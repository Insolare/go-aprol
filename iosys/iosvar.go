package iosys

/*
#cgo CFLAGS: -I /opt/aprol/include

#include <stdlib.h>
#include <PccIosys.h>

extern void iosys_changedCgo(IosVar *v, void *user_data);
extern void  iosys_changerequestCgo(IosVar *v, void *user_data, IosVar *request);
extern void  iosys_idlerCgo(IosVar *v, void *user_data, int idle);
IosFct IosReaderCallbacks = { iosys_changedCgo, NULL, NULL };
IosFct IosWriterCallbacks = { NULL, iosys_changerequestCgo, iosys_idlerCgo };
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

// NewIosVar creates new (or gets existing) variable in Iosys/
//
// `reciever` must implement either IosVarListener (for read-only variable) or IosVarProvider (for sourced variables)
func NewIosVar(name string, reciever interface{}) (*IosVar, error) {
	// Interface validation
	var callbacks *C.IosFct
	switch reciever.(type) {
	case IosVarListener:
		callbacks = &C.IosReaderCallbacks
	case IosVarProvider:
		callbacks = &C.IosWriterCallbacks
	default:
		return nil, fmt.Errorf("goaprol: unsupported i argument")
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	// Clean it while deleting
	handle := cgo.NewHandle(reciever)

	pVar := C.IosVar_new(cName, callbacks, unsafe.Pointer(&handle))

	return &IosVar{
		ptr:    pVar,
		owned:  false,
		parent: &handle,
	}, nil
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
func (v *IosVar) SetValid(valid bool) {
	if v.owned {
		if valid {
			C.IosVar_set_valid(v.ptr)
		} else {
			C.IosVar_invalidate(v.ptr)
		}
	}
}

// Delete frees resources both in Iosys and go/cgo
func (v *IosVar) Delete() {
	C.IosVar_delete(v.ptr)
	cgo.Handle.Delete(*v.parent)

	v.parent = nil
	v.ptr = nil
}
