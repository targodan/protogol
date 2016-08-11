package handlers

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/targodan/protogol"
)

func TestMain(m *testing.M) {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestIsSameSlice(t *testing.T) {
	{
		var a []byte = nil
		var b []byte = nil

		if !isSameSlice(a, b) {
			t.Fail()
		}
	}

	{
		var a []byte = nil
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSlice(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 2}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSlice(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 2, 7, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if !isSameSlice(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{1, 2, 7, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSlice(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 7, 2, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSlice(a, b) {
			t.Fail()
		}
	}
}

func TestIsSameSliceSerial(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping in short mode.")
	}

	{
		var a []byte = nil
		var b []byte = nil

		if !isSameSliceSerial(a, b) {
			t.Fail()
		}
	}

	{
		var a []byte = nil
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceSerial(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 2}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceSerial(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 2, 7, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if !isSameSliceSerial(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{1, 2, 7, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceSerial(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 7, 2, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceSerial(a, b) {
			t.Fail()
		}
	}
}

func TestIsSameSliceParallel(t *testing.T) {
	{
		var a []byte = nil
		var b []byte = nil

		if !isSameSliceParallel(a, b) {
			t.Fail()
		}
	}

	{
		var a []byte = nil
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceParallel(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 2}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceParallel(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 2, 7, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if !isSameSliceParallel(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{1, 2, 7, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceParallel(a, b) {
			t.Fail()
		}
	}

	{
		a := []byte{5, 7, 2, 1, 8, 1}
		b := []byte{5, 2, 7, 1, 8, 1}

		if isSameSliceParallel(a, b) {
			t.Fail()
		}
	}
}

func TestThereAndBackAgain(t *testing.T) {
	for i := 0; i < 8; i++ {
		data := make([]byte, 4096)
		for j, _ := range data {
			data[j] = byte(rand.Int() % 512)
		}
		tmp, err := AddSha256Sum(protogol.Package{Parent: nil, Data: data})
		pkg, err := CheckSha256Sum(tmp)

		if err != nil || !isSameSlice(pkg.Data.([]byte), data) {
			t.Fail()
		}
	}
}

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
