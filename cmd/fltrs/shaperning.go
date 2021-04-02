package fltrs

import (
	"image"
	"image/color"

	"github.com/urfave/cli/v2"
)

var sharpKernel [][]float64

func init() {
	sharpKernel = [][]float64{
		{0.0, -0.5, 0.0},
		{-0.5, 3.0, -0.5},
		{0.0, -0.5, 0.0},
	}
}

func NewSharpSubCommand() *cli.Command {
	return &cli.Command{
		Name:   "sharp",
		Usage:  "apply sharp filter",
		Action: applySharpFilter,
	}
}

func applySharpFilter(c *cli.Context) error {
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

	offset := len(sharpKernel)
	output := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := offset; x < (width - offset); x++ {
		for y := offset; y < (height - offset); y++ {
			acc := make([]float64, 3)

			for i := 0; i < len(sharpKernel); i++ {
				for j := 0; j < len(sharpKernel); j++ {
					xn := x + i - offset
					yn := y + j - offset
					r, g, b, _ := img.At(xn, yn).RGBA()
					acc[0] += (float64(r) / float64(EightBits)) * sharpKernel[i][j]
					acc[1] += (float64(b) / float64(EightBits)) * sharpKernel[i][j]
					acc[2] += (float64(g) / float64(EightBits)) * sharpKernel[i][j]
				}
			}
			output.Set(x, y, color.RGBA{
				R: uint8(acc[0]),
				G: uint8(acc[1]),
				B: uint8(acc[2]),
				A: uint8(Alpha),
			})
		}
	}

	err = e.CreateFile(file, output, c.String("output"))
	if err != nil {
		return err
	}
	return nil
}
