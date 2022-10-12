# go-color

```go
// interfaces
type Color interface {
	color.Color
	RGB() RGB
	HSV() HSV
	HSL() HSL
	YIQ() YIQ
	YUV() YUV
}

// apis
func NewRGBColor(r, g, b float64) Color { return RGB{r, g, b} }
func NewHSVColor(h, s, v float64) Color { return HSV{h, s, v} }
func NewHSLColor(h, s, l float64) Color { return HSL{h, s, l} }
func NewYIQColor(y, i, q float64) Color { return YIQ{y, i, q} }
func NewYUVColor(y, u, v float64) Color { return YUV{y, u, v} }


// use and todo
rgb := NewRGBColor(0, 255, 255)
rgb.HSV()
rgb.YUV()
...
```

- 一个支持各种颜色模式的转换库
- A conversion library that supports various color modes