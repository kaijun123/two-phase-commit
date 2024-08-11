package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"two-phase-commit/participant/store"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

var s *store.KVStore

var isPause = false

var ip string

// TODO: refactor
var state ParticipantStatus = Default
var prepareKey string = ""
var prepareValue string = ""

type ParticipantStatus string

// TODO: use ParticipantRequestType instead for more consistency
// TODO: should pause be one of the states?
// TODO: do we need "committed"
const (
	Default  ParticipantStatus = "default"
	Prepared ParticipantStatus = "prepared"
	// Committed ParticipantStatus = "committed"
)

func main() {
	// Set port
	portPtr := flag.String("port", "8081", "port")
	flag.Parse()
	ip = "localhost:" + *portPtr

	// Initialize Store
	s = store.InitializeStore()

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

		handler(conn)
	}
}

func handler(conn net.Conn) {
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

		// Deserialize the request
		participantReq := utils.DeserializeParticipantRequest(request[:n])
		log.Println("Request received by participant", ip, ":", participantReq.String())

		var participantRes []byte = []byte{}
		switch t := participantReq.GetType(); t {
		case p.ParticipantRequestType_PREPARE:
			if !isPause {
				participantRes = utils.SerializeParticipantResponse(t, true, "")
				state = Prepared
				prepareKey = participantReq.GetKey()
				prepareValue = participantReq.GetValue()
			}
		case p.ParticipantRequestType_COMMIT:
			if !isPause && state == Prepared && prepareKey == participantReq.GetKey() && prepareValue == participantReq.GetValue() {
				participantRes = utils.SerializeParticipantResponse(t, true, "")
				s.Put(participantReq.GetKey(), participantReq.GetValue())
				state = Default
			}
		case p.ParticipantRequestType_PAUSE:
			if participantReq.GetIsAdmin() {
				isPause = true
				participantRes = utils.SerializeParticipantResponse(t, true, "")
			}
		case p.ParticipantRequestType_UNPAUSE:
			if participantReq.GetIsAdmin() {
				isPause = false
			}
		case p.ParticipantRequestType_READ:
			if !isPause {
				value, err := s.Get(*participantReq.Key)
				if err != nil {
					log.Fatal("unable to get key-value pair from the store:", err.Error())
				}
				participantRes = utils.SerializeParticipantResponse(t, true, value)
			}
		default:
			participantRes = utils.SerializeParticipantResponse(t, false, "")
		}
		// Return acknowledges
		_, err = conn.Write(participantRes)
		if err != nil {
			// TCP connection closed
			// coordinator dies?
			if err == io.EOF {
				break
			}
			log.Fatal("Error:", err)
		}
	}
}
