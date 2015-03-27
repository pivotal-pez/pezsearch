package pezsearch_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezsearch"
)

var _ = Describe("Server", func() {
	var server Server
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	BeforeEach(func() {
		server = NewServer()
		recorder = httptest.NewRecorder()
	})

	Describe("GET /v1/search without search parameters", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/search", nil)
		})

		Context("global search", func() {
			It("returns a status code of 400", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(400))
			})

			It("returns a search result", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("You must supply search criteria"))
			})
		})
	})

	Describe("GET /v1/search with empty search parameters", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/search?q", nil)
		})

		Context("global search", func() {
			It("returns a status code of 400", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(400))
			})

			It("returns a search result", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Body).To(ContainSubstring("You must supply search criteria"))
			})
		})
	})

	Describe("GET /v1/search with search parameters", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/search?q=abc", nil)
		})

		Context("global search with results", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns a search result", func() {
				server.ServeHTTP(recorder, request)
				mJSON := mapFromJSON(recorder.Body.Bytes())
				Ω(mJSON["status"]).To(Equal("success"))
				results := sliceFromInterface(mJSON["data"])
				first := results[0].(map[string]interface{})
				last := results[11].(map[string]interface{})
				Ω(len(results)).To(Equal(12))
				Ω(first["_type"]).To(Equal(TypeResource))
				Ω(last["_type"]).To(Equal(ItemResource))
			})
		})
	})

	Describe("GET /v1/search with item scope", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/search?q=abc scope:item", nil)
		})

		Context("global search with results", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns a search result", func() {
				server.ServeHTTP(recorder, request)
				mJSON := mapFromJSON(recorder.Body.Bytes())
				Ω(mJSON["status"]).To(Equal("success"))
				results := sliceFromInterface(mJSON["data"])
				first := results[0].(map[string]interface{})
				last := results[9].(map[string]interface{})
				Ω(len(results)).To(Equal(10))
				Ω(first["_type"]).To(Equal(ItemResource))
				Ω(last["_type"]).To(Equal(ItemResource))
			})
		})
	})

	Describe("GET /v1/search with type scope", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/search?q=abc scope:type", nil)
		})

		Context("global search with results", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Ω(recorder.Code).To(Equal(200))
			})

			It("returns a search result", func() {
				server.ServeHTTP(recorder, request)
				mJSON := mapFromJSON(recorder.Body.Bytes())
				Ω(mJSON["status"]).To(Equal("success"))
				results := sliceFromInterface(mJSON["data"])
				first := results[0].(map[string]interface{})
				last := results[1].(map[string]interface{})
				Ω(len(results)).To(Equal(2))
				Ω(first["_type"]).To(Equal(TypeResource))
				Ω(last["_type"]).To(Equal(TypeResource))
			})
		})
	})
})
