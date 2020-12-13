package utils

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	nan32    = 0x7F800001
	inf32    = 0x7f000000
	neginf32 = 0xff000000
	PI       = math.Pi
)

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf32(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = inf32
	} else {
		v = neginf32
	}
	return math.Float32frombits(v)
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN32() float32 { return math.Float32frombits(nan32) }

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN32(f float32) (is bool) {
	// IEEE 754 says that only NaNs satisfy f != f.
	return f != f
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf32(f float32, sign int) bool {
	// Test for infinity by comparing against maximum float.
	// To avoid the floating-point hardware, could use:
	//	x := Float64bits(f);
	//	return sign >= 0 && x == uvinf || sign <= 0 && x == uvneginf;
	return sign >= 0 && f > math.MaxFloat32 || sign <= 0 && f < -math.MaxFloat32
}

func ToRadians(degrees float32) float32 {
	return degrees * PI / 180
}

func Randf() float32 {
	return rand.Float32()
}

func RandfMinMax(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func Clamp(x, min, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func Minf(x, y float32) float32 {
	if x <= y {
		return x
	}
	return y
}

func Maxf(x, y float32) float32 {
	if x >= y {
		return x
	}
	return y
}

func Sqrtf(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func Sinf(x float32) float32{
	return float32(math.Sin(float64(x)))
}

func Cosf(x float32) float32{
	return float32(math.Cos(float64(x)))
}

func Tanf(x float32) float32{
	return float32(math.Tan(float64(x)))
}

func PrintProgress(current, max int) {
	scale := 50.0
	fmt.Print("Rendering [")
	percentDone := int((float64(current) / float64(max)) * scale)
	for i := 0; i < percentDone; i++ {
		fmt.Print("=")
	}
	for i := percentDone; i < int(scale); i++ {
		fmt.Print(" ")
	}
	fmt.Printf("] %d out of %d \n", current+1, max)
}
