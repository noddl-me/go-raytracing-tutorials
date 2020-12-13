package cmplx32

import "git.maze.io/go/math32"

func Phase(x complex64) float32 { return math32.Atan2(imag(x), real(x)) }
