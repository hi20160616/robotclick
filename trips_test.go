package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hi20160616/robotclick/configs"
)

func TestWorking(t *testing.T) {
	tc := NewTrip()
	if err := tc.working(); err != nil {
		t.Error(err)
	}
}

func TestFindBitmap(t *testing.T) {
	tc := filepath.Join(configs.V.RootPath, "configs", configs.V.Window.BMPPath, "5.png")
	p, err := findBitmap(tc)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p)
}
