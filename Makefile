.PHONY: buf-lint
buf-lint:
	buf lint

.PHONY: buf-gen
buf-gen:
	buf generate

.PHONY: run-server
run-server:
	go run ./cmd/server/main.go

.PHONY: run-client
run-client:
	go run ./cmd/client/main.go

.PHONY: greet-curl
greet-curl:
	curl \
		--header "Content-Type: application/json" \
		--data '{"name": "curl"}' \
		http://localhost:8080/greet.v1.GreetService/Greet

.PHONY: greet-grpc
greet-grpc: SHELL:=/bin/bash
greet-grpc:
	grpcurl \
		-protoset <(buf build -o -) \
		-plaintext \
		-d '{"name": "grpcurl"}' \
		localhost:8080 greet.v1.GreetService/Greet
