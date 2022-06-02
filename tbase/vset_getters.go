package tbase

/*
#cgo CFLAGS: -I /opt/aprol/include
#include <Vset.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

var ErrFieldNotFound = errors.New("field not found")

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

func (v *Vset) GetBoolean(name string) (bool, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cResult := C.VsetGetBoolean(v.vset, cName)
	if cResult != nil {
		if *cResult == 1 {
			return true, nil
		}

		return false, nil
	}

	return false, ErrFieldNotFound
}
