package fltrs

import (
	"image"
	"image/color"

	"github.com/urfave/cli/v2"
)

//3x3
var blurKernel [][]float64

func init() {

	blurKernel = [][]float64{
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	}
}

func NewBlurSubCommand() *cli.Command {
	return &cli.Command{
		Name:   "blur",
		Usage:  "apply the blur filter",
		Action: applyBlurFilter,
	}
}

func applyBlurFilter(c *cli.Context) error {
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

	offset := len(blurKernel)
	output := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := offset; x < (width - offset); x++ {
		for y := offset; y < (height - offset); y++ {
			acc := make([]float64, 3)

			for i := 0; i < len(blurKernel); i++ {
				for j := 0; j < len(blurKernel); j++ {
					xn := x + i - offset
					xy := y + j - offset
					r, g, b, _ := img.At(xn, xy).RGBA()
					acc[0] += (float64(r) / float64(EightBits)) * blurKernel[i][j]
					acc[1] += (float64(g) / float64(EightBits)) * blurKernel[i][j]
					acc[2] += (float64(b) / float64(EightBits)) * blurKernel[i][j]
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
