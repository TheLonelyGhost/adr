package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type AdrConfig struct {
	DecisionsDir string `json:"decisions_directory"`
	DateFormat   string `json:"date_format"`
}

type AdrIndex uint16

type AdrData struct {
	Config    *AdrConfig
	Template  string
	Index     AdrIndex
	BaseDir   string
	ConfigDir string
}

func (a *AdrData) Init() (err error) {
	err = os.MkdirAll(a.ConfigDir, 0o755)
	if err != nil {
		return
	}

	//
	// Write the index changes to disk
	err = os.WriteFile(filepath.Join(a.ConfigDir, "index"), []byte(strconv.FormatUint(uint64(a.Index), 10)), 0o644)
	if err != nil {
		return
	}

	// Write the config file changes
	data, err := json.MarshalIndent(a.Config, "", "  ")
	if err != nil {
		return
	}
	err = os.WriteFile(filepath.Join(a.ConfigDir, "config.json"), []byte(string(data)+"\n"), 0o644)
	if err != nil {
		return
	}

	// Write the default template file contents
	err = os.WriteFile(filepath.Join(a.ConfigDir, "template.md"), []byte(a.Template), 0o644)
	if err != nil {
		return
	}
	return
}

func (a *AdrData) Write() (err error) {
	// Write the index changes to disk
	err = os.WriteFile(filepath.Join(a.ConfigDir, "index"), []byte(fmt.Sprintf("%d\n", a.Index)), 0o644)
	return
}

func (a *AdrData) DecisionsDir() string {
	if filepath.IsAbs(a.Config.DecisionsDir) {
		return a.Config.DecisionsDir
	}
	return filepath.Join(a.BaseDir, a.Config.DecisionsDir)
}
