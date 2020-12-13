package cmplx32

import "git.maze.io/go/math32"

// Abs returns the absolute value (also called the modulus) of x.
func Abs(x complex64) float32 { return math32.Hypot(real(x), imag(x)) }
