package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
)

func Test_RequestId_Middleware(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.Classic()
	m.Use(Middleware())
	m.Get("/", func(requestId RequestId) {
		if requestId != RequestId("1234") {
			t.Fatal("Expected the request id")
		}
	})

	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("X-Request-Id", "1234")
	m.ServeHTTP(recorder, r)
}
