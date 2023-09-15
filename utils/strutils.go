package utils

import (
	"reflect"
	"unsafe"
)

func ByteToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StrToByte(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}
