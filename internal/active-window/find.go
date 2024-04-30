package activewindow

import (
	"fmt"
	"log"
	"strings"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
)

type X11Client struct {
	con *xgbutil.XUtil
}

func NewX11Client() (*X11Client, error) {
	x11 := X11Client{}
	var err error
	x11.con, err = xgbutil.NewConn()
	if err != nil {
		return nil, err
	}
	return &x11, nil
}

func (x11 *X11Client) Close() {
	x11.con.Conn().Close()
}

func (x11 *X11Client) Find(search string) ([]xproto.Window, error) {

	clients, err := ewmh.ClientListGet(x11.con)
	if err != nil {
		return nil, err
	}
	fmt.Println(clients)
	result := make([]xproto.Window, 0)
	for _, wid := range clients {

		name, err := ewmh.WmNameGet(x11.con, wid)
		if err != nil { // not a fatal error
			log.Println(err)
		}
		if strings.Contains(strings.ToLower(name), strings.ToLower(search)) {
			fmt.Printf("name : %v, search : %v\n", name, search)
			result = append(result, wid)
		}
	}
	return result, nil
}

func (x11 *X11Client) SetWindow(wid xproto.Window) error {
	appDesktop, err := ewmh.WmDesktopGet(x11.con, wid)
	if err != nil {
		return err
	}
	currentDesktop, err := ewmh.CurrentDesktopGet(x11.con)
	if err != nil {
		return err
	}
	if appDesktop != currentDesktop {
		ewmh.CurrentDesktopReq(x11.con, int(appDesktop))
	}

	err = ewmh.ActiveWindowReq(x11.con, wid)
	if err != nil {
		return err
	}
	return nil
}
func (x11 *X11Client) ActiveWindow() (uint32, error) {
	wid, err := ewmh.ActiveWindowGet(x11.con)
	return uint32(wid), err
}
