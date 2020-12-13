package math32

import "math"

func Erf(x float32) float32 {
	return float32(math.Erf(float64(x)))
}

func Erfc(x float32) float32 {
	return float32(math.Erfc(float64(x)))
}
