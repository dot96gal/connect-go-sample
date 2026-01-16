FROM golang:1.26rc2-bookworm

WORKDIR /workspace

# install dev tools
RUN apt-get update \
  && apt-get install -y \
  build-essential \
  && apt-get clean \
  && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

# setup go module
COPY go.mod .
COPY go.sum .
RUN go mod download

# install go tools
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.6
