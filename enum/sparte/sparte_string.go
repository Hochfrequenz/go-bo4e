// Code generated by "stringer --type Sparte"; DO NOT EDIT.

package sparte

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[STROM-1]
	_ = x[GAS-2]
	_ = x[FERNWAERME-3]
	_ = x[NAHWAERME-4]
	_ = x[WASSER-5]
	_ = x[ABWASSER-6]
}

const _Sparte_name = "STROMGASFERNWAERMENAHWAERMEWASSERABWASSER"

var _Sparte_index = [...]uint8{0, 5, 8, 18, 27, 33, 41}

func (i Sparte) String() string {
	i -= 1
	if i < 0 || i >= Sparte(len(_Sparte_index)-1) {
		return "Sparte(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Sparte_name[_Sparte_index[i]:_Sparte_index[i+1]]
}
