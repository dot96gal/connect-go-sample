package main

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	greetv1 "github.com/dot96gal/connect-go-sample/gen/greet/v1"
)

func TestGreet(t *testing.T) {
	t.Parallel()

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
		t.Run(tt.scenario, func(t *testing.T) {
			t.Parallel()

			greeter := &GreetServer{}
			req := connect.NewRequest(&greetv1.GreetRequest{Name: tt.name})
			res, err := greeter.Greet(context.Background(), req)
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
