package iosys

/*
#include <PccIosys.h>
*/
import "C"

func (v *IosVar) SetInt(val int) {
	C.IosVar_set_int(v.ptr, (C.int64_t)(val))
}
