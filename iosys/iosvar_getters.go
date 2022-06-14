package iosys

/*
#include <PccIosys.h>
*/
import "C"
import (
	"fmt"
	"math"
	"time"
)

// GetBool returns true if value of Iosys-variable is not equal to 0
func (v *IosVar) GetBool() (bool, error) {
	if v.ptr == nil {
		return false, fmt.Errorf("ptr is nil")
	}

	return C.IosVar_get_int(v.ptr) != 0, nil
}

func (v *IosVar) GetInt() (int64, error) {
	if v.ptr == nil {
		return 0, fmt.Errorf("ptr is nil")
	}

	return int64(C.IosVar_get_int(v.ptr)), nil
}

func (v *IosVar) GetDouble() (float64, error) {
	if v.ptr == nil {
		return 0, fmt.Errorf("ptr is nil")
	}

	return float64(C.IosVar_get_real(v.ptr)), nil
}

func (v *IosVar) GetTimestamp() (time.Time, error) {
	if v.ptr == nil {
		return time.Time{}, fmt.Errorf("ptr is nil")
	}

	ts := float64(C.IosVar_get_timestamp(v.ptr))
	p1, p2 := math.Modf(ts)
	t := time.Unix(int64(p1), int64(p2*1000000000))

	return t, nil
}
