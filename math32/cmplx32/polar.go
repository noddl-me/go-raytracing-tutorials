package cmplx32

// Polar returns the absolute value r and phase θ of x,
// such that x = r * e**θi.
// The phase is in the range [-Pi, Pi].
func Polar(x complex64) (r, θ float32) {
	return Abs(x), Phase(x)
}
