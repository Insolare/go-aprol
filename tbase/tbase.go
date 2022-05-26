package tbase

/*
#cgo CFLAGS: -I /opt/aprol/include
#cgo LDFLAGS: -L /opt/aprol/lib64 -lStd -lIosys -lAprolPGP -lStdSys -lVset -lAprolLoaderComm -lAprolLc -lTbase

#include <stdlib.h>
#include <PccTbase.h>
#//include <Vset.h>

extern int tb_enumerateCgo(Vset *v, void *user_data);
typedef int (*enumcb)(Vset *s, void *closure);
*/
import "C"

import (
	"fmt"
	"unsafe"

	goptr "github.com/mattn/go-pointer"
)

type Tbase struct {
	conn *C.TbaseConnection
	base *C.Tbase
}

type TbaseEnumerator interface {
	EnumeratorCallback(a Vset, b interface{})
}

func (t *Tbase) Connect(host string) error {
	cHost := C.CString(host)
	defer C.free(unsafe.Pointer(cHost))

	t.conn = C.tb_connect(cHost)
	if t.conn == nil {
		return fmt.Errorf("could not connect to Tbase")
	}

	return nil
}

func (t *Tbase) OpenDatabase(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	if t.conn == nil {
		return fmt.Errorf("not connected to Tbase server")
	}

	dbOpenError := C.tb_open(cPath, t.conn, 2050, 0, &t.base)

	if dbOpenError != C.TB_OKAY {
		return fmt.Errorf("error opening database: %d", dbOpenError)
	}

	return nil
}

func (t *Tbase) CloseDatabase() error {
	if t.base == nil {
		return fmt.Errorf("database not opened")
	}

	dbCloseError := C.tb_close(t.base, nil)

	if dbCloseError != C.TB_OKAY {
		return fmt.Errorf("error closing database: %d", dbCloseError)
	}

	return nil
}

func (t *Tbase) Enumerate(path string, en TbaseEnumerator) {
	p := goptr.Save(en)
	defer goptr.Unref(p)

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	C.tb_enumerate(t.base, cPath, nil, 0, C.enumcb(unsafe.Pointer(C.tb_enumerateCgo)), p)
}

func (t *Tbase) Disconnect() {
	if t.base != nil {
		C.tb_close(t.base, nil)
	}

	if t.conn != nil {
		C.tb_disconnect(t.conn)
	}
}
