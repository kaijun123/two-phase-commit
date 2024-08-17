package participants

import (
	"log"
	"net"
	"time"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

// TODO: Add fault tolerance
func TwoPhaseCommit(conn net.Conn, participantStateMap *ParticipantStateMap, req *p.CoordinatorRequest) {
	PreparePhase(conn, participantStateMap, req)

	// check if all the participants are ready
	if !participantStateMap.CheckAllPrepared() {
		log.Println("Not all participants are prepared")
		return
	}

	CommitPhase(conn, participantStateMap, req)
}

func PreparePhase(conn net.Conn, participantStateMap *ParticipantStateMap, req *p.CoordinatorRequest) {
	prepareReq := utils.SerializeParticipantRequest(p.ParticipantRequestType_PREPARE, false, (*req).Key, (*req).Value)
	participantStateMap.BroadcastAndListen(prepareReq, true, participantStateMap.UpdateParticipantStatus)
	// participantStateMap.ToString()
}

func CommitPhase(conn net.Conn, participantStateMap *ParticipantStateMap, req *p.CoordinatorRequest) {
	prepareReq := utils.SerializeParticipantRequest(p.ParticipantRequestType_COMMIT, false, (*req).Key, (*req).Value)
	participantStateMap.BroadcastAndListen(prepareReq, true, participantStateMap.UpdateParticipantStatus)
	// participantStateMap.ToString()
}

func SendHearbeat(participantStateMap *ParticipantStateMap) {
	for {
		prepareReq := utils.SerializeParticipantRequest(p.ParticipantRequestType_CONNECT, false, "", "")
		participantStateMap.BroadcastAndListen(prepareReq, false, participantStateMap.UpdateParticipantStatus)
		time.Sleep(utils.HeartbeatFrequency * time.Second)
	}
}

func TrackHeartbeat(participantStateMap *ParticipantStateMap) {
	for {
		now := time.Now()
		for ip, state := range participantStateMap.States {
			if now.Sub(state.PreviousHeartbeatTime).Seconds() > utils.HeartbeatThreshold {
				state.IsAlive = false
				participantStateMap.States[ip] = state
			}
		}

		participantStateMap.ToString()
		time.Sleep(utils.HeartbeatFrequency * time.Second)
	}
}
