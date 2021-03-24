package fltrs

import (
	"github.com/urfave/cli/v2"
)

const (
	SepiaRedForRed   float64 = 0.393
	SepiaGreenForRed float64 = 0.769
	SepiaBlueForRed  float64 = 0.189

	SepiaRedForGreen   float64 = 0.349
	SepiaGreenForGreen float64 = 0.686
	SepiaBlueForGreen  float64 = 0.168

	SepiaRedForBlue   float64 = 0.272
	SepiaGreenForBlue float64 = 0.534
	SepiaBlueForBlue  float64 = 0.131
)

func NewSepiaSubCommand() *cli.Command {
	effect := new(Effect)
	return &cli.Command{
		Name:  "sepia",
		Usage: "tranform an image into sepia tone",
	}
}
