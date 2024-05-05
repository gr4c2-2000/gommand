package command

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type Storage struct {
	commandMap map[string]Command
	path       string
}

func InitStorage() *Storage {
	s := &Storage{commandMap: make(map[string]Command)}
	s.path = "/home/gr4c2/go/src/github.com/gr4c2-2000/gommand/storage/"
	s.Discover(s.path)
	s.Watch()
	return s

}

func (s *Storage) FindCommand(shellCall string) (*Command, error) {
	exec, err := parseCall(shellCall)
	if err != nil {
		return nil, err
	}
	command, ok := s.commandMap[strings.TrimSpace(exec.CommandStr)]
	if !ok {
		return nil, errors.New("NO COMMAND")
	}
	command.exec = *exec
	return &command, nil

}

func (s *Storage) GetAll() map[string]Command {
	return s.commandMap
}

func (s *Storage) Discover(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		err := s.LoadFile(dir + file.Name())
		if err != nil {
			log.Println("error:", err)
		}
	}
}

func (s *Storage) Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				err := s.LoadFile(event.Name)
				if err != nil {
					log.Println("error:", err)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(s.path)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Storage) LoadFile(dir string) error {
	file_name := filepath.Base(dir)
	if path.Ext(file_name) != ".json" {
		log.Default().Println("Not a json file")
		return nil
	}
	c := Command{}
	bdata, err := os.ReadFile(dir)
	if err != nil {
		log.Default().Println(err)
		return err
	}
	err = json.Unmarshal(bdata, &c)
	if err != nil {
		return err
	}
	s.commandMap[strings.TrimSuffix(file_name, filepath.Ext(file_name))] = c
	return nil
}

func parseCall(s string) (*Exec, error) {
	if s == "" {
		return nil, errors.New("cbdbhc")
	}
	e := Exec{}

	sliceCall := SplitSpecial(s)
	tmpArgs := strings.Split(sliceCall[0], "-")
	for key, val := range tmpArgs {
		if val[0] == '\'' && val[len(val)-1] == '\'' {
			tmpArgs[key] = val[1 : len(val)-1]
		}
		if val[0] == '"' && val[len(val)-1] == '"' {
			tmpArgs[key] = val[1 : len(val)-1]
		}
	}
	e.CommandStr = tmpArgs[0]
	e.TmpArgs = tmpArgs[1:]

	if len(sliceCall) == 2 {
		e.CmdArgs = sliceCall[1]
	}
	return &e, nil
}
