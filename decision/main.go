package decision

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/thelonelyghost/adr/config"

	"github.com/Masterminds/sprig/v3"
)

type DecisionStatus string

const (
	PROPOSED   DecisionStatus = "Proposed"
	ACCEPTED   DecisionStatus = "Accepted"
	DEPRECATED DecisionStatus = "Deprecated"
	SUPERSEDED DecisionStatus = "Superseded"
)

var slugChars = regexp.MustCompile("[^A-Za-z0-9_.-]+")

type Decision struct {
	Index  config.AdrIndex
	Title  string
	Status DecisionStatus
}

func (d Decision) Filename() string {
	return fmt.Sprintf("%d-%s.md", d.Index, slugChars.ReplaceAllString(d.Title, "-"))
}

func New(cfg *config.AdrData, title string) (adr Decision, err error) {
	adr = Decision{
		Title:  strings.Trim(title, "\n \t"),
		Index:  cfg.Index,
		Status: PROPOSED,
	}

	f, err := os.Create(filepath.Join(cfg.DecisionsDir(), adr.Filename()))
	if err != nil {
		return
	}
	defer f.Close()

	log.Printf("Template source: %s\n", cfg.TemplatePath())

	tpl, err := template.New(filepath.Base(cfg.TemplatePath())).Funcs(sprig.FuncMap()).ParseFiles(cfg.TemplatePath())
	if err != nil {
		return
	}

	err = tpl.Execute(f, adr)
	return
}
