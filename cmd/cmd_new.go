package cmd

import (
	"errors"
	"strings"

	"github.com/fatih/color"
	"github.com/thelonelyghost/adr/config"
	"github.com/thelonelyghost/adr/decision"

	// "github.com/thelonelyghost/adr/decision"

	"github.com/urfave/cli/v2"
)

var cmdNew = cli.Command{
	Name: "new",
	Aliases: []string{
		"n",
		"c",
		"create",
	},
	Usage: "Create a new Architecture Decision Record",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "chdir",
			Aliases: []string{"C"},
			Value:   ".",
			Usage:   "Base `DIR` where to search for the closest ADR repository",
		},
		&cli.BoolFlag{
			Name:    "global",
			Aliases: []string{"g"},
			Usage:   "Create an ADR in your global location, not the closest repository to current directory.",
		},
	},
	Action: cmdNewFunc,
}

func cmdNewFunc(c *cli.Context) (err error) {
	var (
		title   string
		adrPath string
	)

	if c.Bool("global") {
		adrPath = config.GlobalConfigDir
	} else {
		baseDir := c.String("chdir")
		if baseDir != "" {
			baseDir = "."
		}
		adrPath = config.FindConfig(baseDir)
	}

	if c.NArg() == 0 {
		err = errors.New("must provide a title for the new Architecture Decision Record")
		return
	} else {
		title = strings.Join(c.Args().Slice(), " ")
	}

	cfg := config.Load(adrPath)
	if cfg == nil {
		if adrPath == config.GlobalConfigDir {
			err = errors.New("ADR repository must be initialized first: `adr init --global`")
		} else {
			err = errors.New("ADR repository must be initialized first: `adr init .`")
		}
		return
	}

	cfg.Index++

	adr, err := decision.New(cfg, title)
	if err != nil {
		return
	}
	color.Green("ADR #%s created: %s", adr.Index, adr.Filename())

	err = cfg.Write()
	return
}
