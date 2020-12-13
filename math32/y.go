package math32

import "math"

func Y0(x float32) float32 {
	return float32(math.Y0(float64(x)))
}

func Y1(x float32) float32 {
	return float32(math.Y1(float64(x)))
}
