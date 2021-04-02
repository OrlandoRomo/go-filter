package cmd

import (
	_ "image/jpeg"
	"os"

	"github.com/OrlandoRomo/imgfltr/cmd/fltrs"
	"github.com/urfave/cli/v2"
)

const (
// For linux and Mac home directory
)

var supportedExtensions map[string]bool

func init() {
	supportedExtensions = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".webp": true,
	}
}

func NewSetCommand() *cli.Command {
	return &cli.Command{
		Name:  "set",
		Usage: "set requires a filter name",
		Subcommands: []*cli.Command{
			fltrs.NewGreySubCommand(),
			fltrs.NewSepiaSubCommand(),
			fltrs.NewNegativeSubCommand(),
			fltrs.NewMirrorSubCommand(),
			fltrs.NewSketchSubCommand(),
			fltrs.NewBlurSubCommand(),
			fltrs.NewSharpSubCommand(),
			fltrs.NewRedSubCommand(),
			fltrs.NewGreenSubCommand(),
			fltrs.NewBlueSubCommand(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "output",
				Usage: "path where the output result will be placed",
				Value: setOutputFlag(),
			},
		},
	}
}

func setOutputFlag() string {
	home, _ := os.UserHomeDir()
	return home
}
