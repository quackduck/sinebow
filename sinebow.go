package sinebow

import (
	"math"
)

// HueToRGB converts a sinebow hue value in the range [0, 360) to an RGB color.
func HueToRGB(h float64) (r, g, b uint8) {
	pi := math.Pi
	h = math.Mod(h, 360) / 360.0
	r = uint8(math.Round(255.0 * (0.5 + 0.5*math.Sin(2*pi*h+pi/2))))
	g = uint8(math.Round(255.0 * (0.5 + 0.5*math.Sin(2*pi*h+pi/2+2*pi/3))))
	b = uint8(math.Round(255.0 * (0.5 + 0.5*math.Sin(2*pi*h+pi/2+4*pi/3))))
	return
}

// https://www.desmos.com/calculator/1blvi1yxtd

// RGBToHue converts an RGB color to a sinebow hue value in the range [0, 360).
func RGBToHue(r, g, b uint8) float64 {
	pi := math.Pi
	guess := math.Asin(2*float64(r)/255.0-1)/(2*pi) + 3.0/4.0
	if b > g {
		guess = 1 - guess
	}
	guess = 360 * guess

	justBelowMinStepWidth := 0.0125
	justAboveMaxStepWidth := 11.0
	i := 0.0
	for i = 0.0; i <= justAboveMaxStepWidth; i += justBelowMinStepWidth {
		_, guessG1, guessB1 := HueToRGB(guess + i)
		_, guessG2, guessB2 := HueToRGB(guess - i)
		if guessG1 == g && guessB1 == b {
			return guess + i
		} else if guessG2 == g && guessB2 == b {
			return guess - i
		}
		//fmt.Println(i)
	}
	return guess
}
