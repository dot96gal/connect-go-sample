package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	greetv1 "github.com/dot96gal/connect-go-sample/gen/greet/v1"
	"github.com/dot96gal/connect-go-sample/gen/greet/v1/greetv1connect"
)

func TestMain(t *testing.T) {
	t.Parallel()

	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)

	tests := []struct {
		scenario string
		name     string
		want     string
	}{
		{
			scenario: "hoge",
			name:     "hoge",
			want:     "Hello, hoge!",
		},
		{
			scenario: "fuga",
			name:     "fuga",
			want:     "Hello, fuga!",
		},
		{
			scenario: "bar",
			name:     "bar",
			want:     "Hello, bar!",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.scenario, func(t *testing.T) {
			t.Parallel()

			client := greetv1connect.NewGreetServiceClient(
				server.Client(),
				server.URL,
			)
			req := connect.NewRequest(&greetv1.GreetRequest{Name: tt.name})
			res, err := client.Greet(context.Background(), req)
			if err != nil {
				t.Error(err)
			}

			got := res.Msg.Greeting
			if got != tt.want {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}
