package main

import (
	"C"
)
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-vgo/robotgo"
	"github.com/hi20160616/gears"
	"github.com/hi20160616/robotclick/configs"
)

type Trip struct {
	BMPathes []string
	Ps       []*Pos
}

type Pos struct {
	X, Y int
}

func NewTrip() *Trip {
	return &Trip{}
}

func (t *Trip) loadBMs() (*Trip, error) {
	files, err := os.ReadDir(configs.V.Window.BMPath)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.Type().IsRegular() {
			t.BMPathes = append(t.BMPathes,
				filepath.Join(configs.V.Window.BMPath, f.Name()))
		}
	}
	return t, nil
}

func (t *Trip) working() error {
	// load bitmaps, add filepath to t
	t, err := t.loadBMs()
	if err != nil {
		return err
	}
	// loop the trips in configs.json
	for _, trip := range configs.V.Trips {
		bPath := filepath.Join(configs.V.Window.BMPath, trip.Name)
		if !gears.Exists(bPath) {
			return fmt.Errorf("no bitmap find out: %s", bPath)
		}
		p, err := findBitmap(bPath)
		if err != nil {
			return err
		}
		switch trip.Action {
		case "click":
			click(p, trip.Double)
		case "input":
			click(p, trip.Double)
			robotgo.TypeStrDelay(trip.Msg, 1)
			robotgo.KeyTap("enter")
		}
	}
	return nil
}

func findBitmap(imgsrc string) (*Pos, error) {
	bm := robotgo.OpenBitmap(imgsrc)
	fx, fy := robotgo.FindBitmap(bm)
	if fx < 0 || fy < 0 {
		return nil, fmt.Errorf("find none: (%d, %d)", fx, fy)
	}
	return &Pos{fx, fy}, nil
}

func click(p *Pos, double bool) {
	robotgo.MouseToggle("up")
	robotgo.MoveClick(p.X, p.Y, "left", double)
}

func typeMsg(p *Pos, msg string) {
	click(p, false)
	robotgo.TypeStrDelay(msg, 1)
	robotgo.KeyTap("enter")
}
