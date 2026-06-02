package main

import (
	"log"
	"net/http"

	"github.com/dot96gal/connect-go-sample/gen/greet/v1/greetv1connect"
)

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	// use h2c so we can serve HTTP/2 without tls
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)

	srv := &http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: protocols,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
