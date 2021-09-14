package main

import (
	"fmt"
	"image"
	"image/color"

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
	return &cli.Command{
		Name:   "sepia",
		Usage:  "apply the sepia filter",
		Action: applySepiaFilter,
	}
}

func applySepiaFilter(c *cli.Context) error {
	filePath := c.Args().First()
	e := new(Effect)

	file, err := e.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	imgConf, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}
	width, height := imgConf.Width, imgConf.Height

	// reset io.Reader
	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	//new image to create
	output := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			go func(x, y int) {
				r, g, b, _ := img.At(x, y).RGBA()
				tr, tg, tb := e.GetSepiaRGB(r/EightBits, g/EightBits, b/EightBits)
				filter := color.RGBA{
					R: tr,
					G: tg,
					B: tb,
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
