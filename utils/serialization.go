package utils

import (
	"log"

	"google.golang.org/protobuf/proto"

	p "two-phase-commit/proto"
)

func SerializeParticipantRequest(t p.ParticipantRequestType, a bool, k string, v string) []byte {
	participantRequest := p.ParticipantRequest{
		Type:    t,
		IsAdmin: a,
		Key:     proto.String(k),
		Value:   proto.String(v),
	}

	bytes, err := proto.Marshal(&participantRequest)
	if err != nil {
		log.Fatal("unable to marshal CoordinatorRequest:", err.Error())
	}
	return bytes
}

func DeserializeParticipantRequest(bytes []byte) *p.ParticipantRequest {
	var participantRequest p.ParticipantRequest

	if err := proto.Unmarshal(bytes, &participantRequest); err != nil {
		log.Fatal("unable to unmarshal ParticipantRequest:", err.Error())
	}

	return &participantRequest
}

func SerializeParticipantResponse(t p.ParticipantRequestType, s bool, v string) []byte {
	participantResponse := p.ParticipantResponse{
		Type:   t,
		Status: s,
		Value:  proto.String(v),
	}

	// fmt.Println("participantResponse.Type:", participantResponse.Type, "participantResponse.Status:", participantResponse.Status)

	bytes, err := proto.Marshal(&participantResponse)

	if err != nil {
		log.Fatal("unable to marshal ParticipantResponse:", err.Error())
	}

	return bytes
}

func DeserializeParticipantResponse(bytes []byte) *p.ParticipantResponse {
	var participantResponse p.ParticipantResponse

	if err := proto.Unmarshal(bytes, &participantResponse); err != nil {
		log.Fatal("unable to unmarshal ParticipantResponse:", err.Error())
	}

	return &participantResponse
}

func SerializeCoordinatorRequest(k string, v string) []byte {
	bytes, err := proto.Marshal(&p.CoordinatorRequest{
		Key:   k,
		Value: v,
	})

	if err != nil {
		log.Fatal("unable to marshal CoordinatorRequest:", err.Error())
	}
	return bytes
}

func DeserializeCoordinatorRequest(bytes []byte) *p.CoordinatorRequest {
	var coordinatorRequest p.CoordinatorRequest

	if err := proto.Unmarshal(bytes, &coordinatorRequest); err != nil {
		log.Fatal("unable to unmarshal CoordinatorRequest:", err.Error())
	}

	return &coordinatorRequest
}

func SerializeCoordinatorResponse(status bool) []byte {
	bytes, err := proto.Marshal(&p.CoordinatorResponse{
		Status: status,
	})

	if err != nil {
		log.Fatal("unable to marshal CoordinatorResponse:", err.Error())
	}
	return bytes
}

func DeserializeCoordinatorResponse(bytes []byte) *p.CoordinatorResponse {
	var coordinatorResponse p.CoordinatorResponse

	if err := proto.Unmarshal(bytes, &coordinatorResponse); err != nil {
		log.Fatal("unable to unmarshal coordinatorResponse:", err.Error())
	}

	return &coordinatorResponse
}
