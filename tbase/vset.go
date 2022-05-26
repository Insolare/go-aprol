package tbase

/*
#cgo CFLAGS: -I /opt/aprol/include
#cgo LDFLAGS: -L /opt/aprol/lib64 -lStd -lIosys -lAprolPGP -lStdSys -lVset -lAprolLoaderComm -lAprolLc -lTbase

#include <stdlib.h>
#include <Vset.h>

Vt_TYPE getVtType(Vt *vt) {
	return vt->type;
}

*/
import "C"

type Vset struct {
	vset     *C.Vset
	Self     string
	Children []Vset
}

// Disposes resources allocated at underlying library.
// If Children exi
//
func (v *Vset) Free() {
	if v.vset != nil {
		C.VsetFree(v.vset)
	}

	for _, c := range v.Children {
		if c.vset != nil {
			C.VsetFree(c.vset)
		}
	}

	v.vset = nil
}

// Fill Children slice with Vsets if they exist
//
// All child-Vsets may also have children!
func (v *Vset) EnumeratorCallback(a Vset, b interface{}) {
	v.Children = append(v.Children, a)
}

func (v *Vset) ListFields() <-chan string {
	var vt *C.Vt
	ch := make(chan string)

	go func() {
		for i := 0; i < int(v.vset.nmax)-1; i++ {
			vt = C.VsetAt(v.vset, C.int(i))
			ch <- C.GoString(C.VtTypename(C.getVtType(vt)))
		}

		close(ch)
	}()

	return ch
}
