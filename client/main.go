package main

import (
	"fmt"
	"io"
	"log"
	"net"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

func main() {
	sendRequest("18081", p.ParticipantRequestType_PAUSE, true, "", "")
	sendRequest("8081", p.ParticipantRequestType_PREPARE, true, "testKey", "testValue")

	// participantRes := utils.SerializeParticipantResponse(p.ParticipantRequestType_PREPARE, false, "")
	// fmt.Println("length:", len(participantRes))

	// participantResponse := p.ParticipantResponse{
	// 	Type:   p.ParticipantRequestType_PREPARE,
	// 	Status: false,
	// 	Value:  proto.String(""),
	// }

	// bytes, err := proto.Marshal(&participantResponse)
	// if err != nil {
	// 	log.Fatal("unable to marshal ParticipantResponse:", err.Error())
	// }
	// fmt.Println("length:", len(bytes))
}

func sendRequest(port string, reqType p.ParticipantRequestType, isAdmin bool, key string, value string) {
	// Make tcp connection
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close() // Ensure the connection is closed properly

	// Serialize the request
	testRequest := utils.SerializeParticipantRequest(reqType, isAdmin, key, value)
	send(conn, testRequest)

	// Read response
	response := make([]byte, 1024) // Ensure buffer size is appropriate
	n, err := conn.Read(response)
	if err != nil {
		if err == io.EOF {
			return // Connection closed
		}
		log.Fatal("Error:", err)
	}

	fmt.Println(utils.DeserializeParticipantResponse(response[:n]).String())
}

func send(conn net.Conn, message []byte) {
	// Send data to the server
	_, err := conn.Write(message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// func listen(conn net.Conn, callback func([]byte) interface{}) any {
// 	for {
// 		response := make([]byte, 1024)
// 		n, err := conn.Read(response)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return nil
// 		}
// 		return callback(response[:n])
// 	}
// }
