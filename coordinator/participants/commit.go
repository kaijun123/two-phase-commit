package participants

import (
	"log"
	"net"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

func TwoPhaseCommit(conn net.Conn, participantStateMap *ParticipantStateMap, req *p.CoordinatorRequest) {
	// Prepare Phase
	prepareReq := utils.SerializeParticipantRequest(p.ParticipantRequestType_PREPARE, false, (*req).Key, (*req).Value)

	// broadcast message to all participants
	participantStateMap.Broadcast(prepareReq)
	// listen to responses from participants
	participantStateMap.Listen(participantStateMap.UpdateParticipantStatus)

	// participantStateMap.ToString()

	// check if all the participants are ready
	if !participantStateMap.CheckAllPrepared() {
		log.Println("Not all participants are prepared")
		return
	}

	// Commit Phase
	commitReq := utils.SerializeParticipantRequest(p.ParticipantRequestType_COMMIT, false, (*req).Key, (*req).Value)

	// broadcast message to all participants
	participantStateMap.Broadcast(commitReq)
	// listen to responses from participants
	participantStateMap.ToString()

	// listen to responses from participants
	participantStateMap.Listen(participantStateMap.UpdateParticipantStatus)
}
