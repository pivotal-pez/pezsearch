package pezinventory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezinventory"
)

var _ = Describe("TypeController", func() {
	var ctrl *TypeController

	BeforeEach(func() {
		ctrl = &TypeController{}
	})

	Context("List Types", func() {
		It("Should return a list of Types", func() {
			Ω(func() {
				res := ctrl.ListTypes()
				Ω(res).ShouldNot(BeNil())
				Ω(res).Should(Equal("List of Types"))
			}).ShouldNot(Panic())
		})
	})

	Context("Get Type Details", func() {
		It("Should return the details of a given Type", func() {
			Ω(func() {
				p := make(map[string]string)
				p["id"] = "12345"
				res := ctrl.GetType(p)
				Ω(res).ShouldNot(BeNil())
				Ω(res).Should(Equal("Details of Type-" + p["id"]))
			}).ShouldNot(Panic())
		})
	})

	Context("List Type Items", func() {
		It("Should return a list of Items for a given Type", func() {
			Ω(func() {
				p := make(map[string]string)
				p["id"] = "12345"
				res := ctrl.ListTypeItems(p)
				Ω(res).ShouldNot(BeNil())
				Ω(res).Should(Equal("List of Items for Type-" + p["id"]))
			}).ShouldNot(Panic())
		})
	})
})
