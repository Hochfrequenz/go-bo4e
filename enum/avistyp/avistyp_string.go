// Code generated by "stringer -type=AvisTyp"; DO NOT EDIT.

package avistyp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ABGELEHNTE_FORDERUNG-1]
	_ = x[ZAHLUNGSAVIS-2]
}

const _AvisTyp_name = "ABGELEHNTE_FORDERUNGZAHLUNGSAVIS"

var _AvisTyp_index = [...]uint8{0, 20, 32}

func (i AvisTyp) String() string {
	i -= 1
	if i < 0 || i >= AvisTyp(len(_AvisTyp_index)-1) {
		return "AvisTyp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AvisTyp_name[_AvisTyp_index[i]:_AvisTyp_index[i+1]]
}
