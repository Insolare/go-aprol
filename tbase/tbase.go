package tbase

/*
#cgo CFLAGS: -I /opt/aprol/include
#cgo LDFLAGS: -L /opt/aprol/lib64 -lStd -lIosys -lAprolPGP -lStdSys -lVset -lAprolLoaderComm -lAprolLc -lTbase

#include <stdlib.h>
#include <PccTbase.h>
#//include <Vset.h>

extern int tb_enumerateCgo(Vset *v, void *user_data);
extern int tb_referCgo(Vset *s, char *field, void *closure);
typedef int (*enumerate_cb)(Vset *s, void *closure);
typedef int (*refer_cb)(Vset *from, char *field, void *closure);
*/
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

type Tbase struct {
	conn *C.TbaseConnection
	base *C.Tbase
}

type tbaseEnumerator interface {
	EnumeratorCallback(v Vset)
}

type tbaseReferer interface {
	RefererCallback(v Vset, field string)
}

type enumerator struct {
	vsets []Vset
}

type referer struct {
	refs []Vset
}

func (r *referer) RefererCallback(v Vset, field string) {
	fmt.Println("Referer callback with field", field)
	r.refs = append(r.refs, v)
}

func (e *enumerator) EnumeratorCallback(v Vset) {
	e.vsets = append(e.vsets, v)
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

func (t *Tbase) Enumerate(path string) []Vset {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	en := &enumerator{
		vsets: make([]Vset, 0),
	}

	handle := cgo.NewHandle(en)
	defer handle.Delete()

	C.tb_enumerate(t.base,
		cPath,
		nil,
		0,
		C.enumerate_cb(unsafe.Pointer(C.tb_enumerateCgo)),
		unsafe.Pointer(&handle))

	return en.vsets
}

func (t *Tbase) GetReferences(path string) []Vset {
	cPath := C.CString(path)

	ref := &referer{
		refs: make([]Vset, 0),
	}

	handle := cgo.NewHandle(ref)
	defer handle.Delete()

	C.tb_refer(t.base,
		cPath,
		nil, nil, nil,
		0,
		C.refer_cb(C.tb_referCgo),
		unsafe.Pointer(&handle))

	C.free(unsafe.Pointer(cPath))

	return ref.refs
}

// Returns Vset containing pointer to underlaying structures
func (t *Tbase) Get(path string) (Vset, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	cVset := C.tb_get(t.base, cPath)
	if cVset == nil {
		return Vset{}, fmt.Errorf("no Vset at given path")
	}

	self := C.GoString(cVset.self)

	result := Vset{
		cVset,
		self,
	}

	return result, nil
}

func (t *Tbase) Exist(path string, option int) (bool, error) {
	var opt C.int
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	switch option {
	case TB_EXIST:
		opt = C.TB_EXIST
	case TB_EXIST_SUB:
		opt = C.TB_EXIST_SUB
	case TB_EXIST_ANY:
		opt = C.TB_EXIST_ANY
	}

	response := C.tb_exist(t.base, cPath, opt)

	if response == C.TB_OKAY {
		return true, nil
	}

	if response == C.TB_NEG {
		return false, nil
	}

	return false, fmt.Errorf("database error: %d", int(response))
}

func (t *Tbase) Disconnect() {
	fmt.Println("Disconnecting tbase")
	if t.base != nil {
		C.tb_close(t.base, nil)
	}

	if t.conn != nil {
		C.tb_disconnect(t.conn)
	}
}
