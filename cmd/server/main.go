package main

import (
	"net/http"

	"github.com/dot96gal/connect-go-sample/gen/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		// use h2c so we can serve HTTP/2 without tls
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
