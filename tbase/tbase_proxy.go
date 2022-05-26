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
	v := goptr.Restore(user_data).(TbaseEnumerator)

	vst := Vset{
		(*C.Vset)(VsetPtr),
		C.GoString((*C.Vset)(VsetPtr).self),
		make([]Vset, 0),
	}

	v.EnumeratorCallback(vst, nil)
}
