package fltrs

type Filter interface {
	GetGreyRGB(r, g, b uint32) uint8
	GetSepiaRGB(r, g, b uint32) (uint8, uint8, uint8)
}

type Effect struct {
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
