package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ProjectName = "robotclick"

type configuration struct {
	RootPath string
	Debug    bool
	Cron     string `json:"cron"`
	Snippets struct {
		Folder string    `json:"folder"`
		Files  []string  `json:"files"`
		Ss     []Snippet // the content of jsons loaded.
	}
}

type Snippet struct {
	Window struct {
		Name    string `json:"name"`
		BMPPath string `json:"bmp_path"`
	} `json:"window"`
	Trips []struct {
		Name   string `json:"name"`
		Action string `json:"action"`
		Double bool   `json:"double"`
		Msg    string `json:"msg"`
		Offset []int  `json:"offset"`
		Keys   []struct {
			Key  string   `json:"key"`
			Attr []string `json:"attr"`
		} `json:"keys"`
	}
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
	// load configs
	cf := filepath.Join(V.RootPath, "configs", "configs.json")
	f, err := os.ReadFile(cf)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, V); err != nil {
		return err
	}

	// load scripts
	for _, sf := range V.Snippets.Files {
		s := &Snippet{}
		scriptPath := filepath.Join(V.Snippets.Folder, sf)
		f, err := os.ReadFile(scriptPath)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(f, &s); err != nil {
			return err
		}
		V.Snippets.Ss = append(V.Snippets.Ss, *s)
	}
	return nil
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
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	ps := strings.Split(root, ProjectName)
	n := 0
	if len(ps) > 1 {
		n = strings.Count(ps[1], string(os.PathSeparator))
	}
	for i := 0; i < n; i++ {
		V.RootPath = filepath.Join("../", "./")
	}
	return nil

}
