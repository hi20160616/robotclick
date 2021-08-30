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
	"time"

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

func (t *Trip) TreatSnippets() error {
	for _, s := range configs.V.Snippets.Ss {
		if err := t.treatSnippet(&s); err != nil {
			return err
		}
	}
	return nil
}

func (t *Trip) treatSnippet(s *configs.Snippet) error {
	// handle the app window
	if err := robotgo.ActiveName(s.Window.Name); err != nil {
		return err
	}
	// loop the trips in configs.json
	for _, trip := range s.Trips {
		switch trip.Action {
		case "click":
			// get position
			if trip.Name != "" {
				p, err := getPos(s, trip.Name)
				if err != nil {
					return err
				}
				p = &Pos{X: p.X + trip.Offset[0], Y: p.Y + trip.Offset[1]}
				click(p, trip.Double)
			} else {
				x, y := robotgo.GetMousePos()
				p := &Pos{X: x + trip.Offset[0], Y: y + trip.Offset[1]}
				click(p, trip.Double)
			}
		case "input-days-ago":
			t := time.Now().AddDate(0, 0, -trip.DaysAgo)
			layout := "2006-01-02"
			if trip.Layout != "" {
				layout = trip.Layout
			}
			robotgo.TypeStr(t.Format(layout), 1)
		case "input":
			robotgo.TypeStr(trip.Msg, 1)
		case "type":
			for _, k := range trip.Keys {
				if k.Attr != nil {
					robotgo.KeyTap(k.Key, k.Attr)
				}
				robotgo.KeyTap(k.Key)
				robotgo.MilliSleep(500)
			}
		}
		robotgo.MilliSleep(trip.Delay)
	}
	return nil
}

func getPos(s *configs.Snippet, name string) (*Pos, error) {
	bPath := filepath.Join(configs.V.RootPath,
		"configs",
		configs.V.Snippets.Folder,
		s.Window.BMPPath, name)
	if !gears.Exists(bPath) {
		return nil, fmt.Errorf("no bitmap find out: %s", bPath)
	}
	return findBitmap(bPath)
}

func findBitmap(imgsrc string) (*Pos, error) {
	cb := robotgo.OpenBitmap(imgsrc)
	defer robotgo.FreeBitmap(cb)
	fx, fy := robotgo.FindBitmap(cb, nil, configs.V.Tolerance) // last arg is tolerance
	if fx < 0 || fy < 0 {
		return nil, fmt.Errorf("find none: (%d, %d)", fx, fy)
	}
	return &Pos{fx, fy}, nil
}

func click(p *Pos, double bool) {
	robotgo.MoveMouseSmooth(p.X, p.Y, 1.0, 0.3)
	robotgo.Click()
}
