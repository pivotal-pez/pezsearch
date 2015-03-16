package pezinventory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezinventory"
)

var _ = Describe("ItemController", func() {
	var ctrl *ItemController

	BeforeEach(func() {
		ctrl = &ItemController{}
	})

	Context("List Items", func() {
		It("Should return a list of Items", func() {
			Ω(func() {
				res := ctrl.ListItems()
				Ω(res).ShouldNot(BeNil())
				Ω(res).Should(Equal("List of Items"))
			}).ShouldNot(Panic())
		})
	})

	Context("Get Item Details", func() {
		It("Should return the details of a given Item", func() {
			Ω(func() {
				p := make(map[string]string)
				p["id"] = "12345"
				res := ctrl.GetItem(p)
				Ω(res).ShouldNot(BeNil())
				Ω(res).Should(Equal("Details of Item-" + p["id"]))
			}).ShouldNot(Panic())
		})
	})

	Context("Get Item History", func() {
		It("Should return the history for a given Item", func() {
			Ω(func() {
				p := make(map[string]string)
				p["id"] = "12345"
				res := ctrl.GetItemHistory(p)
				Ω(res).ShouldNot(BeNil())
				Ω(res).Should(Equal("History for Item-" + p["id"]))
			}).ShouldNot(Panic())
		})
	})
})
