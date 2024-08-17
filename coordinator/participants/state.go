package participants

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

// TODO: Current implementation is optimistic. Assumes that none of the participants can die.
// Need some form of hearbeat and mechanism to track who is alive and dead, to allow for some form of fault-tolerance

// type ParticipantStatus string

// const (
// 	Prepared  ParticipantStatus = "prepared"
// 	Committed ParticipantStatus = "committed"
// )

type ParticipantState struct {
	Conn                  net.Conn
	Ip                    string
	IsAlive               bool
	Status                p.ParticipantRequestType
	ReqChannel            chan []byte // buffered channel
	ResChannel            chan []byte // non-buffered channel
	PreviousHeartbeatTime time.Time
}

func (s ParticipantState) ToString() {
	fmt.Printf("ip: %v, isAlive: %v, status: %v\n", s.Ip, s.IsAlive, s.Status.String())
}

type ParticipantStateMap struct {
	States map[string]ParticipantState
	Mutex  sync.Mutex
}

func CreateParticipantStateMap() *ParticipantStateMap {
	return &ParticipantStateMap{
		States: make(map[string]ParticipantState),
	}
}

func (m *ParticipantStateMap) ToString() {
	for _, v := range (*m).States {
		v.ToString()
	}
}

func (m *ParticipantStateMap) AddParticipant(ip string, conn net.Conn) (*ParticipantState, error) {
	m.Mutex.Lock()
	state := ParticipantState{
		Conn:                  conn,
		Ip:                    ip,
		IsAlive:               false,
		Status:                p.ParticipantRequestType_DISCONNECT,
		ReqChannel:            make(chan []byte, 1), // buffered channel
		ResChannel:            make(chan []byte),    // non-buffered channel
		PreviousHeartbeatTime: time.Now(),
	}

	_, exists := m.States[ip]
	if !exists {
		m.States[ip] = state
		m.Mutex.Unlock()
		return &state, nil
	}
	m.Mutex.Unlock()
	return nil, fmt.Errorf("participant already exists")
}

// broadcasts message to participant goroutines
// onlyAliveParticipants is used to set if we only want to broadcast to alive participants
func (m *ParticipantStateMap) Broadcast(message []byte, onlyAliveParticipants bool) {
	for _, state := range (*m).States {
		if onlyAliveParticipants && !state.IsAlive {
			continue
		}
		state.ReqChannel <- message
	}
}

// listen to responses from the participants
// onlyAliveParticipants is used to set if we only want to listen from alive participants
func (m *ParticipantStateMap) Listen(onlyAliveParticipants bool, callback func(ip string, response []byte)) {
	for _, state := range (*m).States {
		if onlyAliveParticipants && !state.IsAlive {
			continue
		}
		response := <-state.ResChannel
		callback(state.Ip, response)
	}
}

func (m *ParticipantStateMap) BroadcastAndListen(message []byte, onlyAliveParticipants bool, callback func(ip string, response []byte)) {
	m.Mutex.Lock()

	// broadcast message to all participants
	m.Broadcast(message, onlyAliveParticipants)
	// listen to responses from participants
	m.Listen(onlyAliveParticipants, callback)

	m.Mutex.Unlock()
}

func (m *ParticipantStateMap) UpdateParticipantStatus(ip string, response []byte) {
	state, exists := m.States[ip]

	if exists {
		participantRes := utils.DeserializeParticipantResponse(response)
		fmt.Println("participantRes:", participantRes.Type.String(), participantRes.GetStatus(), participantRes.GetValue())

		t := participantRes.GetType()
		status := participantRes.GetStatus()
		switch t {
		case p.ParticipantRequestType_CONNECT:
			if status {
				state.IsAlive = true
				state.PreviousHeartbeatTime = time.Now()
				state.Status = t
			}
		case p.ParticipantRequestType_DISCONNECT: // occurs when the participant did not return a response and is assumed to be paused
			if status {
				state.IsAlive = false
				state.Status = t
			}
		default:
			state.Status = t
		}
		m.States[ip] = state
	} else {
		log.Fatal("participant not in the participantStateMap")
	}
}

func (m *ParticipantStateMap) CheckAllPrepared() bool {
	for _, state := range m.States {
		if state.Status != p.ParticipantRequestType_PREPARE {
			return false
		}
	}
	return true
}
