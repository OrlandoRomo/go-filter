package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var effect *Effect

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
		NewGraySubCommand(),
		NewNegativeSubCommand(),
		NewRedSubCommand(),
		NewBlueSubCommand(),
		NewGreenSubCommand(),
		NewMirrorSubCommand(),
		NewSepiaSubCommand(),
		NewSketchSubCommand(),
		NewSharpSubCommand(),
	}
	return filterCommand
}

func setOutputFlag() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}
