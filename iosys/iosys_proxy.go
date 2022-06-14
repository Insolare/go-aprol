package iosys

/*
#include <PccIosys.h>
*/

import "C"

import (
	"runtime/cgo"
	"unsafe"
)

/*
* Iosys connection events proxies
 */

//export iosysConnectedProxy
func iosysConnectedProxy(closure unsafe.Pointer) {
	handle := *(*cgo.Handle)(closure)
	reciever := handle.Value().(IosysConnectionEventReciever)

	reciever.OnConnected()
}

//export iosysDisconnectedProxy
func iosysDisconnectedProxy(closure unsafe.Pointer) {
	handle := *(*cgo.Handle)(closure)
	reciever := handle.Value().(IosysConnectionEventReciever)

	reciever.OnDisconnected()
}

/*
* IosVar event proxies
 */

//export iosysChangedProxy
func iosysChangedProxy(v unsafe.Pointer, user_data unsafe.Pointer) {
	handle := *(*cgo.Handle)(user_data)
	reciever := handle.Value().(IosVarEvtReciever)

	reciever.OnChange()
}

//export iosysChangeRequestProxy
func iosysChangeRequestProxy(v unsafe.Pointer, user_data unsafe.Pointer) {
	handle := *(*cgo.Handle)(user_data)
	reciever := handle.Value().(IosVarEvtReciever)

	reciever.OnChangeRequest()
}

//export iosysIdlerProxy
func iosysIdlerProxy(v unsafe.Pointer, user_data unsafe.Pointer) {
	handle := *(*cgo.Handle)(user_data)
	reciever := handle.Value().(IosVarEvtReciever)

	reciever.OnIdleChange()
}
