// Code generated by "stringer --type Geschaeftspartnerrolle"; DO NOT EDIT.

package geschaeftspartnerrolle

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LIEFERANT-1]
	_ = x[DIENSTLEISTER-2]
	_ = x[KUNDE-3]
	_ = x[INTERESSENT-4]
	_ = x[MARKTPARTNER-5]
}

const _Geschaeftspartnerrolle_name = "LIEFERANTDIENSTLEISTERKUNDEINTERESSENTMARKTPARTNER"

var _Geschaeftspartnerrolle_index = [...]uint8{0, 9, 22, 27, 38, 50}

func (i Geschaeftspartnerrolle) String() string {
	i -= 1
	if i < 0 || i >= Geschaeftspartnerrolle(len(_Geschaeftspartnerrolle_index)-1) {
		return "Geschaeftspartnerrolle(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Geschaeftspartnerrolle_name[_Geschaeftspartnerrolle_index[i]:_Geschaeftspartnerrolle_index[i+1]]
}
