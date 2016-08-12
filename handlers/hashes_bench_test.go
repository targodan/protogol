package handlers

import (
	"testing"
)

var result bool

func benchmarkIsSliceEqualsSerial(s1 []byte, s2 []byte, b *testing.B) {
	var dummy bool
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dummy = isSliceEqualSerial(s1, s2)
	}
	result = dummy
}

func benchmarkIsSliceEqualsParallel(s1 []byte, s2 []byte, b *testing.B) {
	var dummy bool
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dummy = isSliceEqualParallel(s1, s2)
	}
	result = dummy
}

func BenchmarkIsSliceEqualsSerial128(b *testing.B) {
	d := make([]byte, 128)
	benchmarkIsSliceEqualsSerial(d, d, b)
}

func BenchmarkIsSliceEqualsParallel128(b *testing.B) {
	d := make([]byte, 128)
	benchmarkIsSliceEqualsParallel(d, d, b)
}

func BenchmarkIsSliceEqualsSerial256(b *testing.B) {
	d := make([]byte, 256)
	benchmarkIsSliceEqualsSerial(d, d, b)
}

func BenchmarkIsSliceEqualsParallel256(b *testing.B) {
	d := make([]byte, 256)
	benchmarkIsSliceEqualsParallel(d, d, b)
}

func BenchmarkIsSliceEqualsSerial512(b *testing.B) {
	d := make([]byte, 512)
	benchmarkIsSliceEqualsSerial(d, d, b)
}

func BenchmarkIsSliceEqualsParallel512(b *testing.B) {
	d := make([]byte, 512)
	benchmarkIsSliceEqualsParallel(d, d, b)
}

func BenchmarkIsSliceEqualsSerial1024(b *testing.B) {
	d := make([]byte, 1024)
	benchmarkIsSliceEqualsSerial(d, d, b)
}

func BenchmarkIsSliceEqualsParallel1024(b *testing.B) {
	d := make([]byte, 1024)
	benchmarkIsSliceEqualsParallel(d, d, b)
}

func BenchmarkIsSliceEqualsSerial2048(b *testing.B) {
	d := make([]byte, 2048)
	benchmarkIsSliceEqualsSerial(d, d, b)
}

func BenchmarkIsSliceEqualsParallel2048(b *testing.B) {
	d := make([]byte, 2048)
	benchmarkIsSliceEqualsParallel(d, d, b)
}

func BenchmarkIsSliceEqualsSerial4096(b *testing.B) {
	d := make([]byte, 4096)
	benchmarkIsSliceEqualsSerial(d, d, b)
}

func BenchmarkIsSliceEqualsParallel4096(b *testing.B) {
	d := make([]byte, 4096)
	benchmarkIsSliceEqualsParallel(d, d, b)
}
