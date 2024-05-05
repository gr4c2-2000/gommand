
ARG GO_VERSION=1.22.1

# Use golang base image
FROM golang:${GO_VERSION}

# Define Variables 
ENV GEN_GO_VERSION=1.33.0
ENV GEN_GO_GRPC_VERSION=1.3.0
ENV PROTOC_LINK=https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip
ENV PROTOC_FILE=protoc-21.12-linux-x86_64.zip

# Install protobuf tools, zip, protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GEN_GO_VERSION} && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${GEN_GO_GRPC_VERSION} && \
    apt-get update -y && apt-get install -y protobuf-compiler --no-install-recommends && \
    apt-get install -y git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app

# Disable VCS stamping
ENV GOFLAGS=-buildvcs=false
# Define environment variables
ENV PATH="$PATH:/go/bin"