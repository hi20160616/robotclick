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
	// Cron      string  `json:"cron"`
	Tolerance float64 `json:"tolerance"`
	Snippets  struct {
		Folder string    `json:"folder"`
		Files  []string  `json:"files"`
		Ss     []Snippet // the content of jsons loaded.
	}
}

type Snippet struct {
	Cron     string `json:"cron"`
	FileName string
	Window   struct {
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
		DaysAgo int    `json:"days_ago"`
		Layout  string `json:"layout"`
		Delay   int    `json:"delay"` // million seconds
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
		s, err := loadSnippet(sf)
		if err != nil {
			return err
		}
		V.Snippets.Ss = append(V.Snippets.Ss, *s)
	}
	return nil
}

func loadSnippet(snippetName string) (*Snippet, error) {
	s := &Snippet{FileName: snippetName}
	snippetPath := filepath.Join(filepath.Join(V.RootPath, "configs"), V.Snippets.Folder, s.FileName)
	f, err := os.ReadFile(snippetPath)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(f, &s); err != nil {
		return nil, err
	}
	return s, nil
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
