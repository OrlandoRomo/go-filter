package cmd

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/OrlandoRomo/imgfltr/cmd/fltrs"
	"github.com/urfave/cli/v2"
)

const (
	// to get RGB into 8 bits representation
	EightBits uint32 = 257
	Alpha     int    = 255
	// For linux and Mac home directory
	Home = "$HOME"
)

var supportedExtensions map[string]bool

func init() {
	supportedExtensions = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".webp": true,
	}
}

func NewSetCommand() *cli.Command {
	return &cli.Command{
		Name:  "set",
		Usage: "set requires a filter name",
		Subcommands: []*cli.Command{
			fltrs.NewGreySubCommand(),
			fltrs.NewSepiaSubCommand(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "output",
				Usage: "path where the output result will be placed",
				Value: os.Getenv(Home),
			},
		},
	}
}

func SetAction(c *cli.Context) error {
	filePath := c.Args().First()
	if len(filePath) == 0 {
		return errors.New("transform command requires a path as argument")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	if !isValidExtension(file.Name()) {
		return errors.New("image file is not supported. Supported extensions (.png, .jpg, .jpeg, .webp)")
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
			// rgb := getGreyScaleRGB(r/EightBits, g/EightBits, b/EightBits)
			// filter := color.RGBA{
			// 	R: rgb,
			// 	G: rgb,
			// 	B: rgb,
			// 	A: Alpha,
			// }
			// result.Set(x, y, filter)
			tr, tg, tb := fltrs.GetSepiaTone(r/EightBits, g/EightBits, b/EightBits)
			filter := color.RGBA{
				R: tr,
				G: tg,
				B: tb,
				A: uint8(Alpha),
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
