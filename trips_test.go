package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hi20160616/robotclick/configs"
)

func TestTreatSnippet(t *testing.T) {
	// test test1.json
	tc1 := NewTrip(&configs.V.Snippets.Ss[0])
	if err := tc1.treatSnippet(); err != nil {
		t.Error(err)
	}
	// test test2.json
	tc2 := NewTrip(&configs.V.Snippets.Ss[1])
	if err := tc2.treatSnippet(); err != nil {
		t.Error(err)
	}
}

func TestFindBitmap(t *testing.T) {
	tc := filepath.Join(configs.V.RootPath, "configs",
		configs.V.Snippets.Ss[0].Window.BMPPath, "5.png")
	p, err := findBitmap(tc)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p)
}
