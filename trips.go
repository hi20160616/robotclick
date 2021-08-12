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
			p = &Pos{X: p.X + trip.Offset[0], Y: p.Y + trip.Offset[1]}
			click(p, trip.Double)
		case "input":
			robotgo.TypeStr(trip.Msg, 1)
		}
		// robotgo.MilliSleep(500)
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
	defer robotgo.FreeBitmap(cb)
	// s := robotgo.TostringBitmap(cb)
	// fmt.Println(s)
	fx, fy := robotgo.FindBitmap(cb, nil, 0.1) // last arg is tolerance
	if fx < 0 || fy < 0 {
		return nil, fmt.Errorf("find none: (%d, %d)", fx, fy)
	}
	return &Pos{fx, fy}, nil
}

func click(p *Pos, double bool) {
	robotgo.MoveMouseSmooth(p.X, p.Y, 1.0, 0.3)
	robotgo.Click()
}
