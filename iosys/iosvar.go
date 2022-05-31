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
	"unsafe"

	goptr "github.com/mattn/go-pointer"
)

// Must be embedded to user type
type IosVar struct {
	ptr   *C.IosVar
	owned bool
}

type IosVarEvtReciever interface {
	OnChange()
	OnChangeRequest()
	OnIdleChange()
}

func NewIosVar(name string, i IosVarEvtReciever) IosVar {
	cName := C.CString(name)

	p := goptr.Save(i)

	pVar := C.IosVar_new(cName, &C.IosCallbacks, p)
	C.Ios_sync()

	C.free(unsafe.Pointer(cName))

	return IosVar{
		ptr: pVar,
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

func (v *IosVar) SetValid() {
	if v.owned {
		C.IosVar_set_valid(v.ptr)
	}
}

func (v *IosVar) SetInvalid() {
	if v.owned {
		C.IosVar_invalidate(v.ptr)
	}
}

func (v *IosVar) Delete() {
	C.IosVar_delete(v.ptr)

	v.ptr = nil
}
