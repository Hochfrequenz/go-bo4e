// Code generated by "stringer --type Plausibilisierungshinweis"; DO NOT EDIT.

package plausibilisierungshinweis

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Z83_KUNDENSELBSTABLESUNG-1]
	_ = x[Z84_LEERSTAND-2]
	_ = x[Z85_REALERZAEHLERUEBERLAUFGEPRUEFT-3]
	_ = x[Z86_PLAUSIBELWGKONTROLLABLESUNG-4]
	_ = x[Z87_PLAUSIBELWGKUNDENHINWEIS-5]
	_ = x[ZC3_AUSTAUSCHDESERSATZWERTES-6]
	_ = x[ZR5_RECHENWERT-7]
}

const _Plausibilisierungshinweis_name = "Z83_KUNDENSELBSTABLESUNGZ84_LEERSTANDZ85_REALERZAEHLERUEBERLAUFGEPRUEFTZ86_PLAUSIBELWGKONTROLLABLESUNGZ87_PLAUSIBELWGKUNDENHINWEISZC3_AUSTAUSCHDESERSATZWERTESZR5_RECHENWERT"

var _Plausibilisierungshinweis_index = [...]uint8{0, 24, 37, 71, 102, 130, 158, 172}

func (i Plausibilisierungshinweis) String() string {
	i -= 1
	if i < 0 || i >= Plausibilisierungshinweis(len(_Plausibilisierungshinweis_index)-1) {
		return "Plausibilisierungshinweis(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Plausibilisierungshinweis_name[_Plausibilisierungshinweis_index[i]:_Plausibilisierungshinweis_index[i+1]]
}
