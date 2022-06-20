package iosys

/*
#include <PccIosys.h>
*/
import "C"
import (
	"math"
	"time"
)

// GetBool returns true if value of Iosys-variable is not equal to 0
func (v *IosVar) GetBool() bool {
	return C.IosVar_get_int(v.ptr) != 0
}

func (v *IosVar) GetInt() int64 {
	return int64(C.IosVar_get_int(v.ptr))
}

func (v *IosVar) GetDouble() float64 {
	return float64(C.IosVar_get_real(v.ptr))
}

func (v *IosVar) GetTimestamp() time.Time {
	ts := float64(C.IosVar_get_timestamp(v.ptr))
	p1, p2 := math.Modf(ts)
	t := time.Unix(int64(p1), int64(p2*1000000000))

	return t
}
