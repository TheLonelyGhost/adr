package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	userHomeDir, _  = os.UserHomeDir()
	GlobalConfigDir = filepath.Join(userHomeDir, ".adr")
	defaultTemplate = `# {{ .Index }}. {{ title .Title }}

**Date:** {{ now | date "2006-01-02" }}

## Status

{{.Status}}

## Context

...

## Decision

...

## Consequences

...
`
)

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func isValidConfigDir(targetPath string) bool {
	if !pathExists(filepath.Join(targetPath, "index")) {
		return false
	}

	contents, err := os.ReadFile(filepath.Join(targetPath, "index"))
	if err != nil {
		return false
	}
	if _, err := strconv.ParseUint(strings.Trim(string(contents), "\t \n"), 10, 16); err != nil {
		return false
	}

	return true
}

// FindConfig starts with a given path on the filesystem and
// continues walking up the parent directories until it either
// finds a directory containing `.adr/` in it, or it reaches
// the root directory. If it reaches the root directory, it will
// report the user-global location, `~/.adr/`, as the location
// of the configs.
// The return value will always be the closest ancestor's `.adr/`
// directory, or `~/.adr/` (regardless of if it exists) if there
// is no ancestor containing `.adr/`.
func FindConfig(startingPath string) (configPath string) {
	// starting with current directory, walk each parent directory until hit root
	cursor, _ := filepath.Abs(startingPath)
	for {
		if cursor == filepath.Dir(cursor) {
			// if cursor has reached the root directory, pull from global location for user
			configPath = GlobalConfigDir
			return
		}

		if isValidConfigDir(filepath.Join(cursor, ".adr")) {
			configPath = filepath.Join(cursor, ".adr")
			return
		}

		cursor = filepath.Dir(cursor)
	}
}

func Load(configDirPath string) (data *AdrData) {
	// Set some defaults
	data = &AdrData{
		Config: &AdrConfig{
			DecisionsDir: "decisions",
			DateFormat:   "2006-01-02 15:04:05",
		},
		Template:  string(defaultTemplate),
		Index:     AdrIndex(0),
		ConfigDir: configDirPath,
		BaseDir:   filepath.Dir(configDirPath),
	}
	// if _, err := os.Stat(configDirPath); os.IsNotExist(err) {
	// 	return nil
	// }

	if contents, err := os.ReadFile(filepath.Join(configDirPath, "config.json")); err == nil {
		json.Unmarshal(contents, data.Config)
	}
	if contents, err := os.ReadFile(filepath.Join(configDirPath, "index")); err == nil {
		index, err := strconv.ParseUint(strings.Trim(string(contents), "\t \n"), 10, 16)
		if err != nil {
			panic(err)
		}
		data.Index = AdrIndex(index)
	}
	if contents, err := os.ReadFile(filepath.Join(configDirPath, "template.md")); err == nil {
		data.Template = string(contents)
	}

	return
}
