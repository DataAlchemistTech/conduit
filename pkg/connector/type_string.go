// Code generated by "stringer -type=Type -trimprefix Type"; DO NOT EDIT.

package connector

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TypeSource-1]
	_ = x[TypeDestination-2]
}

const _Type_name = "SourceDestination"

var _Type_index = [...]uint8{0, 6, 17}

func (i Type) String() string {
	i -= 1
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
