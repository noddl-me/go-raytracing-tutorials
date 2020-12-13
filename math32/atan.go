package math32

import "math"

func Atan(x float32) float32 {
	return float32(math.Atan(float64(x)))
}

func Atan2(x, y float32) float32 {
	return float32(math.Atan2(float64(x), float64(y)))
}

func Atanh(x float32) float32 {
	return float32(math.Atanh(float64(x)))
}
