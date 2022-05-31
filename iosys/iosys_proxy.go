package iosys

/*
#include <PccIosys.h>

extern void iosysChangedProxy(void*, void*);
extern void iosysChangeRequestProxy(void*, void*);
extern void iosysIdlerProxy(void*, void*);
*/
import "C"

import (
	"unsafe"

	goptr "github.com/mattn/go-pointer"
)

//export iosysChangedProxy
func iosysChangedProxy(v unsafe.Pointer, user_data unsafe.Pointer) {
	cb := goptr.Restore(user_data).(IosVarEvtReciever)

	cb.OnChange()
}

//export iosysChangeRequestProxy
func iosysChangeRequestProxy(v unsafe.Pointer, user_data unsafe.Pointer) {
	cb := goptr.Restore(user_data).(IosVarEvtReciever)

	cb.OnChangeRequest()
}

//export iosysIdlerProxy
func iosysIdlerProxy(v unsafe.Pointer, user_data unsafe.Pointer) {
	cb := goptr.Restore(user_data).(IosVarEvtReciever)

	cb.OnIdleChange()
}
