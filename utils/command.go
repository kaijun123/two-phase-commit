package utils

import (
	"os/exec"

	"log"
)

func PauseParticipant(port string) []byte {
	cmd := exec.Command("kill", "-STOP", port)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("could not run command: ", err)
	}
	return out
}

func UnpauseParticipant(port string) []byte {
	cmd := exec.Command("kill", "-CONT", port)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("could not run command: ", err)
	}
	return out
}
