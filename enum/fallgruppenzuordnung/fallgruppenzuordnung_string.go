// Code generated by "stringer --type Fallgruppenzuordnung"; DO NOT EDIT.

package fallgruppenzuordnung

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GABI_RLMmT-1]
	_ = x[GABI_RLMoT-2]
	_ = x[GABI_RLMNEV-3]
}

const _Fallgruppenzuordnung_name = "GABI_RLMmTGABI_RLMoTGABI_RLMNEV"

var _Fallgruppenzuordnung_index = [...]uint8{0, 10, 20, 31}

func (i Fallgruppenzuordnung) String() string {
	i -= 1
	if i < 0 || i >= Fallgruppenzuordnung(len(_Fallgruppenzuordnung_index)-1) {
		return "Fallgruppenzuordnung(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Fallgruppenzuordnung_name[_Fallgruppenzuordnung_index[i]:_Fallgruppenzuordnung_index[i+1]]
}
