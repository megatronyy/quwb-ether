package main

import (
	"unsafe"
	"reflect"
)

var (
	s   = make([]byte, 20)
	ptr = unsafe.Pointer(&s[0])
	//ptr    unsafe.Pointer
	length = 20
)

var s1 = struct {
	addr uintptr
	len  int
	cap  int
}{uintptr(ptr), length, length}

func main() {
	var o []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	sliceHeader.Len = 20
	sliceHeader.Cap = 20
	sliceHeader.Data = uintptr(ptr)
}
