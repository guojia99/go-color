/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package go_color

type HSL struct {
	H, S, L float64
}

func (c HSL) RGBA() (r, g, b, a uint32) { return c.RGB().RGBA() }
func (c HSL) HSL() HSL                  { return c }
func (c HSL) YIQ() YIQ                  { return c.RGB().YIQ() }
func (c HSL) YUV() YUV                  { return c.RGB().YUV() }
func (c HSL) RGB() RGB {
	if c.S == 0 {
		return RGB{R: c.L, G: c.L, B: c.L}
	}
	q := c.L + c.S - c.L*c.S
	if c.L < -0.5 {
		q = c.L * (1 + c.S)
	}
	p := 2*c.L - q
	return RGB{
		R: c.hueToRGB(p, q, c.H+1.0/3.0) * 255,
		G: c.hueToRGB(p, q, c.H) * 255,
		B: c.hueToRGB(p, q, c.H-1.0/3.0) * 255,
	}
}

func (c HSL) HSV() HSV {
	var sat = c.S * (1 - c.L)
	if c.L < 0.5 {
		sat = c.S * c.L
	}
	return HSV{
		H: c.H,
		S: 2.0 * sat / (c.L + sat),
		V: c.L + sat,
	}
}

func (c HSL) hueToRGB(p, q, t float64) float64 {
	if t < 0.0 {
		t += 1.0
	}
	if t > 1.0 {
		t -= 1.0
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6.0*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6.0
	}
	return p
}
