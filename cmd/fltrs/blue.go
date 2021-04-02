package fltrs

import (
	"image"
	"image/color"

	"github.com/urfave/cli/v2"
)

const (
	MaxBlue         = 255
	MinRedForBlue   = 180
	MinGreenForBlue = 180
)

func NewBlueSubCommand() *cli.Command {
	return &cli.Command{
		Name:   "blue",
		Usage:  "apply blue filter",
		Action: applyBlueFilter,
	}
}

func applyBlueFilter(c *cli.Context) error {
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
			newR, newG, newB := e.GetBlueRGB(r/EightBits, g/EightBits, b/EightBits)
			filter := color.RGBA{
				R: newR,
				G: newG,
				B: newB,
				A: uint8(Alpha),
			}
			output.Set(x, y, filter)
		}
	}
	err = e.CreateFile(file, output, c.String("output"))
	if err != nil {
		return err
	}

	return nil
}
