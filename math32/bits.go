package math32

import "math"

// Mathematical constants.
const (
	E   = math.E
	Pi  = math.Pi
	Phi = math.Phi

	Sqrt2   = math.Sqrt2
	SqrtE   = math.SqrtE
	SqrtPi  = math.SqrtPi
	SqrtPhi = math.SqrtPhi

	Ln2    = math.Ln2
	Log2E  = 1 / Ln2
	Ln10   = math.Ln10
	Log10E = 1 / Ln10
)

// Floating-point limit values. Max is the largest finite value representable by the type. SmallestNonzero is the smallest positive, non-zero value representable by the type.
const (
	MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)
)

const (
	uvnan    = 0x7FC00001
	uvinf    = 0x7F800000
	uvneginf = 0xFF800000
	mask     = 0xFF
	shift    = 32 - 8 - 1
	bias     = 127
)

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = uvinf
	} else {
		v = uvneginf
	}
	return math.Float32frombits(v)
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float32 { return math.Float32frombits(uvnan) }

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN(f float32) (is bool) {
	return f != f
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf(f float32, sign int) bool {
	return sign >= 0 && f > math.MaxFloat32 || sign <= 0 && f < -math.MaxFloat32
}

// Signbit returns true if x is negative or negative zero.
func Signbit(x float32) bool {
	return math.Float32bits(x)&(1<<31) != 0
}

// normalize returns a normal number y and exponent exp
// satisfying x == y × 2**exp. It assumes x is finite and non-zero.
func normalize(x float32) (y float32, exp int) {
	const SmallestNormal = 1.1754943508222875079687365e-38 // 2**-(127 - 1)
	if Abs(x) < SmallestNormal {
		return x * (1 << shift), -shift
	}
	return x, 0
}
