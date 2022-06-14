package iosys

/*
#cgo CFLAGS: -I /opt/aprol/include

#include <stdlib.h>
#include <PccIosys.h>

extern void iosys_changedCgo(IosVar *v, void *user_data);
extern void  iosys_changerequestCgo(IosVar *v, void *user_data, IosVar *request);
extern void  iosys_idlerCgo(IosVar *v, void *user_data, int idle);
IosFct IosCallbacks = { iosys_changedCgo, iosys_changerequestCgo, iosys_idlerCgo };
*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

// Must be embedded to user type
type IosVar struct {
	ptr    *C.IosVar
	owned  bool
	parent *cgo.Handle
}

type IosVarListener interface {
	OnChange()
}

type IosVarProvider interface {
	OnChangeRequest()
	OnIdleChange()
}

type IosVarEvtReciever interface {
	OnChange()
	OnChangeRequest()
	OnIdleChange()
}

func NewIosVar(name string, i IosVarEvtReciever) IosVar {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	// Clean it while deleting
	handle := cgo.NewHandle(i)

	pVar := C.IosVar_new(cName, &C.IosCallbacks, unsafe.Pointer(&handle))
	C.Ios_sync() // Maybe make it separate?..
	fmt.Println("IosVar created")

	return IosVar{
		ptr:    pVar,
		owned:  false,
		parent: &handle,
	}
}

func (v *IosVar) Source() error {
	ret := C.IosVar_source(v.ptr)

	if ret == 1 {
		return fmt.Errorf("provider already exists")
	}

	if ret == -1 {
		return fmt.Errorf("invalid argument")
	}

	return nil
}

func (v *IosVar) Unsource() error {
	// TODO:
	// -1 - invalid arg
	// 1 - not a provider
	// 0 - ok
	C.IosVar_unsource(v.ptr)

	return nil
}

// SetValid sets Valid flag to Iosys variable if variable is sourced by our app
func (v *IosVar) SetValid() {
	if v.owned {
		C.IosVar_set_valid(v.ptr)
	}
}

// SetInvalid sets Invalid flag to Iosys variable if variable is sourced by our app
func (v *IosVar) SetInvalid() {
	if v.owned {
		C.IosVar_invalidate(v.ptr)
	}
}

// Delete frees resources both in Iosys and go/cgo
func (v *IosVar) Delete() {
	C.IosVar_delete(v.ptr)
	cgo.Handle.Delete(*v.parent)

	C.Ios_sync() // Maybe not needed

	v.ptr = nil
}
