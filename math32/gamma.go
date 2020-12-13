package math32

import "math"

func Gamma(x float32) float32 {
	return float32(math.Gamma(float64(x)))
}

func Lgamma(x float32) (lgamma float32, sign int) {
	lg, sign := math.Lgamma(float64(x))
	return float32(lg), sign
}
