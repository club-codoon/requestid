# Martini Request Id

This is a middleware for martini that will extract (or generate) the request id
from the `X-Request-Id` header, and make it available to be injected.

## Usage


```go
m := martini.Classic()
m.Use(requestid.Middleware())
m.Get("/", func(requestId requestid.RequestId) {
  return 200, "The request id is: " + requestId
})
```
