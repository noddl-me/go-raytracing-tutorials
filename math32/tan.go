// This file contains a subset of functions of the std
// math library from Go, but converted from float64 to float32.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import "math"

/*
	Floating-point tangent.
*/

// The original C code, the long comment, and the constants
// below were from http://netlib.sandia.gov/cephes/cmath/sin.c,
// available from http://www.netlib.org/cephes/cmath.tgz.
// The go code is a simplified version of the original C.
//
// For more information see:
// https://github.com/golang/go/blob/master/src/math/tan.go
//
// Cephes Math Library Release 2.8:  June, 2000
// Copyright 1984, 1987, 1989, 1992, 2000 by Stephen L. Moshier
//
// The readme file at http://netlib.sandia.gov/cephes/ says:
//    Some software in this archive may be from the book _Methods and
// Programs for Mathematical Functions_ (Prentice-Hall or Simon & Schuster
// International, 1989) or from the Cephes Mathematical Library, a
// commercial product. In either event, it is copyrighted by the author.
// What you see here may be used freely but it comes with no support or
// guarantee.
//
//   The two known misprints in the book are repaired here in the
// source listings for the gamma function and the incomplete beta
// integral.
//
//   Stephen L. Moshier
//   moshier@na-net.ornl.gov

// tan coefficients
var _tanP = [...]float32{
	-1.30936939181383777646E4, // 0xc0c992d8d24f3f38
	1.15351664838587416140E6,  // 0x413199eca5fc9ddd
	-1.79565251976484877988E7, // 0xc1711fead3299176
}
var _tanQ = [...]float32{
	1.00000000000000000000E0,
	1.36812963470692954678E4,  //0x40cab8a5eeb36572
	-1.32089234440210967447E6, //0xc13427bc582abc96
	2.50083801823357915839E7,  //0x4177d98fc2ead8ef
	-5.38695755929454629881E7, //0xc189afe03cbe5a31
}

// Tan returns the tangent of the radian argument x.
//
// Special cases are:
//	Tan(±0) = ±0
//	Tan(±Inf) = NaN
//	Tan(NaN) = NaN
func Tan(x float32) float32

func tan(x float32) float32 {
	const (
		PI4A = 7.85398125648498535156E-1                             // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668E-8                             // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645E-15                            // 0x3ce8469898cc5170,
		M4PI = 1.273239544735162542821171882678754627704620361328125 // 4/pi
	)
	// special cases
	switch {
	case x == 0 || IsNaN(x):
		return x // return ±0 || NaN()
	case IsInf(x, 0):
		return NaN()
	}

	// make argument positive but save the sign
	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	j := int64(x * M4PI) // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := float32(j)      // integer part of x/(Pi/4), as float

	/* map zeros and singularities to origin */
	if j&1 == 1 {
		j++
		y++
	}

	z := ((x - y*PI4A) - y*PI4B) - y*PI4C
	zz := z * z

	if zz > 1e-14 {
		y = z + z*(zz*(((_tanP[0]*zz)+_tanP[1])*zz+_tanP[2])/((((zz+_tanQ[1])*zz+_tanQ[2])*zz+_tanQ[3])*zz+_tanQ[4]))
	} else {
		y = z
	}
	if j&2 == 2 {
		y = -1 / y
	}
	if sign {
		y = -y
	}
	return y
}

// TanGo is the pure Go Tan function.
var TanGo = tan

func Tanh(x float32) float32 {
	return float32(math.Tanh(float64(x)))
}
