package participants

import (
	"fmt"
	"log"
	"net"
)

// TODO: Current implementation is optimistic. Assumes that none of the participants can die.
// Need some form of hearbeat and mechanism to track who is alive and dead, to allow for some form of fault-tolerance

type ParticipantStatus string

const (
	Prepared  ParticipantStatus = "prepared"
	Committed ParticipantStatus = "committed"
)

type ParticipantState struct {
	// IsAlive bool
	Ip         string
	Status     ParticipantStatus
	ReqChannel chan []byte // buffered channel
	ResChannel chan []byte // non-buffered channel
	Conn       net.Conn
}

type ParticipantStateMap struct {
	States map[string]ParticipantState
}

func CreateParticipantStateMap() *ParticipantStateMap {
	return &ParticipantStateMap{
		States: make(map[string]ParticipantState),
	}
}

func (p *ParticipantStateMap) AddParticipant(ip string, conn net.Conn) (*ParticipantState, error) {
	state := ParticipantState{
		// IsAlive: isAlive,
		Ip:         ip,
		Status:     "connected",
		ReqChannel: make(chan []byte, 1), // buffered channel
		ResChannel: make(chan []byte),    // non-buffered channel
		Conn:       conn,
	}

	_, exists := p.States[ip]
	if !exists {
		p.States[ip] = state
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
func (p *ParticipantStateMap) Broadcast(message []byte) {
	for _, state := range (*p).States {
		state.ReqChannel <- message
	}
}

// listen to responses from the participants
func (p *ParticipantStateMap) Listen(callback func(s ParticipantState, response []byte)) {
	for _, state := range (*p).States {
		response := <-state.ResChannel
		callback(state, response)
	}
}

func (p *ParticipantStateMap) UpdateParticipantStatus(state ParticipantState, status []byte) {
	ip := state.Ip
	state, exists := p.States[ip]
	if exists {
		fmt.Printf("ip: %s state: %s\n", ip, ParticipantStatus(status))
		state.Status = ParticipantStatus(status)
		fmt.Println("state:", state.Status)
	} else {
		log.Fatal("participant not in the participantStateMap")
	}
}
