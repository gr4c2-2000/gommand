package activewindow

import (
	"errors"
	"log"

	"github.com/BurntSushi/xgb/xproto"
)

type Service struct {
	storage   *Storage
	X11Client *X11Client
}

func InitService(storage *Storage, x11Client *X11Client) *Service {
	return &Service{storage: storage, X11Client: x11Client}
}

func (ser *Service) Next(appName string) error {
	apps, err := ser.X11Client.Find(appName)
	if err != nil {
		return err
	}
	if len(apps) == 0 {
		return errors.New("NO RUNNING APPS")
	}
	var id xproto.Window

	defer func() {
		err := ser.X11Client.SetWindow(id)
		if err != nil {
			log.Default().Printf("Set Window Error: %v", err)
			return
		}
		ser.storage.Replace(appName, uint32(id))
	}()

	item := ser.storage.Find(appName)
	if item == nil {
		id = apps[0]
		return nil
	}
	//todo : add skip rotation if active windows is != lastId
	// id, err := ActiveWindow()

	// if item.LastId ==
	for i := 0; i < len(apps); i++ {
		if uint32(apps[i]) == item.LastId && i < len(apps)-1 {
			id = apps[i+1]
			return nil
		}
	}
	id = apps[0]
	return nil
}
