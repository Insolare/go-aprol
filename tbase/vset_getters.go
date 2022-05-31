package tbase

/*
#cgo CFLAGS: -I /opt/aprol/include
#include <Vset.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func (v *Vset) GetString(name string) (string, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cResult := C.VsetGetString(v.vset, cName)

	if cResult != nil {
		return C.GoString(cResult), nil
	}

	return "", fmt.Errorf("field not found")
}

func (v *Vset) GetInt(name string) (int, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cResult := C.VsetGetInt(v.vset, cName)

	if cResult != nil {
		return int(*cResult), nil
	}

	return 0, fmt.Errorf("field not found")
}
