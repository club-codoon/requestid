package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
)

var (
	extractTests = []struct {
		Options  *Options
		Header   string
		Expected RequestId
	}{
		{&Options{}, "1234", "1234"},
		{&Options{Generate: false}, "", ""},
	}
)

func Test_Extract(t *testing.T) {
	for _, test := range extractTests {
		recorder := httptest.NewRecorder()
		m := martini.Classic()
		m.Use(Middleware(test.Options))
		m.Get("/", func(requestId RequestId) {
			if requestId != test.Expected {
				t.Fatalf("Expected %v, got %v", test.Expected, requestId)
			}
		})
		r, _ := http.NewRequest("GET", "/", nil)
		r.Header.Set("X-Request-Id", test.Header)
		m.ServeHTTP(recorder, r)
	}
}
