package main

import (
	"fmt"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

func main() {
	resp := utils.SendCoordinatorRequest("8080", "testKey", "testValue")
	fmt.Println("status:", resp.GetStatus())

	// resp := utils.SendParticipantRequest("8080", p.ParticipantRequestType_PREPARE, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	// resp = utils.SendParticipantRequest("8080", p.ParticipantRequestType_COMMIT, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	resp2 := utils.SendParticipantRequest("18081", p.ParticipantRequestType_READ, true, "", "")
	fmt.Println("type:", resp2.GetType(), "status:", resp2.GetStatus())

	// resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_PAUSE, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	// resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_PREPARE, true, "testKey", "testValue")

	// resp = utils.SendParticipantRequest("8080", p.ParticipantRequestType_PREPARE, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	// resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_UNPAUSE, true, "", "")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())
}
