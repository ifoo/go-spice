// Code generated by "stringer -type=Capability"; DO NOT EDIT.

package red

import "fmt"

const _Capability_name = "CapabilityAuthSpiceCapabilityAuthSASL"

var _Capability_index = [...]uint8{0, 19, 37}

func (i Capability) String() string {
	i -= 1
	if i >= Capability(len(_Capability_index)-1) {
		return fmt.Sprintf("Capability(%d)", i+1)
	}
	return _Capability_name[_Capability_index[i]:_Capability_index[i+1]]
}