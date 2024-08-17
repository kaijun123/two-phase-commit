package utils

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	p "two-phase-commit/proto"
)

func SendParticipantRequest(port string, reqType p.ParticipantRequestType, isAdmin bool, key string, value string) *p.ParticipantResponse {
	conn := connect(port)
	defer conn.Close() // Ensure the connection is closed properly

	// Serialize the request
	request := SerializeParticipantRequest(reqType, isAdmin, key, value)
	send(conn, request)

	// Read response
	response := make([]byte, 1024) // Ensure buffer size is appropriate
	conn.SetReadDeadline(time.Now().Add(ReadDeadlineDuration * time.Second))
	n, err := conn.Read(response)
	if err != nil {
		if err == io.EOF {
			return nil // Connection closed
		}
		fmt.Println("Error:", err)
		return nil
	}

	fmt.Println(DeserializeParticipantResponse(response[:n]).String())
	return DeserializeParticipantResponse(response[:n])
}

func SendCoordinatorRequest(port string, key string, value string) *p.CoordinatorResponse {
	conn := connect(port)
	defer conn.Close() // Ensure the connection is closed properly

	// Serialize the request
	request := SerializeCoordinatorRequest(key, value)
	send(conn, request)

	// Read response
	response := make([]byte, 1024) // Ensure buffer size is appropriate
	n, err := conn.Read(response)
	if err != nil {
		if err == io.EOF {
			return nil // Connection closed
		}
		fmt.Println("Error:", err)
		return nil
	}

	fmt.Println(DeserializeCoordinatorResponse(response[:n]).String())
	return DeserializeCoordinatorResponse(response[:n])
}

func connect(port string) net.Conn {
	// Make tcp connection
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return conn
}

func send(conn net.Conn, message []byte) {
	// Send data to the server
	_, err := conn.Write(message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
