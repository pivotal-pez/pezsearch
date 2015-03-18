package server_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezinventory/server"
)

var _ = Describe("Server", func() {
	var server Server
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	BeforeEach(func() {
		server = NewServer()
		recorder = httptest.NewRecorder()
	})

	Describe("GET /v1/types", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/types", nil)
		})

		Context("when types exist", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns a list of types", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("List of Types"))
			})
		})
	})

	Describe("GET /v1/types/:id", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/types/12345", nil)
		})

		Context("with valid Type id", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns Type details", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("Details of Type"))
			})
		})
	})

	Describe("GET /v1/types/:id/items", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/types/12345/items", nil)
		})

		Context("when Items exist for Type", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns a list of Items for Type", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("List of Items for Type"))
			})
		})
	})

	Describe("GET /v1/items", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/items", nil)
		})

		Context("when Items exist", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns a list of Items", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("List of Items"))
			})
		})
	})

	Describe("GET /v1/items/:id", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/items/12345", nil)
		})

		Context("with valid Item id", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns Item details", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("Details of Item"))
			})
		})
	})

	Describe("GET /v1/items/:id/history", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/items/12345/history", nil)
		})

		Context("when history exists for Item", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns the history for an Item", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("History for Item"))
			})
		})
	})
})
