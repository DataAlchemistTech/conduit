// Code generated by "stringer -type=Status -trimprefix Status"; DO NOT EDIT.

package pipeline

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusRunning-1]
	_ = x[StatusSystemStopped-2]
	_ = x[StatusUserStopped-3]
	_ = x[StatusDegraded-4]
}

const _Status_name = "RunningSystemStoppedUserStoppedDegraded"

var _Status_index = [...]uint8{0, 7, 20, 31, 39}

func (i Status) String() string {
	i -= 1
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
