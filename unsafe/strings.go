// +build !appengine

package unsafe

import (
	"reflect"
	"unsafe"
)

// ByteSlice returns s as a byte slice without copying,
// any attempt to modifythe resulting slice might end all life as we know.
func ByteSlice(s *string) []byte {
	return VoidToBytes(unsafe.Pointer(s), len(*s))
}

// VoidToBytes is a helper function to convert a void* buffer from a cgo call
// to a byte slice
func VoidToBytes(p unsafe.Pointer, ln int) []byte {
	sh := &reflect.SliceHeader{
		Data: uintptr(p),
		Len:  ln,
		Cap:  ln,
	}
	return *(*[]byte)(unsafe.Pointer(sh))
}
