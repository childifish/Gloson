package under

import "unsafe"

type TypeChanger struct {
}

func (tc *TypeChanger)Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func (tc *TypeChanger)Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

