package math32

import "math"

func J0(x float32) float32 {
	return float32(math.J0(float64(x)))
}

func J1(x float32) float32 {
	return float32(math.J1(float64(x)))
}

func Jn(n int, x float32) float32 {
	return float32(math.Jn(n, float64(x)))
}

func Yn(n int, x float32) float32 {
	return float32(math.Yn(n, float64(x)))
}
