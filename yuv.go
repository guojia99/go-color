package go_color

type YUV struct {
	Y, U, V float64
}

func (c YUV) RGBA() (r, g, b, a uint32) { return c.RGB().RGBA() }
func (c YUV) YUV() YUV                  { return c }
func (c YUV) HSV() HSV                  { return c.RGB().HSV() }
func (c YUV) HSL() HSL                  { return c.RGB().HSL() }
func (c YUV) YIQ() YIQ                  { return c.RGB().YIQ() }
func (c YUV) RGB() RGB {
	c.Y -= 16
	c.U -= 128
	c.V -= 128
	return RGB{
		R: checkRGB0To255(1.164*c.Y + 0.000*c.U + 1.596*c.V),
		G: checkRGB0To255(1.164*c.Y - 0.392*c.U - 0.813*c.V),
		B: checkRGB0To255(1.164*c.Y + 2.017*c.U + 0.000*c.V),
	}
}
