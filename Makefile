DOCKER_IMAGE_NAME := proto-gen
DOCKERFILE_PATH := .docker/proto-gen/
BUILD_DOCKER := docker build --platform=linux/amd64 -t $(DOCKER_IMAGE_NAME) $(DOCKERFILE_PATH)
GOMMOND_COMMAND := protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./internal/proto/gommand.proto
EXEC_IN_DOCKER := docker run --rm -v $(shell pwd):/app $(DOCKER_IMAGE_NAME)
SERVER_MAIN := ./cmd/gommandd/
CMD_MAIN := ./cmd/gmd/
AW_MAIN := ./cmd/activate-window/
GO_BUILD := go build -o ./bin


grpc: 
	$(BUILD_DOCKER) && $(EXEC_IN_DOCKER) $(GOMMOND_COMMAND)

build_server: 
	$(BUILD_DOCKER) && $(EXEC_IN_DOCKER) $(GO_BUILD) $(SERVER_MAIN)

build_command: 
	$(BUILD_DOCKER) && $(EXEC_IN_DOCKER) $(GO_BUILD) $(CMD_MAIN)

build_activate_window:
	$(BUILD_DOCKER) && $(EXEC_IN_DOCKER) $(GO_BUILD) $(AW_MAIN)

go_build_server: 
	$(GO_BUILD) $(SERVER_MAIN)

go_build_gmd: 
	$(GO_BUILD) $(CMD_MAIN)


