package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	greetv1 "github.com/dot96gal/connect-go-sample/gen/greet/v1"
	"github.com/dot96gal/connect-go-sample/gen/greet/v1/greetv1connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		// connect.WithGRPC(),
		// connect.WithGRPCWeb,
	)

	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "connect-go-sample-client"}),
	)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.Msg.Greeting)
}
