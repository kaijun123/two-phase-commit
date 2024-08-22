package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
	"two-phase-commit/participant/store"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

// TODO: should pause be one of the states?
// TODO: do we need "committed"
type State string

const (
	Init         State = "init"
	Ready        State = "ready"
	Commit       State = "commit"
	Abort        State = "abort"
	Inconclusive State = "inconclusive" // only for p2p
)

type participantState struct {
	mutex         sync.Mutex // mutex to prevent race condition
	coordinatorIp string
	clientIp      string
	timeout       bool
	peerIps       []string
	// isPause                 bool
	// isStale                 bool // occurs when the participant is partitioned from the coordinator
	commitState             State
	previousPrepareResponse State
	previousKey             string
	previousValue           string
	// previousHeartbeatTime time.Time
}

var s *store.KVStore
var state participantState

func main() {
	bootstrap()
	fmt.Println("coordinatorIp:", state.coordinatorIp)
	fmt.Println("clientIp:", state.clientIp)

	// go trackHearbeat()
	go listenTCP(state.clientIp, handler)
	listenTCP(state.coordinatorIp, handler)
}

func handler(ip string, conn net.Conn) {
	for {
		if state.commitState == Ready {
			conn.SetReadDeadline(time.Now().Add(utils.ReadDeadlineDuration * time.Second))
		}

		// Listen for requests
		request := make([]byte, 1024)
		n, err := conn.Read(request)
		if err != nil {
			// TODO: coord dies, TCP connection closed: the participant should timeout
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error:", err)

				// Timeout: did not receive commit message after prepare
				if state.commitState == Ready {
					state.timeout = true
				}
			}
		}

		if state.timeout {
			// 2 scenarios:
			// 1) coord died
			// 2) participant is partitioned
			terminationProtocol() // no response needed

		} else {
			// Deserialize the request
			participantReq := utils.DeserializeParticipantRequest(request[:n])
			fmt.Println("Request received by participant", ip, ":", participantReq.String())

			var participantRes []byte
			t := participantReq.GetType()

			// ignore the case whereby the message after prepare is not commit: assume no byzantine/ unforeseen behaviour
			// if state.timeout || (state.commitState == Ready && (t != p.MessageType_Commit && t != p.MessageType_Abort)) {
			// 	t = p.MessageType_Abort
			// }

			fmt.Println(t.String())

			switch t {
			case p.MessageType_Prepare:
				state.mutex.Lock()
				// TODO: how to differentiate between repeated prepare statements and new prepare statements (coord reboot process)
				prepareResponse := utils.RandOutcome()
				participantRes = utils.SerializeParticipantResponse(t, prepareResponse, "", "", "")
				state.commitState = Ready
				if prepareResponse {
					state.previousPrepareResponse = Commit
					state.previousKey = participantReq.GetKey()
					state.previousValue = participantReq.GetValue()
				} else {
					state.previousPrepareResponse = Abort
				}

			case p.MessageType_Commit:
				key := participantReq.GetKey()
				value := participantReq.GetValue()
				if state.commitState == Ready && state.previousKey == key && state.previousValue == value {
					participantRes = utils.SerializeParticipantResponse(t, true, "", "", "")
					s.Put(key, value)
					state.commitState = Commit
					state.mutex.Unlock()
				}

			case p.MessageType_Abort:
				if state.commitState == Ready {
					participantRes = utils.SerializeParticipantResponse(t, true, "", "", "")
					state.commitState = Abort
					state.mutex.Unlock()
				}

			case p.MessageType_Read:
				state.mutex.Lock()
				key := participantReq.GetKey()
				value, err := s.Get(key)
				if err != nil {
					fmt.Println("unable to get key-value pair from the store:", err.Error())
					participantRes = utils.SerializeParticipantResponse(t, false, "", "", "")
				} else {
					participantRes = utils.SerializeParticipantResponse(t, true, "", key, value)
				}
				state.mutex.Unlock()

			case p.MessageType_Delete:
				if participantReq.GetIsAdmin() {
					state.mutex.Lock()
					err := s.Remove(participantReq.GetKey())
					if err != nil {
						fmt.Println("unable to get key-value pair from the store:", err.Error())
						participantRes = utils.SerializeParticipantResponse(t, false, "", "", "")
					} else {
						participantRes = utils.SerializeParticipantResponse(t, true, "", "", "")
					}
					state.mutex.Unlock()
				}

			case p.MessageType_GetPid:
				pid := string(os.Getpid())
				participantRes = utils.SerializeParticipantResponse(t, true, "", "", pid)

			case p.MessageType_GetStatus:
				participantRes = utils.SerializeParticipantResponse(t, true, "", "", string(state.commitState))

			case p.MessageType_P2P:
				if state.previousPrepareResponse == Abort {
					// if sent abort to coordinator, return abort
					participantRes = utils.SerializeParticipantResponse(t, true, string(Abort), "", "")
				} else if state.commitState == Commit {
					// if received commit message from coordinator
					participantRes = utils.SerializeParticipantResponse(t, true, string(Commit), state.previousKey, state.previousValue)
				} else if state.commitState == Abort {
					// if received abort message from coordinator
					participantRes = utils.SerializeParticipantResponse(t, true, string(Abort), "", "")
				} else {
					// if no request received (ie timeout), return inconclusive
					participantRes = utils.SerializeParticipantResponse(t, true, string(Inconclusive), "", "")
				}
			}

			// default response is fail
			if len(participantRes) == 0 {
				participantRes = utils.SerializeParticipantResponse(t, false, "", "", "")
			}

			// Return acknowledges if not timeout
			_, err = conn.Write(participantRes)
			if err != nil {
				// TCP connection closed
				// coordinator dies?
				// what if the connection just dies and the coordinator/ client does not get the response?
				if err == io.EOF {
					break
				}
				log.Fatal("Error:", err)
			}
		}

		conn.SetReadDeadline(time.Time{}) // remove the setReadDeadline
		state.timeout = false
	}
}

func bootstrap() {
	// Set coordinator and client port
	portPtr := flag.String("port", "8081", "port")
	flag.Parse()
	coordinatorIp := "localhost:" + *portPtr
	clientIp := "localhost:" + "1" + *portPtr // client port = coordinator port + 10000

	peerIpArray := utils.ReadConfigFile("../participants.txt", true)
	// fmt.Println("peerIpArray:", peerIpArray)

	state = participantState{
		coordinatorIp: coordinatorIp,
		clientIp:      clientIp,
		timeout:       false,
		peerIps:       peerIpArray,
		// isPause:       false,
		// isStale:       false,
		commitState:             Init,
		previousPrepareResponse: Inconclusive,
		previousKey:             "",
		previousValue:           "",
	}

	// Initialize Store
	s = store.InitializeStore()
}

func listenTCP(ip string, callback func(ip string, conn net.Conn)) {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Participant is running at", ip)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		callback(ip, conn)
	}
}

// Used to determine if the
// If timeout and no participant received any message: safe to assume that the coord died => Abort
// If timeout and one participant died => stuck
// If timeout and one participant received commit => commit
// If timeout and one participant sent abort => abort
func terminationProtocol() {
	// if current participant sent false => abort
	if state.previousPrepareResponse == Abort {
		state.commitState = Abort
		state.mutex.Unlock()
		return
	}

	peerResults := Inconclusive
	peerTimeouts := false
	key := ""
	value := ""
	// P2P with other participants
	for _, ip := range state.peerIps {
		if ip == state.clientIp { // skip if it is current participant
			continue
		} else {
			resp := utils.SendParticipantRequest(ip, p.MessageType_P2P, false, "", "", false)
			if resp == nil {
				// connection time-out: participant assumed to be dead
				peerTimeouts = true
			} else {
				if resp.GetAction() == string(Abort) {
					peerResults = Abort
					break
				} else if resp.GetAction() == string(Commit) {
					peerResults = Commit
					key = resp.GetKey()
					value = resp.GetValue()
					break
				}
			}
		}
	}

	if peerResults == Inconclusive && !peerTimeouts {
		// case 1: all participant alive, and no one received coordinator message => abort
		state.commitState = Abort
		state.previousKey = ""
		state.previousValue = ""

	} else if peerResults == Inconclusive && peerTimeouts {
		// case 2: some participants dead, and those alive did not receive coordinator message => stuck
		// need to unlock mutex for the coordinator to resend messages when alive

	} else if peerResults != Inconclusive {
		// case 3: some participants received a message
		if peerResults == Commit {
			s.Put(key, value)
			state.commitState = Commit
		} else if peerResults == Abort {
			state.commitState = Abort
		}
	}

	state.mutex.Unlock()
}

// func trackHearbeat() {
// 	for {
// 		now := time.Now()
// 		if now.Sub(state.previousHeartbeatTime).Seconds() > utils.HeartbeatThreshold {
// 			state.isStale = true
// 		}
// 		time.Sleep(utils.HeartbeatFrequency * time.Second)
// 	}
// }
