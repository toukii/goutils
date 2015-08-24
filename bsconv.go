package goutils

import (
	"reflect"
	"unsafe"
)

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ToByte(v string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&v))
	return *(*[]byte)(unsafe.Pointer(sh.Data))
}
