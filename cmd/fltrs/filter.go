package fltrs

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	// to get RGB into 8 bits representation
	EightBits     uint32 = 257
	Alpha         int    = 255
	maxCharacters int    = 7
	letterRunes   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var supportedExtensions map[string]bool

func init() {
	rand.Seed(time.Now().UnixNano())
	supportedExtensions = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".webp": true,
	}
}

type Filter interface {
	GetGreyRGB(r, g, b uint32) uint8
	GetSepiaRGB(r, g, b uint32) (uint8, uint8, uint8)
}

func isValidExtension(name string) bool {
	extension := filepath.Ext(name)
	_, ok := supportedExtensions[extension]
	return ok
}
func randomName() string {
	b := make([]byte, maxCharacters)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Effect struct{}

func (e *Effect) ReadFile(filePath string) (*os.File, error) {
	if len(filePath) == 0 {
		return nil, errors.New("`set` command requires a path as argument")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	if !isValidExtension(file.Name()) {
		return nil, errors.New("file is not supported. Supported extensions (.png, .jpg, .jpeg, .webp)")
	}
	return file, nil
}

func (e *Effect) CreateFile(f *os.File, img image.Image, outputPath string) error {
	ext := filepath.Ext(f.Name())
	file, err := os.Create(fmt.Sprintf("%s/%s%s", outputPath, randomName(), ext))
	if err != nil {
		return err
	}
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

//GetGreyRGB returns RGB values that represent the grey scale tone.
func (e *Effect) GetGreyRGB(r, g, b uint32) uint8 {
	// using the luminosity algorithm 0.21 * R + 0.72 * G + 0.07 * B
	return uint8((RedWaveLength * float64(r)) + (GreenWaveLength * float64(g)) + (BlueWaveLength * float64(b)))
}

// GetSepiaRGB gets the RGB values that represent the sepia scale tone
func (e *Effect) GetSepiaRGB(r, g, b uint32) (uint8, uint8, uint8) {
	// tr = 0.393R + 0.769G + 0.189B
	// tg = 0.349R + 0.686G + 0.168B
	// tb = 0.272R + 0.534G + 0.131B
	tr := (SepiaRedForRed * float64(r)) + (SepiaGreenForRed * float64(g)) + (SepiaBlueForRed * float64(b))
	tg := (SepiaRedForGreen * float64(r)) + (SepiaGreenForGreen * float64(g)) + (SepiaBlueForGreen * float64(b))
	tb := (SepiaRedForBlue * float64(r)) + (SepiaGreenForBlue * float64(g)) + (SepiaBlueForBlue * float64(b))

	if int(tr) > RGBA {
		tr = float64(RGBA)
	}

	if int(tg) > RGBA {
		tg = float64(RGBA)
	}

	if int(tb) > RGBA {
		tb = float64(RGBA)
	}

	return uint8(tr), uint8(tg), uint8(tb)
}
