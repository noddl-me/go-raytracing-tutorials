package utils

import (
	"raytracing-tutorials/math32"
	"testing"
)

func BenchmarkSqrtf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = Sqrtf(k)
	}
}

func BenchmarkSqrt32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = math32.Sqrt(k)
	}
}

func BenchmarkSinf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = Sinf(k)
	}
}

func BenchmarkSinGo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = math32.SinGo(k)
	}
}

func BenchmarkSin32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = math32.Sin(k)
	}
}

func BenchmarkCosf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = Cosf(k)
	}
}
func BenchmarkCos32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = math32.Cos(k)
	}
}

func BenchmarkCosGo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = math32.CosGo(k)
	}
}

func BenchmarkTanf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = Tanf(k)
	}
}
func BenchmarkTan32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = math32.Tan(k)
	}
}

func BenchmarkTanGo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := float32(i) * 5.0
		k = math32.TanGo(k)
	}
}