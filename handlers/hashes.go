package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"

	"github.com/targodan/protogol"
)

type InvalidHashError struct {
	Expected []byte
	Received []byte
}

func (err InvalidHashError) Error() string {
	return "Invalid Hash. \"" + base64.StdEncoding.EncodeToString(err.Expected) + "\" != \"" + base64.StdEncoding.EncodeToString(err.Received) + "\""
}

func isSameSlice(a []byte, b []byte) bool {
	return isSameSliceSerial(a, b)
}

func isSameSliceSerial(a []byte, b []byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func isSameSliceParallel(a []byte, b []byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	const numParts = 3
	slicePartLen := len(a) / numParts
	c := make(chan bool, numParts)
	for i := 0; i < numParts; i++ {
		go func(c chan bool, i int) {
			limit := (i + 1) * slicePartLen
			if limit > len(a)-1 {
				limit = len(a) - 1
			}
			for j := i * slicePartLen; j < limit; j++ {
				if a[j] != b[j] {
					c <- false
					return
				}
			}
			c <- true
		}(c, i)
	}

	for i := 0; i < numParts; i++ {
		if !(<-c) {
			return false
		}
	}

	return true
}

func AddSha256Sum(pkg protogol.Package) (protogol.Package, error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, pkg.Data)
	hash := sha256.Sum256(buf.Bytes())
	return protogol.Package{Parent: &pkg, Data: hash[:]}, nil
}

func CheckSha256Sum(pkg protogol.Package) (ret protogol.Package, err error) {
	ret = protogol.Unpack(pkg)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, pkg.Parent.Data)
	hash := sha256.Sum256(buf.Bytes())
	if !isSameSlice(hash[:], pkg.Data.([]byte)) {
		err = InvalidHashError{Expected: hash[:], Received: pkg.Data.([]byte)}
	}

	return
}
