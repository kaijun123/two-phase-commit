package main

import (
	"2PC/participant/store"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var s *store.KVStore

func main() {
	// Set port
	portPtr := flag.String("port", "8081", "port")
	flag.Parse()
	ip := "localhost:" + *portPtr

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
			log.Println("Request received by participant", *portPtr, ":", request[:n])

			// Return acknowledges
			response := []byte("prepared")
			_, err = conn.Write(response)
			if err != nil {
				log.Fatal("Error:", err)
			}
		}
	}
}
