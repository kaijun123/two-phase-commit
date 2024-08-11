package participants

import (
	"fmt"
	"log"
	"net"
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
	// IsAlive bool
	Ip         string
	Status     p.ParticipantRequestType
	ReqChannel chan []byte // buffered channel
	ResChannel chan []byte // non-buffered channel
	Conn       net.Conn
}

func (s ParticipantState) ToString() {
	fmt.Printf("ip: %s, status: %s\n", s.Ip, s.Status.String())
}

type ParticipantStateMap struct {
	States map[string]ParticipantState
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
	state := ParticipantState{
		// IsAlive: isAlive,
		Ip:         ip,
		Status:     p.ParticipantRequestType_CONNECTED,
		ReqChannel: make(chan []byte, 1), // buffered channel
		ResChannel: make(chan []byte),    // non-buffered channel
		Conn:       conn,
	}

	_, exists := m.States[ip]
	if !exists {
		m.States[ip] = state
		return &state, nil
	}
	return nil, fmt.Errorf("participant already exists")
}

// func (p *ParticipantStateMap) UpdateParticipantState(ip string, isAlive bool) error {
// 	if state, exists := p.States[ip]; exists {
// 		state.IsAlive = isAlive
// 		return nil
// 	}
// 	return fmt.Errorf("participant does not exist")
// }

// broadcasts message to all the participant goroutines
func (m *ParticipantStateMap) Broadcast(message []byte) {
	for _, state := range (*m).States {
		state.ReqChannel <- message
	}
}

// listen to responses from the participants
func (m *ParticipantStateMap) Listen(callback func(ip string, response []byte)) {
	for _, state := range (*m).States {
		response := <-state.ResChannel
		callback(state.Ip, response)
	}
}

func (m *ParticipantStateMap) UpdateParticipantStatus(ip string, response []byte) {
	state, exists := m.States[ip]

	if exists {
		participantRes := utils.DeserializeParticipantResponse(response)
		fmt.Println("participantRes:", participantRes.Type.String(), participantRes.GetStatus(), participantRes.GetValue())

		if (*participantRes).Status {
			state.Status = participantRes.Type
			m.States[ip] = state
		}

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
