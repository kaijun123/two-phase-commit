package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"two-phase-commit/participant/store"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

type participantStatus string

// TODO: use ParticipantRequestType instead for more consistency
// TODO: should pause be one of the states?
// TODO: do we need "committed"
const (
	Default  participantStatus = "default"
	Prepared participantStatus = "prepared"
	// Committed ParticipantStatus = "committed"
)

type participantState struct {
	mutex         sync.Mutex // mutex to prevent race condition
	coordinatorIp string
	clientIp      string
	isPause       bool
	status        participantStatus
	prepareKey    string
	prepareValue  string
}

var s *store.KVStore
var state participantState

func main() {
	bootstrap()
	fmt.Println("coordinatorIp:", state.coordinatorIp)
	fmt.Println("clientIp:", state.clientIp)

	go listenTCP(state.clientIp, handler)
	listenTCP(state.coordinatorIp, handler)
}

func handler(ip string, conn net.Conn) {
	for {
		// Listen for requests
		request := make([]byte, 1024)
		n, err := conn.Read(request)
		if err != nil {
			// TCP connection closed
			// If lose connection with coordinator, data becomes stale. no more reads?
			if err == io.EOF {
				break
			}
			log.Fatal("Error:", err)
		}

		// lock mutex
		state.mutex.Lock()

		// Deserialize the request
		participantReq := utils.DeserializeParticipantRequest(request[:n])
		log.Println("Request received by participant", ip, ":", participantReq.String())

		var participantRes []byte
		t := participantReq.GetType()
		switch t {
		case p.ParticipantRequestType_PREPARE:
			if !state.isPause {
				participantRes = utils.SerializeParticipantResponse(t, true, "")
				state.status = Prepared
				state.prepareKey = participantReq.GetKey()
				state.prepareValue = participantReq.GetValue()
			}
		case p.ParticipantRequestType_COMMIT:
			if !state.isPause && state.status == Prepared && state.prepareKey == participantReq.GetKey() && state.prepareValue == participantReq.GetValue() {
				participantRes = utils.SerializeParticipantResponse(t, true, "")
				s.Put(participantReq.GetKey(), participantReq.GetValue())
				state.status = Default
			}
		case p.ParticipantRequestType_PAUSE:
			if participantReq.GetIsAdmin() {
				state.isPause = true
				participantRes = utils.SerializeParticipantResponse(t, true, "")
			}
		case p.ParticipantRequestType_UNPAUSE:
			if participantReq.GetIsAdmin() {
				state.isPause = false
			}
		case p.ParticipantRequestType_READ:
			if !state.isPause {
				value, err := s.Get(*participantReq.Key)
				if err != nil {
					log.Fatal("unable to get key-value pair from the store:", err.Error())
				}
				participantRes = utils.SerializeParticipantResponse(t, true, value)
			}
		default:
			participantRes = utils.SerializeParticipantResponse(t, false, "")
		}

		if len(participantRes) == 0 {
			participantRes = utils.SerializeParticipantResponse(t, false, "")
		}

		// unlock mutex
		state.mutex.Unlock()

		// Return acknowledges
		_, err = conn.Write(participantRes)
		// fmt.Println("recipient:", conn.RemoteAddr().String())
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
}

func bootstrap() {
	// Set coordinator and client port
	portPtr := flag.String("port", "8081", "port")
	flag.Parse()
	coordinatorIp := "localhost:" + *portPtr
	clientIp := "localhost:" + "1" + *portPtr // client port = coordinator port + 10000

	state = participantState{
		coordinatorIp: coordinatorIp,
		clientIp:      clientIp,
		isPause:       false,
		status:        Default,
		prepareKey:    "",
		prepareValue:  "",
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
