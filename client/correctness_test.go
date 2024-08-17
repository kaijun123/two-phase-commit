package main_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	p "two-phase-commit/proto"
	"two-phase-commit/utils"
)

const (
	testKey   = "testKey"
	testValue = "testValue"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var _ = Describe("Basic participant functions", func() {
	It("Basic pause & unpause by admin and non-admin", func() {
	})

	It("Basic pause & unpause by admin and non-admin", func() {
		By("Test pause by non-admin")
		resp := utils.SendParticipantRequest("18081", p.ParticipantRequestType_PAUSE, false, "", "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_PAUSE))
		Expect(resp.GetStatus()).To(Equal(false))

		By("Test pause by admin")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_PAUSE, true, "", "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_PAUSE))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test unpause by non-admin")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_UNPAUSE, false, "", "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_UNPAUSE))
		Expect(resp.GetStatus()).To(Equal(false))

		By("Test unpause by admin")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_UNPAUSE, true, "", "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_UNPAUSE))
		Expect(resp.GetStatus()).To(Equal(true))
	})

	It("Basic participant flow", func() {
		By("Test pause")
		resp := utils.SendParticipantRequest("18081", p.ParticipantRequestType_PAUSE, true, "", "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_PAUSE))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Prepare after pause")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_PREPARE, false, testKey, testValue)
		fmt.Println(resp.String())
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_PREPARE))
		Expect(resp.GetStatus()).To(Equal(false))

		By("Test Unpause after pause")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_UNPAUSE, true, "", "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_UNPAUSE))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Prepare after unpause")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_PREPARE, false, testKey, testValue)
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_PREPARE))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Commit after unpause")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_COMMIT, false, testKey, testValue)
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_COMMIT))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Read after commit")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_READ, false, testKey, "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_READ))
		Expect(resp.GetValue()).To(Equal(testValue))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Delete after commit")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_DELETE, true, testKey, "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_DELETE))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Read after commit")
		resp = utils.SendParticipantRequest("18081", p.ParticipantRequestType_READ, false, testKey, "")
		Expect(resp.GetType()).To(Equal(p.ParticipantRequestType_READ))
		Expect(resp.GetStatus()).To(Equal(false))
	})
})

// func main() {

// 	// Case 1: Test participant functionalities

// 	sendRequest("18081", p.ParticipantRequestType_PAUSE, true, "", "")
// 	sendRequest("8081", p.ParticipantRequestType_PREPARE, true, "testKey", "testValue")

// 	// participantRes := utils.SerializeParticipantResponse(p.ParticipantRequestType_PREPARE, false, "")
// 	// fmt.Println("length:", len(participantRes))

// 	// participantResponse := p.ParticipantResponse{
// 	// 	Type:   p.ParticipantRequestType_PREPARE,
// 	// 	Status: false,
// 	// 	Value:  proto.String(""),
// 	// }

// 	// bytes, err := proto.Marshal(&participantResponse)
// 	// if err != nil {
// 	// 	log.Fatal("unable to marshal ParticipantResponse:", err.Error())
// 	// }
// 	// fmt.Println("length:", len(bytes))
// }
