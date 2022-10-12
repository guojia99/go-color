package go_color

import (
	"image/color"
)

// Color each color mode can be calculated by algorithm to get a new color mode
// The algorithm of this method is implemented according to the documentation
//  1. with HSV and HSL
//     see https://en.wikipedia.org/wiki/HSL_and_HSV
//  2. with YIQ
//     see https://en.wikipedia.org/wiki/YIQ
//  3. with YUV
//     see https://en.wikipedia.org/wiki/YUV
type Color interface {
	color.Color
	RGB() RGB
	HSV() HSV
	HSL() HSL
	YIQ() YIQ
	YUV() YUV
}

func NewRGBColor(r, g, b float64) Color { return RGB{r, g, b} }
func NewHSVColor(h, s, v float64) Color { return HSV{h, s, v} }
func NewHSLColor(h, s, l float64) Color { return HSL{h, s, l} }
func NewYIQColor(y, i, q float64) Color { return YIQ{y, i, q} }
func NewYUVColor(y, u, v float64) Color { return YUV{y, u, v} }
