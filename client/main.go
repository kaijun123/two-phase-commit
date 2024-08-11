package main

import (
	"fmt"
	"net"
	"two-phase-commit/utils"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	testRequest := utils.SerializeCoordinatorRequest("testKey", "testValue")

	// Send data to the server
	_, err = conn.Write(testRequest)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read and process data from the server
	// ...
}
