package tbase

/*
#include <Vset.h>
*/
import "C"

import (
	"unsafe"

	goptr "github.com/mattn/go-pointer"
)

//export tbEnumerateProxy
func tbEnumerateProxy(VsetPtr unsafe.Pointer, user_data unsafe.Pointer) {
	v := goptr.Restore(user_data).(tbaseEnumerator)

	vst := Vset{
		(*C.Vset)(VsetPtr),
		C.GoString((*C.Vset)(VsetPtr).self),
	}

	v.EnumeratorCallback(vst)
}

//export tbReferProxy
func tbReferProxy(VsetPtr unsafe.Pointer, cField unsafe.Pointer, user_data unsafe.Pointer) {
	v := goptr.Restore(user_data).(tbaseReferer)
	field := C.GoString((*C.char)(cField))

	vst := Vset{
		(*C.Vset)(VsetPtr),
		C.GoString((*C.Vset)(VsetPtr).self),
	}

	v.RefererCallback(vst, field)
}
