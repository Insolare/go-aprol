package tbase

/*
#cgo CFLAGS: -I /opt/aprol/include

#include <stdlib.h>
#include <Vset.h>

Vt_TYPE getVtType(Vt *vt) {
	return vt->type;
}

*/
import "C"

type Vset struct {
	vset *C.Vset
	Self string
}

// Free disposes resources allocated at underlying library.
func (v *Vset) Free() {
	if v.vset != nil {
		C.VsetFree(v.vset)
	}

	v.vset = nil
}
