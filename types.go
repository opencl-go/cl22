package cl22

// #include "api.h"
import "C"

// Bool represents a boolean value in the OpenCL API.
// It is not guaranteed to be the same size as the bool in kernels.
type Bool C.cl_bool

const (
	// False is the Bool value representing "false".
	False Bool = C.CL_FALSE
	// True is the Bool value representing "true".
	True Bool = C.CL_TRUE
)

// BoolFrom returns the Bool equivalent of a boolean value.
func BoolFrom(b bool) Bool {
	if b {
		return True
	}
	return False
}

// ToGoBool returns false if the Bool value is False, and true otherwise.
func (b Bool) ToGoBool() bool {
	return b != False
}
