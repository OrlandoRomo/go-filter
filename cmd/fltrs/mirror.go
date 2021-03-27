package fltrs

import (
	"image"

	"github.com/urfave/cli/v2"
)

func NewMirrorSubCommand() *cli.Command {
	return &cli.Command{
		Name:   "mirror",
		Usage:  "apply the flip or mirror filter",
		Action: applyMirrorFilter,
	}
}

func applyMirrorFilter(c *cli.Context) error {
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
			xp := imgConf.Width - x - 1
			color := img.At(x, y)
			// reversing matrix horizontally
			output.Set(xp, y, color)
		}
	}

	err = e.CreateFile(file, output, c.String("output"))
	if err != nil {
		return err
	}
	return nil
}
