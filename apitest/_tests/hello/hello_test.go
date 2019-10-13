package hello_test

import (
	"net/http"

	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
)

var _ = gk.Describe("Hello", func() {
	gk.Describe("Categorizing book length", func() {
		gk.Context("With more than 300 pages", func() {
			gk.It("should be a novel", func() {
				gm.Expect("NOVEL").To(gm.Equal("NOVEL"))
			})
		})
	})

	gk.Describe("Check Request to Hello API", func() {
		gk.Context("GET REQUEST", func() {
			gk.It("/hello", func() {
				request, _ := http.NewRequest("GET", "/hello", nil)
				request.RequestURI = "/hello"

				response := testHTTPServer.Response(request)
				gm.Expect(response.HeaderMap.Get("Content-Type")).To(gm.Equal("application/json"))
			})
		})
	})
})
