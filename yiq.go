/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package go_color

type YIQ struct {
	Y, I, Q float64
}

func (c YIQ) RGBA() (r, g, b, a uint32) { return c.RGB().RGBA() }
func (c YIQ) YUV() YUV                  { return YUV{} }

func (c YIQ) RGB() RGB {
	return RGB{
		R: c.Y + 0.9563*c.I + 0.6210*c.Q,
		G: c.Y - 0.2721*c.I - 0.6474*c.Q,
		B: c.Y - 1.1070*c.I + 1.7046*c.Q,
	}
}

func (c YIQ) HSV() HSV {
	//TODO implement me
	panic("implement me")
}

func (c YIQ) HSL() HSL {
	//TODO implement me
	panic("implement me")
}

func (c YIQ) YIQ() YIQ {
	//TODO implement me
	panic("implement me")
}
