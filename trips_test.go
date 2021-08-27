package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hi20160616/robotclick/configs"
)

func TestTreatSnippet(t *testing.T) {
	tc := NewTrip()
	// test test1.json
	if err := tc.treatSnippet(&configs.V.Snippets.Ss[0]); err != nil {
		t.Error(err)
	}
	// test test2.json
	if err := tc.treatSnippet(&configs.V.Snippets.Ss[1]); err != nil {
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
