package main

import (
	"image"
	"image/color"

	"github.com/urfave/cli/v2"
)

const (
	IntensityFactor uint8 = 120
	HighestValue    uint8 = 255
	MeanValue       uint8 = 150
	LowestValue     uint8 = 0
)

func NewSketchSubCommand() *cli.Command {
	return &cli.Command{
		Name:   "sketch",
		Usage:  "apply the sketch filter",
		Action: applySketchFilter,
	}
}

func applySketchFilter(c *cli.Context) error {
	filePath := c.Args().First()

	e := new(Effect)

	file, err := e.ReadFile(filePath)
	if err != nil {
		return err
	}

	imgConf, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}

	width, height := imgConf.Width, imgConf.Height

	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	output := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			newR, newG, newB := e.GetSketchRGB(r/EightBits, g/EightBits, b/EightBits)
			sketch := color.RGBA{
				R: newR,
				G: newG,
				B: newB,
				A: uint8(Alpha),
			}
			output.Set(x, y, sketch)
		}
	}

	err = e.CreateFile(file, output, c.String("output"))
	if err != nil {
		return err
	}
	return nil
}
