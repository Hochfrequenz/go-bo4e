// Code generated by "stringer --type Waermenutzung"; DO NOT EDIT.

package waermenutzung

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SPEICHERHEIZUNG-1]
	_ = x[WAERMEPUMPE-2]
	_ = x[DIREKTHEIZUNG-3]
}

const _Waermenutzung_name = "SPEICHERHEIZUNGWAERMEPUMPEDIREKTHEIZUNG"

var _Waermenutzung_index = [...]uint8{0, 15, 26, 39}

func (i Waermenutzung) String() string {
	i -= 1
	if i < 0 || i >= Waermenutzung(len(_Waermenutzung_index)-1) {
		return "Waermenutzung(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Waermenutzung_name[_Waermenutzung_index[i]:_Waermenutzung_index[i+1]]
}
