package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Run() (err error) {
	app := cli.NewApp()
	app.Name = "adr"
	app.Usage = "Work with Architecture Decision Records (ADRs)"
	app.Version = "0.2.0"
	app.Flags = []cli.Flag{}

	app.Commands = []*cli.Command{
		&cmdNew,
		&cmdInit,
	}

	err = app.Run(os.Args)
	return
}
