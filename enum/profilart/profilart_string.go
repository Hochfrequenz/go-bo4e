// Code generated by "stringer --type Profilart"; DO NOT EDIT.

package profilart

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ART_STANDARDLASTPROFIL-1]
	_ = x[ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL-2]
	_ = x[ART_LASTPROFIL-3]
}

const _Profilart_name = "ART_STANDARDLASTPROFILART_TAGESPARAMETERABHAENGIGES_LASTPROFILART_LASTPROFIL"

var _Profilart_index = [...]uint8{0, 22, 62, 76}

func (i Profilart) String() string {
	i -= 1
	if i < 0 || i >= Profilart(len(_Profilart_index)-1) {
		return "Profilart(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Profilart_name[_Profilart_index[i]:_Profilart_index[i+1]]
}
