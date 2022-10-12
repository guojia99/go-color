package go_color

import (
	"fmt"
	"testing"
)

func TestHSV_RGB(t *testing.T) {
	hsv := HSV{100, 60, 40}
	fmt.Println(hsv.RGB())
}

func TestRGB_HSV(t *testing.T) {
	rgb := RGB{45, 215, 0}
	fmt.Println(rgb.HSV())
}

func TestRGB_HLS(t *testing.T) {
	rgb := RGB{82, 0, 87}
	fmt.Printf("%+v\n", rgb.HSL())
	fmt.Printf("want %+v\n", HSL{H: 296, S: 1, L: 0.17058824})
}

func Test_maxAndMin(t *testing.T) {
	type args struct {
		c Color
	}
	tests := []struct {
		name    string
		args    args
		wantMax float64
		wantMin float64
	}{
		{
			name: "ok",
			args: args{
				c: RGB{0, 1, 2},
			},
			wantMax: 2,
			wantMin: 0,
		},
		{
			name: "ok2",
			args: args{
				c: RGB{0, 2, 1},
			},
			wantMax: 2,
			wantMin: 0,
		},
		{
			name: "ok3",
			args: args{
				c: RGB{1, 2, 0},
			},
			wantMax: 2,
			wantMin: 0,
		},
		{
			name: "ok4",
			args: args{
				c: RGB{1, 0, 2},
			},
			wantMax: 2,
			wantMin: 0,
		},
		{
			name: "ok5",
			args: args{
				c: RGB{2, 0, 1},
			},
			wantMax: 2,
			wantMin: 0,
		},
		{
			name: "ok6",
			args: args{
				c: RGB{2, 1, 0},
			},
			wantMax: 2,
			wantMin: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotMin := maxAndMin(tt.args.c)
			if gotMax != tt.wantMax {
				t.Errorf("maxAndMin() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotMin != tt.wantMin {
				t.Errorf("maxAndMin() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}
