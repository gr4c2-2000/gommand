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

func (ser *Service) Set(appName string) error {
	id, err := ser.nextWindow(appName)

	if err != nil {
		log.Default().Printf("Find next window Error: %v", err)
		return err
	}

	err = ser.X11Client.SetWindow(id)
	if err != nil {
		log.Default().Printf("Set Window Error: %v", err)
		return err
	}
	ser.storage.Replace(appName, uint32(id))
	return nil
}

func (ser *Service) nextWindow(appName string) (xproto.Window, error) {
	apps, err := ser.X11Client.Find(appName)
	if err != nil {
		return 0, err
	}
	if len(apps) == 0 {
		return 0, errors.New("NO RUNNING APPS")
	}

	item := ser.storage.Find(appName)

	if item == nil {
		return apps[0], nil
	}

	activeWindow, err := ser.X11Client.ActiveWindow()

	if err == nil && isRunning(int32(item.LastId), apps) && item.LastId != activeWindow {
		return xproto.Window(item.LastId), nil
	}

	for i := 0; i < len(apps); i++ {
		if uint32(apps[i]) == item.LastId && i < len(apps)-1 {
			return apps[i+1], nil
		}
	}

	return apps[0], nil
}

func isRunning(lastId int32, apps []xproto.Window) bool {
	for _, id := range apps {
		if lastId == int32(id) {
			return true
		}
	}
	return false
}
