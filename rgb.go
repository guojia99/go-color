/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package go_color

import "math"

type RGB struct {
	R, G, B float64
}

func (c RGB) RGBA() (r, g, b, a uint32) { return uint32(c.R), uint32(c.G), uint32(c.B), uint32(1) }

func (c RGB) RGB() RGB { return c }

func (c RGB) YIQ() YIQ {
	rgb := c.except255()
	yiq := YIQ{
		Y: (0.299000*rgb.R + 0.587000*rgb.G + 0.114000*rgb.B) * 100,
		I: (0.595716*rgb.R - 0.274453*rgb.G - 0.321264*rgb.B) * 100,
		Q: (0.211456*rgb.R - 0.522591*rgb.G + 0.311135*rgb.B) * 100,
	}
	if yiq.I < 0 {
		yiq.I += 0.5957 * 100
	}
	if yiq.Q < 0 {
		yiq.Q += 0.5226 * 100
	}
	return yiq
}

func (c RGB) HSL() HSL {
	var hsl HSL
	rgb := c.except255()
	max, min := maxAndMin(rgb)
	delta := max - min
	if hsl.L = (max + min) / 2.0; delta == 0 {
		return hsl
	}
	if hsl.S = delta / (2.0 - max - min); hsl.L <= 0.5 {
		hsl.S = delta / (max + min)
	}
	hue := 0.0
	if rgb.R == max {
		hue = ((rgb.G - rgb.B) / 6.0) / delta
	} else if rgb.G == max {
		hue = (1.0 / 3.0) + ((rgb.B-rgb.R)/6.0)/delta
	} else {
		hue = (2.0 / 3.0) + ((rgb.R-rgb.G)/6.0)/delta
	}
	if hue < 0 {
		hue += 1
	}
	if hue > 1 {
		hue -= 1
	}
	hsl.H = float64(uint32(hue * 360))
	return hsl
}

func (c RGB) HSV() HSV {
	var hsv HSV
	rgb := c.except255()
	max, min := maxAndMin(rgb)
	delta := max - min

	if delta == 0 {
		hsv.H = 0
	} else if rgb.R == max {
		hsv.H = math.Mod(60.0*((rgb.G-rgb.B)/delta)+360.0, 360.0)
	} else if rgb.G == max {
		hsv.H = math.Mod(60.0*((rgb.B-rgb.R)/delta)+120.0, 360.0)
	} else if rgb.B == max {
		hsv.H = math.Mod(60.0*((rgb.R-rgb.G)/delta)+240.0, 360.0)
	}

	if max != 0 {
		hsv.S = (delta / max) * 100
	}
	hsv.V = max * 100
	return hsv
}

func (c RGB) YUV() YUV {
	return YUV{
		Y: 0.257*c.R + 0.504*c.G + 0.098*c.B + 16.0,
		U: -0.148*c.R - 0.291*c.G + 0.439*c.B + 128.0,
		V: 0.439*c.R - 0.368*c.G - 0.071*c.B + 128.0,
	}
}

func (c RGB) except255() RGB {
	return RGB{
		R: c.R / 255.0,
		G: c.G / 255.0,
		B: c.B / 255.0,
	}
}

func (c RGB) except255rgbC() RGB {
	rgb := c.except255()
	max, min := maxAndMin(c)
	return RGB{
		R: (max - rgb.R) / (max - min),
		G: (max - rgb.G) / (max - min),
		B: (max - rgb.B) / (max - min),
	}
}
