package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// if wantClientPort == true => return the
func ReadConfigFile(path string, wantClientPort bool) []string {
	ipArray := []string{}

	// obtain the relative path to the participants file
	relativePath, _ := filepath.Abs(path)

	file, err := os.Open(relativePath)
	if err != nil {
		log.Fatal(err)
	}

	// Close the file
	defer file.Close()

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := scanner.Text()
		domain := strings.Split(ip, ":")[0]
		port := strings.Split(ip, ":")[1]

		if wantClientPort {
			port = "1" + port
			ip = domain + ":" + port
		}
		ipArray = append(ipArray, ip)
	}
	// check for the error that occurred during the scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ipArray
}
