package main

import (
	"log"

	"github.com/gr4c2-2000/gommand/internal/command"
	"github.com/gr4c2-2000/gommand/internal/daemon"
	"github.com/gr4c2-2000/gommand/internal/grpc"
)

//OLD WAY
// storage := command.InitStorage()
// ex := command.NewExecService(storage)
// ms := command.NewManageService(storage)
// serv, err := grpc.NewServer(context.Background(), ex, ms)
// if err != nil {
// 	log.Fatal(err)
// }
// err = serv.Wait()
// if err != nil {
// 	log.Fatal(err)
// }

func main() {
	storage := command.InitStorage()
	ex := command.NewExecService(storage)
	ms := command.NewManageService(storage)
	sw := grpc.NewDaemonServerWrapper(ex, ms)
	dm, err := daemon.NewDaemonService("gommandd", "Aliases replacment", sw)
	if err != nil {
		log.Fatal(err)
	}
	dm.Run()
}
