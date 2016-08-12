package handlers

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/targodan/protogol"

	. "github.com/smartystreets/goconvey/convey"
)

const testRandomDataNTimes = 16

func TestMain(m *testing.M) {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

type futSameSlice func(a []byte, b []byte) bool

func testSameSliceFut(fut futSameSlice, t *testing.T) {
	Convey("These slices should be equal.", t, func() {
		{
			var a []byte
			var b []byte

			So(fut(a, b), ShouldBeTrue)
		}

		{
			a := []byte{5, 2, 7, 1, 8, 1}
			b := []byte{5, 2, 7, 1, 8, 1}

			So(fut(a, b), ShouldBeTrue)
		}
	})

	Convey("These slices should not be equal.", t, func() {
		{
			var a []byte
			b := []byte{5, 2, 7, 1, 8, 1}

			So(fut(a, b), ShouldBeFalse)
		}

		{
			a := []byte{5, 2}
			b := []byte{5, 2, 7, 1, 8, 1}

			So(fut(a, b), ShouldBeFalse)
		}

		{
			a := []byte{1, 2, 7, 1, 8, 1}
			b := []byte{5, 2, 7, 1, 8, 1}

			So(fut(a, b), ShouldBeFalse)
		}

		{
			a := []byte{5, 7, 2, 1, 8, 1}
			b := []byte{5, 2, 7, 1, 8, 1}

			So(fut(a, b), ShouldBeFalse)
		}
	})
}

func TestIsSameSlice(t *testing.T) {
	testSameSliceFut(isSameSlice, t)
}

func TestIsSameSliceSerial(t *testing.T) {
	testSameSliceFut(isSameSliceSerial, t)
}

func TestIsSameSliceParallel(t *testing.T) {
	testSameSliceFut(isSameSliceParallel, t)
}

func TestThereAndBackAgainSha256(t *testing.T) {
	for i := 0; i < testRandomDataNTimes; i++ {
		data := make([]byte, 4096)
		for j := range data {
			data[j] = byte(rand.Int() % 512)
		}
		tmp, err := AddSha256Sum(protogol.Package{Parent: nil, Data: data})
		pkg, err := CheckSha256Sum(tmp)

		if err != nil || !isSameSlice(pkg.Data.([]byte), data) {
			t.Fail()
		}
	}
}

func TestAddSha256(t *testing.T) {
	data := []byte{
		0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x6d, 0x61, 0x69, 0x6e,
		0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x28, 0x0a, 0x09,
		0x22, 0x6f, 0x73, 0x22, 0x0a, 0x09, 0x22, 0x72, 0x75, 0x6e, 0x74, 0x69,
		0x6d, 0x65, 0x22, 0x0a, 0x0a, 0x09, 0x5f, 0x20, 0x22, 0x67, 0x69, 0x74,
		0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x72, 0x67,
		0x6f, 0x64, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62,
		0x61, 0x73, 0x65, 0x36, 0x34, 0x22, 0x0a, 0x09, 0x2e, 0x20, 0x22, 0x67,
		0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61,
		0x72, 0x67, 0x6f, 0x64, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x67, 0x65, 0x6e,
		0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x0a, 0x09,
		0x5f, 0x20, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
		0x6d, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x61, 0x6e, 0x2f, 0x67,
		0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x62, 0x79,
		0x74, 0x65, 0x73, 0x22, 0x0a, 0x0a, 0x09, 0x22, 0x67, 0x69, 0x74, 0x68,
		0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x72, 0x66, 0x61, 0x76,
		0x65, 0x2f, 0x63, 0x6c, 0x69, 0x22, 0x0a, 0x29, 0x0a, 0x0a, 0x63, 0x6f,
		0x6e, 0x73, 0x74, 0x20, 0x41, 0x50, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x20,
		0x3d, 0x20, 0x22, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x2d, 0x64, 0x65, 0x76,
		0x22, 0x0a, 0x0a, 0x76, 0x61, 0x72, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
		0x6e, 0x64, 0x73, 0x20, 0x5b, 0x5d, 0x63, 0x6c, 0x69, 0x2e, 0x43, 0x6f,
		0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x0a, 0x0a, 0x66, 0x75, 0x6e, 0x63, 0x20,
		0x69, 0x6e, 0x69, 0x74, 0x28, 0x29, 0x20, 0x7b, 0x0a, 0x09, 0x72, 0x75,
		0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x47, 0x4f, 0x4d, 0x41, 0x58, 0x50,
		0x52, 0x4f, 0x43, 0x53, 0x28, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
		0x2e, 0x4e, 0x75, 0x6d, 0x43, 0x50, 0x55, 0x28, 0x29, 0x29, 0x0a, 0x7d,
		0x0a, 0x0a, 0x66, 0x75, 0x6e, 0x63, 0x20, 0x6d, 0x61, 0x69, 0x6e, 0x28,
		0x29, 0x20, 0x7b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x20, 0x3a, 0x3d, 0x20,
		0x63, 0x6c, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x70, 0x70, 0x28, 0x29,
		0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x3d,
		0x20, 0x22, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x22, 0x0a, 0x09, 0x61, 0x70,
		0x70, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x20, 0x3d, 0x20, 0x22, 0x41,
		0x20, 0x6e, 0x69, 0x66, 0x74, 0x79, 0x20, 0x74, 0x6f, 0x6f, 0x6c, 0x20,
		0x74, 0x6f, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x20,
		0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x20, 0x66, 0x6f, 0x72,
		0x20, 0x67, 0x6f, 0x2e, 0x22, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x56,
		0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x3d, 0x20, 0x41, 0x50, 0x50,
		0x5f, 0x56, 0x45, 0x52, 0x0a, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x43,
		0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x20, 0x3d, 0x20, 0x43, 0x6f,
		0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x0a, 0x0a, 0x09, 0x61, 0x70, 0x70,
		0x2e, 0x52, 0x75, 0x6e, 0x28, 0x6f, 0x73, 0x2e, 0x41, 0x72, 0x67, 0x73,
		0x29, 0x0a, 0x7d, 0x0a,
	}

	hash := []byte{
		0xb3, 0x11, 0xa0, 0x2c, 0xaf, 0x11, 0xa4, 0x4f, 0xa3, 0x12, 0x8c, 0x84,
		0x92, 0xe6, 0x14, 0xcc, 0x76, 0xc6, 0x69, 0xe5, 0x65, 0xa8, 0xf3, 0x64,
		0x70, 0xc5, 0xc3, 0x23, 0x01, 0x54, 0xc3, 0x15,
	}

	pkg, err := AddSha256Sum(protogol.Package{Parent: nil, Data: data})

	Convey("The hash calculation should be succesfull.", t, func() {
		So(err, ShouldBeNil)
	})

	Convey("The hashed data should match the precalculated hash.", t, func() {
		So(pkg.Data, ShouldResemble, hash)
	})
}

func TestCheckSha256(t *testing.T) {
	data := []byte{
		0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x6d, 0x61, 0x69, 0x6e,
		0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x28, 0x0a, 0x09,
		0x22, 0x6f, 0x73, 0x22, 0x0a, 0x09, 0x22, 0x72, 0x75, 0x6e, 0x74, 0x69,
		0x6d, 0x65, 0x22, 0x0a, 0x0a, 0x09, 0x5f, 0x20, 0x22, 0x67, 0x69, 0x74,
		0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x72, 0x67,
		0x6f, 0x64, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62,
		0x61, 0x73, 0x65, 0x36, 0x34, 0x22, 0x0a, 0x09, 0x2e, 0x20, 0x22, 0x67,
		0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61,
		0x72, 0x67, 0x6f, 0x64, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x67, 0x65, 0x6e,
		0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x0a, 0x09,
		0x5f, 0x20, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
		0x6d, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x61, 0x6e, 0x2f, 0x67,
		0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x62, 0x79,
		0x74, 0x65, 0x73, 0x22, 0x0a, 0x0a, 0x09, 0x22, 0x67, 0x69, 0x74, 0x68,
		0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x72, 0x66, 0x61, 0x76,
		0x65, 0x2f, 0x63, 0x6c, 0x69, 0x22, 0x0a, 0x29, 0x0a, 0x0a, 0x63, 0x6f,
		0x6e, 0x73, 0x74, 0x20, 0x41, 0x50, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x20,
		0x3d, 0x20, 0x22, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x2d, 0x64, 0x65, 0x76,
		0x22, 0x0a, 0x0a, 0x76, 0x61, 0x72, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
		0x6e, 0x64, 0x73, 0x20, 0x5b, 0x5d, 0x63, 0x6c, 0x69, 0x2e, 0x43, 0x6f,
		0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x0a, 0x0a, 0x66, 0x75, 0x6e, 0x63, 0x20,
		0x69, 0x6e, 0x69, 0x74, 0x28, 0x29, 0x20, 0x7b, 0x0a, 0x09, 0x72, 0x75,
		0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x47, 0x4f, 0x4d, 0x41, 0x58, 0x50,
		0x52, 0x4f, 0x43, 0x53, 0x28, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
		0x2e, 0x4e, 0x75, 0x6d, 0x43, 0x50, 0x55, 0x28, 0x29, 0x29, 0x0a, 0x7d,
		0x0a, 0x0a, 0x66, 0x75, 0x6e, 0x63, 0x20, 0x6d, 0x61, 0x69, 0x6e, 0x28,
		0x29, 0x20, 0x7b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x20, 0x3a, 0x3d, 0x20,
		0x63, 0x6c, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x70, 0x70, 0x28, 0x29,
		0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x3d,
		0x20, 0x22, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x22, 0x0a, 0x09, 0x61, 0x70,
		0x70, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x20, 0x3d, 0x20, 0x22, 0x41,
		0x20, 0x6e, 0x69, 0x66, 0x74, 0x79, 0x20, 0x74, 0x6f, 0x6f, 0x6c, 0x20,
		0x74, 0x6f, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x20,
		0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x20, 0x66, 0x6f, 0x72,
		0x20, 0x67, 0x6f, 0x2e, 0x22, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x56,
		0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x3d, 0x20, 0x41, 0x50, 0x50,
		0x5f, 0x56, 0x45, 0x52, 0x0a, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x43,
		0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x20, 0x3d, 0x20, 0x43, 0x6f,
		0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x0a, 0x0a, 0x09, 0x61, 0x70, 0x70,
		0x2e, 0x52, 0x75, 0x6e, 0x28, 0x6f, 0x73, 0x2e, 0x41, 0x72, 0x67, 0x73,
		0x29, 0x0a, 0x7d, 0x0a,
	}

	hash := []byte{
		0xb3, 0x11, 0xa0, 0x2c, 0xaf, 0x11, 0xa4, 0x4f, 0xa3, 0x12, 0x8c, 0x84,
		0x92, 0xe6, 0x14, 0xcc, 0x76, 0xc6, 0x69, 0xe5, 0x65, 0xa8, 0xf3, 0x64,
		0x70, 0xc5, 0xc3, 0x23, 0x01, 0x54, 0xc3, 0x15,
	}

	wrongHash := []byte{
		0x00, 0x11, 0xa0, 0x2c, 0xaf, 0x11, 0xa4, 0x4f, 0xa3, 0x12, 0x8c, 0x84,
		0x92, 0xe6, 0x14, 0xcc, 0x76, 0xc6, 0x69, 0xe5, 0x65, 0xa8, 0xf3, 0x64,
		0x70, 0xc5, 0xc3, 0x23, 0x01, 0x54, 0xc3, 0x15,
	}

	pkg := protogol.NewPackage(data)
	res, err := CheckSha256Sum(pkg.Pack(hash))

	Convey("Package should be unpacked after.", t, func() {
		So(res, ShouldResemble, pkg)
	})

	Convey("Hash should fit.", t, func() {
		So(err, ShouldBeNil)
	})

	res, err = CheckSha256Sum(pkg.Pack(wrongHash))

	Convey("Package should be unpacked after.", t, func() {
		So(res, ShouldResemble, pkg)
	})

	Convey("Hash should be invalid.", t, func() {
		So(err, ShouldNotBeNil)
		So(err, ShouldHaveSameTypeAs, InvalidHashError{})
	})
}
