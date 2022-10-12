package go_color

import "math"

type HSV struct {
	H, S, V float64
}

func (c HSV) RGBA() (r, g, b, a uint32) { return c.RGB().RGBA() }
func (c HSV) HSV() HSV                  { return c }
func (c HSV) YIQ() YIQ                  { return c.RGB().YIQ() }
func (c HSV) YUV() YUV                  { return c.RGB().YUV() }

func (c HSV) RGB() RGB {
	if c.S == 0 {
		val := float64(uint32(c.V * 255))
		return RGB{R: val, G: val, B: val}
	}

	C := (c.V / 100.0) * (c.S / 100.0)
	X := C * (1 - math.Abs(math.Mod(c.H/60, 2)-1))
	mod := (c.V / 100) - C

	var r, g, b float64
	switch {
	case c.H < 60:
		r, g, b = C, X, 0
	case c.H < 120:
		r, g, b = X, C, 0
	case c.H < 180:
		r, g, b = 0, C, X
	case c.H < 240:
		r, g, b = 0, X, C
	case c.H < 300:
		r, g, b = X, 0, C
	case c.H < 360:
		r, g, b = C, 0, X
	}
	return RGB{
		R: checkRGB0To255((r + mod) * 255),
		G: checkRGB0To255((g + mod) * 255),
		B: checkRGB0To255((b + mod) * 255),
	}
}

func (c HSV) HSL() HSL {
	var hsl = HSL{
		H: c.H,
		L: c.H / 2.0,
	}
	if (2.0-c.S)*c.V < 1.0 {
		hsl.S = c.S * c.V / ((2.0 - c.S) * c.V)
		return hsl
	}
	hsl.S = c.S * c.V / (2.0 - (2.0-c.S)*c.V)
	return hsl
}
