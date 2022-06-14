package tbase

/*
#include <Vset.h>
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

//export tbEnumerateProxy
func tbEnumerateProxy(VsetPtr unsafe.Pointer, user_data unsafe.Pointer) {
	handle := *(*cgo.Handle)(user_data)
	reciever := handle.Value().(tbaseEnumerator)

	vst := Vset{
		(*C.Vset)(VsetPtr),
		C.GoString((*C.Vset)(VsetPtr).self),
	}

	reciever.EnumeratorCallback(vst)
}

//export tbReferProxy
func tbReferProxy(VsetPtr unsafe.Pointer, cField unsafe.Pointer, user_data unsafe.Pointer) {
	handle := *(*cgo.Handle)(user_data)
	reciever := handle.Value().(tbaseReferer)

	field := C.GoString((*C.char)(cField))

	vst := Vset{
		(*C.Vset)(VsetPtr),
		C.GoString((*C.Vset)(VsetPtr).self),
	}

	reciever.RefererCallback(vst, field)
}
