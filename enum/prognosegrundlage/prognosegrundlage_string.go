// Code generated by "stringer --type Prognosegrundlage"; DO NOT EDIT.

package prognosegrundlage

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WERTE-1]
	_ = x[PROFILE-2]
}

const _Prognosegrundlage_name = "WERTEPROFILE"

var _Prognosegrundlage_index = [...]uint8{0, 5, 12}

func (i Prognosegrundlage) String() string {
	i -= 1
	if i < 0 || i >= Prognosegrundlage(len(_Prognosegrundlage_index)-1) {
		return "Prognosegrundlage(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Prognosegrundlage_name[_Prognosegrundlage_index[i]:_Prognosegrundlage_index[i+1]]
}
