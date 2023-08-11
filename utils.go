/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package go_color

func maxAndMin(c Color) (max float64, min float64) {
	var v1, v2, v3 float64
	switch data := c.(type) {
	case RGB:
		v1, v2, v3 = data.R, data.G, data.B
	case HSL:
		v1, v2, v3 = data.H, data.L, data.S
	case HSV:
		v1, v2, v3 = data.H, data.S, data.V
	case YIQ:
		v1, v2, v3 = data.Y, data.I, data.Q
	case YUV:
		v1, v2, v3 = data.Y, data.U, data.V
	}
	max, min = v1, v2
	if min > max {
		max, min = min, max
	}
	if max < v3 {
		max = v3
	}
	if min > v3 {
		min = v3
	}
	return
}

func checkRGB0To255(v float64) float64 {
	if v < 0 {
		return 0.0
	}
	if v > 255.0 {
		return 255.0
	}
	return v
}
