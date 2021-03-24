package fltrs

import (
	"github.com/urfave/cli/v2"
)

const (
	RedWaveLength   float64 = 0.21
	GreenWaveLength float64 = 0.72
	BlueWaveLength  float64 = 0.07
	RGBA            int     = 255
)

func NewGreySubCommand() *cli.Command {
	return &cli.Command{
		Name:   "grey",
		Usage:  "apply the grey scale filter",
		Action: applyGreyFilter,
	}
}

func applyGreyFilter(c *cli.Context) error {

	return nil
}
