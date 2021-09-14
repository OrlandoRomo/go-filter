package main

import (
	"os"
	"sync"

	"github.com/urfave/cli/v2"
)

var (
	effect *Effect
	wg     sync.WaitGroup
)

func NewFilterCommand() *cli.Command {
	filterCommand := &cli.Command{
		Name:  "filter",
		Usage: "filter will apply the given filter name",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "output",
				Usage: "path where the output result will be placed",
				Value: setOutputFlag(),
			},
		},
	}
	filterCommand.Subcommands = []*cli.Command{
		NewListCommand(),
		NewGreySubCommand(),
		NewNegativeSubCommand(),
		NewRedSubCommand(),
		NewBlueSubCommand(),
		NewGreenSubCommand(),
		NewMirrorSubCommand(),
	}
	return filterCommand
}

func setOutputFlag() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}
