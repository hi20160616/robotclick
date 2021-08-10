package main

import (
	"C"
)
import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"

	"github.com/go-vgo/robotgo"
	"github.com/hi20160616/gears"
	"github.com/hi20160616/robotclick/configs"
)

type Trip struct {
}

type Pos struct {
	X, Y int
}

func NewTrip() *Trip {
	return &Trip{}
}

func (t *Trip) working() error {
	// handle the app window
	if err := robotgo.ActiveName(configs.V.Window.Name); err != nil {
		return err
	}
	// loop the trips in configs.json
	for _, trip := range configs.V.Trips {
		switch trip.Action {
		case "click":
			// get position
			p, err := getPos(trip.Name)
			if err != nil {
				return err
			}
			click(p, trip.Double)
		case "input":
			robotgo.TypeStrDelay(trip.Msg, 1)
			// robotgo.KeyTap("enter")
		}
	}
	return nil
}

func getPos(name string) (*Pos, error) {
	bPath := filepath.Join(configs.V.RootPath, "configs", configs.V.Window.BMPPath, name)
	if !gears.Exists(bPath) {
		return nil, fmt.Errorf("no bitmap find out: %s", bPath)
	}
	return findBitmap(bPath)
}

func findBitmap(imgsrc string) (*Pos, error) {
	cb := robotgo.OpenBitmap(imgsrc)
	fx, fy := robotgo.FindBitmap(cb)
	if fx < 0 || fy < 0 {
		return nil, fmt.Errorf("find none: (%d, %d)", fx, fy)
	}
	return &Pos{fx, fy}, nil
}

func click(p *Pos, double bool) {
	robotgo.MouseToggle("up")
	robotgo.MoveClick(p.X, p.Y, "left", double)
}
