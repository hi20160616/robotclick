package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var ProjectName = "robotclick"

type configuration struct {
	RootPath string
	Window   struct {
		Name string `json:"name"`
	} `json:"window"`
	Cron string `json:"cron"`
}

var V = &configuration{}

func setRootPath() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	V.RootPath = root
	if strings.Contains(os.Args[0], ".test") {
		return rootPath4Test()
	}
	return nil
}

func load() error {
	cf := filepath.Join(V.RootPath, "configs/configs.json")
	f, err := os.ReadFile(cf)
	if err != nil {
		return err
	}
	return json.Unmarshal(f, V)
}

func init() {
	if err := setRootPath(); err != nil {
		log.Printf("configs init error: %v", err)
	}
	if err := load(); err != nil {
		log.Printf("configs load error: %v", err)
	}
}

func rootPath4Test() error {
	ps := strings.Split(V.RootPath, ProjectName)
	n := 0
	if runtime.GOOS == "windows" {
		n = strings.Count(ps[1], "\\")
	} else {
		n = strings.Count(ps[1], "/")
	}
	for i := 0; i < n; i++ {
		V.RootPath = filepath.Join("../", V.RootPath)
	}
	return nil
}
