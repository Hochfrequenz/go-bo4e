// Code generated by "stringer --type RechnungspositionsAbschlag"; DO NOT EDIT.

package rechnungspositionsabschlag

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GEMEINDERABATT-1]
	_ = x[ABSCHLAG_ANPASSUNG-2]
}

const _RechnungspositionsAbschlag_name = "GEMEINDERABATTABSCHLAG_ANPASSUNG"

var _RechnungspositionsAbschlag_index = [...]uint8{0, 14, 32}

func (i RechnungspositionsAbschlag) String() string {
	i -= 1
	if i < 0 || i >= RechnungspositionsAbschlag(len(_RechnungspositionsAbschlag_index)-1) {
		return "RechnungspositionsAbschlag(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RechnungspositionsAbschlag_name[_RechnungspositionsAbschlag_index[i]:_RechnungspositionsAbschlag_index[i+1]]
}
