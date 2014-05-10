// The requestid package is a partini middleware that will parse
// the X-Request-Id header from a request and make it injectible.
package requestid

import (
	"net/http"

	"github.com/go-martini/martini"
)

type RequestId string

// A middleware to extract the request id.
func Middleware() martini.Handler {
	return func(c martini.Context, req *http.Request) {
		requestId := req.Header.Get("X-Request-Id")

		c.Map(RequestId(requestId))
		c.Next()
	}
}
