package main_test

import (
	"testing"

	"two-phase-commit/utils"

	p "two-phase-commit/proto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	testKey    = "testKey"
	testValue  = "testValue"
	testValue2 = "testValue2"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var _ = Describe("Basic participant functions", func() {
	It("Partitioned Participant: Participant receives Prepare message, but not Commit", func() {
		resp := utils.SendParticipantRequest("18081", p.MessageType_Prepare, false, testKey, testValue, true)
		Expect(resp.GetType()).To(Equal(p.MessageType_Prepare))
		Expect(resp.GetStatus()).To(Equal(true))

		resp = utils.SendParticipantRequest("18081", p.MessageType_GetStatus, false, "", "", false)
		Expect(resp.GetType()).To(Equal(p.MessageType_GetStatus))
		Expect(resp.GetStatus()).To(Equal(true))
		Expect(resp.GetValue()).To(Equal("abort"))
	})

	// It("Basic Prepare-Commit flow", func() {
	// 	// Issue: with the current locking mechanism, what happens when the coord dies, and reboots and resends a global commit?
	// 	// Or should the restart process for coord be from prepare. Then make sure that the participants resend the same votes?
	// })

	It("Successful Prepare-Commit flow", func() {
		By("Test Prepare")
		resp := utils.SendParticipantRequest("18081", p.MessageType_Prepare, false, testKey, testValue, false)
		Expect(resp.GetType()).To(Equal(p.MessageType_Prepare))
		Expect(resp.GetStatus()).To(Equal(true))

		// By("Test Prepare Status")
		// resp = utils.SendParticipantRequest("18081", p.MessageType_GetStatus, false, "", "", false)
		// fmt.Println(resp.GetType().String())
		// Expect(resp.GetType()).To(Equal(p.MessageType_GetStatus))
		// Expect(resp.GetStatus()).To(Equal(true))
		// Expect(resp.GetValue()).To(Equal("ready"))

		By("Test Commit")
		resp = utils.SendParticipantRequest("18081", p.MessageType_Commit, false, testKey, testValue, false)
		Expect(resp.GetType()).To(Equal(p.MessageType_Commit))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Commit Status")
		resp = utils.SendParticipantRequest("18081", p.MessageType_GetStatus, false, "", "", false)
		Expect(resp.GetType()).To(Equal(p.MessageType_GetStatus))
		Expect(resp.GetStatus()).To(Equal(true))
		Expect(resp.GetValue()).To(Equal("commit"))

		By("Test Read after commit")
		resp = utils.SendParticipantRequest("18081", p.MessageType_Read, false, testKey, "", false)
		Expect(resp.GetType()).To(Equal(p.MessageType_Read))
		Expect(resp.GetStatus()).To(Equal(true))
		Expect(resp.GetValue()).To(Equal(testValue))

		By("Test Prepare")
		resp = utils.SendParticipantRequest("18081", p.MessageType_Prepare, false, testKey, testValue2, false)
		Expect(resp.GetType()).To(Equal(p.MessageType_Prepare))
		Expect(resp.GetStatus()).To(Equal(true))

		// By("Test Prepare Status")
		// resp = utils.SendParticipantRequest("18081", p.MessageType_GetStatus, false, "", "", false)
		// Expect(resp.GetType()).To(Equal(p.MessageType_GetStatus))
		// Expect(resp.GetStatus()).To(Equal(true))
		// Expect(resp.GetValue()).To(Equal("ready"))

		By("Test Abort")
		resp = utils.SendParticipantRequest("18081", p.MessageType_Abort, false, testKey, testValue2, false)
		Expect(resp.GetType()).To(Equal(p.MessageType_Abort))
		Expect(resp.GetStatus()).To(Equal(true))

		By("Test Abort Status")
		resp = utils.SendParticipantRequest("18081", p.MessageType_GetStatus, false, "", "", false)
		Expect(resp.GetType()).To(Equal(p.MessageType_GetStatus))
		Expect(resp.GetStatus()).To(Equal(true))
		Expect(resp.GetValue()).To(Equal("abort"))

		By("Test Read after abort")
		resp = utils.SendParticipantRequest("18081", p.MessageType_Read, false, testKey, "", false)
		Expect(resp.GetType()).To(Equal(p.MessageType_Read))
		Expect(resp.GetStatus()).To(Equal(true))
		Expect(resp.GetValue()).To(Equal(testValue))

		By("Test Delete")
		resp = utils.SendParticipantRequest("18081", p.MessageType_Delete, true, testKey, "", false)
		Expect(resp.GetType()).To(Equal(p.MessageType_Delete))
		Expect(resp.GetStatus()).To(Equal(true))
	})
})
