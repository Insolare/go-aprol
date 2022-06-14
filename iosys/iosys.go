package iosys

/*
#cgo CFLAGS: -I /opt/aprol/include
#cgo LDFLAGS: -L /opt/aprol/lib64 -lStd -lIosys -lAprolPGP -lStdSys -lVset -lAprolLoaderComm -lAprolLc -lTbase

#include <stdlib.h>
#include <PccIosys.h>

typedef void (*changed_fct)(IosVar *s, void *closure);
typedef void (*changerequest_fct)(IosVar *s, void *closure);
typedef void (*idler_fct)(IosVar *s, void *closure);
typedef void (*ios_conn_fct)(void* conn, int sock, void* closure);
typedef void (*ios_fail_fct)(void* conn, void* closure);

extern void iosys_connectedCgo(void* conn, int sock, void* closure);
extern void iosys_disconnectedCgo(void* conn, void* closure);
extern void goaprol_ios_mainloop();
*/
import "C"

import (
	"context"
	"runtime/cgo"
	"sync"
	"unsafe"
)

type IosysConnectionEventReciever interface {
	OnConnected()
	OnDisconnected()
}

type IosysConnection struct {
	conn   *C.IosConn
	handle cgo.Handle
}

func Initialize() {
	C.Ios_init()
}

func Finalize() {
	C.Ios_fin()
}

func NewIosysConnection(host string, object IosysConnectionEventReciever) IosysConnection {
	cHost := C.CString(host)
	defer C.free(unsafe.Pointer(cHost))

	handle := cgo.NewHandle(object)

	conn := C.IosConn_new(cHost)
	C.IosConn_set_conn(conn, (C.ios_conn_fct)(C.iosys_connectedCgo), unsafe.Pointer(&handle))
	C.IosConn_set_fail(conn, (C.ios_fail_fct)(C.iosys_disconnectedCgo), unsafe.Pointer(&handle))

	return IosysConnection{
		conn:   conn,
		handle: handle,
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
	cn.conn = nil
	cn.handle.Delete()
}

// StartMainLoop launches Iosys-mainloop via underlying library.
func StartMainloop(wg *sync.WaitGroup) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		//runtime.LockOSThread()
		//defer runtime.UnlockOSThread()
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
