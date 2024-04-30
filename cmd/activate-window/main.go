package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	activewindow "github.com/gr4c2-2000/gommand/internal/active-window"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args)
		fmt.Println("NO ARGS")
		os.Exit(1)
	}
	st, err := activewindow.InitStorage("./storage.json")
	if err != nil {
		log.Fatal(err)
	}
	x11, err := activewindow.NewX11Client()
	if err != nil {
		log.Fatal(err)
	}
	defer x11.Close()
	service := activewindow.InitService(st, x11)
	err = service.Next(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}

}
