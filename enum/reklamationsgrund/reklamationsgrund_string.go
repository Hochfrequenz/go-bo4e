// Code generated by "stringer --type Reklamationsgrund"; DO NOT EDIT.

package reklamationsgrund

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FEHLT-1]
	_ = x[ZU_NIEDRIG-2]
	_ = x[ZU_HOCH-3]
}

const _Reklamationsgrund_name = "FEHLTZU_NIEDRIGZU_HOCH"

var _Reklamationsgrund_index = [...]uint8{0, 5, 15, 22}

func (i Reklamationsgrund) String() string {
	i -= 1
	if i < 0 || i >= Reklamationsgrund(len(_Reklamationsgrund_index)-1) {
		return "Reklamationsgrund(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Reklamationsgrund_name[_Reklamationsgrund_index[i]:_Reklamationsgrund_index[i+1]]
}
