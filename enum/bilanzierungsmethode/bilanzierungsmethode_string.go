// Code generated by "stringer --type Bilanzierungsmethode"; DO NOT EDIT.

package bilanzierungsmethode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RLM-1]
	_ = x[SLP-2]
	_ = x[TLP_GEMEINSAM-3]
	_ = x[TLP_GETRENNT-4]
	_ = x[PAUSCHAL-5]
}

const _Bilanzierungsmethode_name = "RLMSLPTLP_GEMEINSAMTLP_GETRENNTPAUSCHAL"

var _Bilanzierungsmethode_index = [...]uint8{0, 3, 6, 19, 31, 39}

func (i Bilanzierungsmethode) String() string {
	i -= 1
	if i < 0 || i >= Bilanzierungsmethode(len(_Bilanzierungsmethode_index)-1) {
		return "Bilanzierungsmethode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Bilanzierungsmethode_name[_Bilanzierungsmethode_index[i]:_Bilanzierungsmethode_index[i+1]]
}