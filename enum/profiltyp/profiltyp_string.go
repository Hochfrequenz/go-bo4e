// Code generated by "stringer --type Profiltyp"; DO NOT EDIT.

package profiltyp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SLP_SEP-1]
	_ = x[TLP_TEP-2]
}

const _Profiltyp_name = "SLP_SEPTLP_TEP"

var _Profiltyp_index = [...]uint8{0, 7, 14}

func (i Profiltyp) String() string {
	i -= 1
	if i < 0 || i >= Profiltyp(len(_Profiltyp_index)-1) {
		return "Profiltyp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Profiltyp_name[_Profiltyp_index[i]:_Profiltyp_index[i+1]]
}
