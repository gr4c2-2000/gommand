package main

import (
	"log"

	"github.com/gr4c2-2000/gommand/internal/command"
	"github.com/gr4c2-2000/gommand/internal/grpc"
)

func main() {
	storage := command.InitStorage()
	ex := command.NewExecService(storage)
	ms := command.NewManageService(storage)
	err := grpc.NewServer(ex, ms)
	if err != nil {
		log.Fatal(err)
	}
}
