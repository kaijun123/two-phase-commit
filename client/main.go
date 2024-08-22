package main

import (
	"fmt"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

func main() {
	participantRes := utils.SerializeParticipantResponse(p.MessageType_Prepare, false, "", "", "")
	fmt.Println(len(participantRes))

	// resp := utils.SendCoordinatorRequest("8080", "testKey", "testValue")
	// fmt.Println("status:", resp.GetStatus())

	// resp := utils.SendParticipantRequest("8080", p.MessageType_Prepare, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	// resp = utils.SendParticipantRequest("8080", p.MessageType_Commit, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	// utils.SendParticipantRequest("18081", p.MessageType_Prepare, true, "testKey", "testValue", false)
	// fmt.Println("type:", resp2.GetType(), "status:", resp2.GetStatus())

	// resp = utils.SendParticipantRequest("18081", p.MessageType_PAUSE, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	// resp = utils.SendParticipantRequest("18081", p.MessageType_Prepare, true, "testKey", "testValue")

	// resp = utils.SendParticipantRequest("8080", p.MessageType_Prepare, true, "testKey", "testValue")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())

	// resp = utils.SendParticipantRequest("18081", p.MessageType_UNPAUSE, true, "", "")
	// fmt.Println("type:", resp.GetType(), "status:", resp.GetStatus())
}
