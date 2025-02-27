package cmd

import (
	"path/filepath"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/thelonelyghost/adr/config"
)

var cmdInit = cli.Command{
	Name:        "init",
	Aliases:     []string{"i"},
	Usage:       "Initializes the ADR configuration",
	UsageText:   "adr init .",
	Description: "Initializes a new Architecture Decision Record (ADR) repository",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "global",
			Aliases: []string{"g"},
			Usage:   "Initialize a centrally-located ADR repository",
		},
		&cli.StringFlag{
			Name:    "data",
			Aliases: []string{"d"},
			Value:   "decisions",
			Usage:   "Relative path to a directory where ADRs will be managed. Will be created if it does not yet exist.",
		},
	},
	Action: cmdInitFunc,
}

func cmdInitFunc(c *cli.Context) (err error) {
	var cfg *config.AdrData
	isGlobal := false

	if c.Bool("global") {
		isGlobal = true
		cfg = config.Load(config.GlobalConfigDir)
	} else if c.NArg() > 0 {
		cfg = config.Load(filepath.Join(c.Args().Get(0), ".adr"))
	} else {
		cfg = config.Load(filepath.Join(".", ".adr"))
	}

	decisionsPath := c.String("data")
	if decisionsPath != "" {
		if isGlobal {
			cfg.Config.DecisionsDir, _ = filepath.Abs(decisionsPath)
		} else {
			cfg.Config.DecisionsDir = decisionsPath
		}
	}
	err = cfg.Init()

	code := color.New(color.Bold, color.FgBlack, color.BgGreen).SprintfFunc()

	color.Green("ADR repository has been initialized.")
	color.Green("Create your first record: `%s`", code("adr new Hello World"))
	return
}
