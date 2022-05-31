package iosys

/*
#cgo CFLAGS: -I /opt/aprol/include
#cgo LDFLAGS: -L /opt/aprol/lib64 -lStd -lIosys -lAprolPGP -lStdSys -lVset -lAprolLoaderComm -lAprolLc -lTbase

#include <stdlib.h>
#include <PccIosys.h>

typedef void (*changed_fct)(IosVar *s, void *closure);
typedef void (*changerequest_fct)(IosVar *s, void *closure);
typedef void (*idler_fct)(IosVar *s, void *closure);

extern void goaprol_ios_mainloop();
*/
import "C"

import (
	"context"
	"sync"
	"unsafe"
)

type EventReciever interface {
	OnConnected()
	OnDisconnected()
}

type IosysConnection struct {
	conn      *C.IosConn
	callbacks EventReciever
	vars      []unsafe.Pointer
}

func Initialize() {
	C.Ios_init()
}

func Finalize() {
	C.Ios_fin()
}

func NewIosysConnection(host string) IosysConnection {
	cHost := C.CString(host)
	defer C.free(unsafe.Pointer(cHost))

	return IosysConnection{
		C.IosConn_new(cHost),
		nil,
		make([]unsafe.Pointer, 0),
	}
}

func (cn *IosysConnection) Connect(reconnect int) {
	C.IosConn_connect(cn.conn, C.int(reconnect))
	C.Ios_sync()
}

func (cn *IosysConnection) IsConnected() bool {
	r := C.IosConn_connected_p(cn.conn)

	return r > 0
}

func (cn *IosysConnection) Delete() {
	C.IosConn_delete(cn.conn)
}

// StartMainLoop launches Iosys-mainloop via underlying library.
func StartMainloop(wg *sync.WaitGroup) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				C.goaprol_ios_mainloop()
			}
		}
	}()

	return cancel
}
