package cmd

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const (
	// to get RGB into 8 bits representation
	EightBits       = 257
	RedWaveLength   = 0.21
	GreenWaveLength = 0.72
	BlueWaveLength  = 0.07
	// For linux and Mac home directory
	Home = "$HOME"
)

var supportedExtensions map[string]bool

func init() {
	supportedExtensions = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
}

func NewTransformCommand() *cli.Command {
	return &cli.Command{
		Name:    "transform",
		Aliases: []string{"t"},
		Usage:   "transform takes a path as argument where the image is located.",
		Action:  transformAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "grey",
				Usage: "transform an image into grey-scale",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "path where the output result will be placed",
				Value: os.Getenv(Home),
			},
		},
	}
}

func transformAction(c *cli.Context) error {
	filePath := c.Args().First()
	if len(filePath) == 0 {
		return errors.New("transform command requires a path as argument")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	if !isValidExtension(file.Name()) {
		return errors.New("image file is not supported. Supported extensions (.png, .jpg, .jpeg)")
	}

	// Move the rest of the code another independent function

	imgConf, _, err := image.DecodeConfig(file)
	if err != nil {
		return errors.New(err.Error())
	}
	width, height := imgConf.Width, imgConf.Height

	// reset io.Reader
	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		return errors.New(err.Error())
	}
	// new image
	result := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// TODO: implement a switch-case for more
			rgb := getGreyScaleRGB(r/EightBits, g/EightBits, b/EightBits)
			filter := color.RGBA{
				R: rgb,
				G: rgb,
				B: rgb,
				A: 255,
			}
			result.Set(x, y, filter)
		}
	}

	f, err := os.Create("result.png")
	if err != nil {
		return errors.New("error after creating the desire result")
	}

	err = png.Encode(f, result)
	if err != nil {
		return errors.New("could not transform the current image")
	}
	fmt.Printf("file created: %s\n", f.Name())
	return nil
}

func isValidExtension(name string) bool {
	extension := filepath.Ext(name)
	_, ok := supportedExtensions[extension]
	return ok
}

func getGreyScaleRGB(r, g, b uint32) uint8 {
	// using the luminosity algorithm 0.21 * R + 0.72 * G + 0.07 * B
	return uint8((RedWaveLength * float64(r)) + (GreenWaveLength * float64(g)) + (BlueWaveLength * float64(b)))
}
