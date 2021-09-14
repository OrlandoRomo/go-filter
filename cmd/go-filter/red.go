package main

import (
	"image"
	"image/color"

	"github.com/urfave/cli/v2"
)

const (
	MaxRed         = 255
	MinGreenForRed = 200
	MinBlueForRed  = 200
)

func NewRedSubCommand() *cli.Command {
	return &cli.Command{
		Name:   "red",
		Usage:  "apply red filter",
		Action: applyRedFilter,
	}
}

func applyRedFilter(c *cli.Context) error {
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
			go func(x, y int) {
				r, g, b, _ := img.At(x, y).RGBA()
				newR, newG, newB := e.GetRedRGB(r/EightBits, g/EightBits, b/EightBits)
				filter := color.RGBA{
					R: newR,
					G: newG,
					B: newB,
					A: uint8(Alpha),
				}
				output.Set(x, y, filter)
			}(x, y)
		}
	}
	err = e.CreateFile(file, output, c.String("output"))
	if err != nil {
		return err
	}

	return nil
}
