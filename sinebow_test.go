package sinebow

import (
	"testing"
)

func TestRGBToHue(test *testing.T) {
	for t := 0.0; t <= 361; t += 0.0001 {
		r, g, b := HueToRGB(t)
		r2, g2, b2 := HueToRGB(RGBToHue(r, g, b))
		if r != r2 || g != g2 || b != b2 {
			test.Errorf("RGBToHue() = %v, want %v for hue %v", t, RGBToHue(r, g, b), t)
		}
	}
}
