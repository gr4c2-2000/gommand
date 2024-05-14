package main

import (
	"log"
	"os"
	"strings"

	activewindow "github.com/gr4c2-2000/gommand/internal/active-window"
	"github.com/gr4c2-2000/gommand/internal/command"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("NO ARGS")
	}
	st, err := activewindow.InitStorage("./.gommand-activate-window.storage")
	if err != nil {
		log.Fatal(err)
	}
	x11, err := activewindow.NewX11Client()
	if err != nil {
		log.Fatal(err)
	}
	defer x11.Close()
	service := activewindow.InitService(st, x11)
	err = service.Set(strings.TrimSpace(os.Args[1]))
	if err == nil {
		return
	}

	if err != nil && len(os.Args) < 3 {
		log.Fatal(err)
	}

	dir, _ := os.Getwd()

	err = command.ExecSync("bash", os.Args[2], os.Args[2:], dir)
	if err != nil {
		log.Fatal(err)
	}

}
