.PHONY: buf-lint
buf-lint:
	buf lint

.PHONY: buf-gen
buf-gen:
	buf generate

.PHONY: run-server
run-server:
	go run ./cmd/server/...

.PHONY: run-client
run-client:
	go run ./cmd/client/...

.PHONY: lint-client
lint-client:
	golangci-lint run ./cmd/client/...

.PHONY: lint-server
lint-server:
	golangci-lint run ./cmd/server/...

.PHONY: test-server
test-server:
	go test -race ./cmd/server/...

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
