package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"two-phase-commit/coordinator/participants"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

var participantStateMap *participants.ParticipantStateMap

func main() {
	participantStateMap = participants.CreateParticipantStateMap()
	bootstrap()
	// participantStateMap.ToString()

	// // send first heartbeat
	go participants.SendHearbeat(participantStateMap)
	go participants.TrackHeartbeat(participantStateMap)
	// participantStateMap.ToString()

	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	// Accept connections
	for {
		log.Println("waiting for new connections")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error:", err)
		}

		for {
			request := make([]byte, 1024)
			n, err := conn.Read(request)
			if err != nil {
				if err == io.EOF { // client closed connection
					log.Println("Client closed the connection.")
					break
				}
				log.Fatal("Error:", err)
			}
			log.Println("Message from client:", string(request[:n]))

			coordinatorRequest := utils.DeserializeCoordinatorRequest(request[:n])

			participants.TwoPhaseCommit(conn, participantStateMap, coordinatorRequest)
		}
	}
}

func bootstrap() {
	ipArray := utils.ReadConfigFile("../participants.txt", false)
	fmt.Println("ipArray:", ipArray)

	for _, ip := range ipArray {
		// Connect to the participants
		conn, err := net.Dial("tcp", ip)
		log.Println("Connected to participants at", ip)
		if err != nil {
			log.Fatal(err)
		}

		// add participants
		state, err := participantStateMap.AddParticipant(ip, conn)
		if err != nil {
			log.Fatal(err)
		}

		// handle participants
		go handleParticipants(state)
	}
}

func handleParticipants(state *participants.ParticipantState) {
	conn := state.Conn
	reqChannel := state.ReqChannel
	resChannel := state.ResChannel

	defer conn.Close()

	for {
		request := <-reqChannel

		// send packets to participants
		n, err := conn.Write(request)
		if err != nil {
			log.Fatal("Error1:", err)
			return
		}

		log.Println("Request to participant at", state.Ip, ":", string(request[:n]))

		// read packets from participants
		response := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(utils.ReadDeadlineDuration * time.Second))
		n, err = conn.Read(response)
		if err != nil {
			fmt.Println("Error2:", err)
			response = utils.SerializeParticipantResponse(p.MessageType_Disconnect, true, "")
			n = len(response)
		}
		log.Println("Response from participant at", state.Ip, ":", string(response[:n]))

		resChannel <- response[:n]
	}
}
