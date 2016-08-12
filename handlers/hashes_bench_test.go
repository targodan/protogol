package handlers

import (
	"testing"
)

var result bool

func benchmarkIsSameSlicesSerial(s1 []byte, s2 []byte, b *testing.B) {
	var dummy bool
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dummy = isSameSliceSerial(s1, s2)
	}
	result = dummy
}

func benchmarkIsSameSlicesParallel(s1 []byte, s2 []byte, b *testing.B) {
	var dummy bool
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dummy = isSameSliceParallel(s1, s2)
	}
	result = dummy
}

func BenchmarkIsSameSlicesSerial128(b *testing.B) {
	d := make([]byte, 128)
	benchmarkIsSameSlicesSerial(d, d, b)
}

func BenchmarkIsSameSlicesParallel128(b *testing.B) {
	d := make([]byte, 128)
	benchmarkIsSameSlicesParallel(d, d, b)
}

func BenchmarkIsSameSlicesSerial256(b *testing.B) {
	d := make([]byte, 256)
	benchmarkIsSameSlicesSerial(d, d, b)
}

func BenchmarkIsSameSlicesParallel256(b *testing.B) {
	d := make([]byte, 256)
	benchmarkIsSameSlicesParallel(d, d, b)
}

func BenchmarkIsSameSlicesSerial512(b *testing.B) {
	d := make([]byte, 512)
	benchmarkIsSameSlicesSerial(d, d, b)
}

func BenchmarkIsSameSlicesParallel512(b *testing.B) {
	d := make([]byte, 512)
	benchmarkIsSameSlicesParallel(d, d, b)
}

func BenchmarkIsSameSlicesSerial1024(b *testing.B) {
	d := make([]byte, 1024)
	benchmarkIsSameSlicesSerial(d, d, b)
}

func BenchmarkIsSameSlicesParallel1024(b *testing.B) {
	d := make([]byte, 1024)
	benchmarkIsSameSlicesParallel(d, d, b)
}

func BenchmarkIsSameSlicesSerial2048(b *testing.B) {
	d := make([]byte, 2048)
	benchmarkIsSameSlicesSerial(d, d, b)
}

func BenchmarkIsSameSlicesParallel2048(b *testing.B) {
	d := make([]byte, 2048)
	benchmarkIsSameSlicesParallel(d, d, b)
}

func BenchmarkIsSameSlicesSerial4096(b *testing.B) {
	d := make([]byte, 4096)
	benchmarkIsSameSlicesSerial(d, d, b)
}

func BenchmarkIsSameSlicesParallel4096(b *testing.B) {
	d := make([]byte, 4096)
	benchmarkIsSameSlicesParallel(d, d, b)
}
