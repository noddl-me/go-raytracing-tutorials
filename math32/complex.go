package math32

func ComplexExp(x complex64) complex64 {
	r := Exp(real(x))
	s, c := Sincos(imag(x))
	return complex(r*c, s*c)
}
