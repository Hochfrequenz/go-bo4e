// Code generated by "stringer --type Rufnummernart"; DO NOT EDIT.

package rufnummernart

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RUF_ZENTRALE-1]
	_ = x[FAX_ZENTRALE-2]
	_ = x[SAMMELRUF-3]
	_ = x[SAMMELFAX-4]
	_ = x[ABTEILUNGRUF-5]
	_ = x[ABTEILUNGFAX-6]
	_ = x[RUF_DURCHWAHL-7]
	_ = x[FAX_DURCHWAHL-8]
	_ = x[MOBIL_NUMMER-9]
}

const _Rufnummernart_name = "RUF_ZENTRALEFAX_ZENTRALESAMMELRUFSAMMELFAXABTEILUNGRUFABTEILUNGFAXRUF_DURCHWAHLFAX_DURCHWAHLMOBIL_NUMMER"

var _Rufnummernart_index = [...]uint8{0, 12, 24, 33, 42, 54, 66, 79, 92, 104}

func (i Rufnummernart) String() string {
	i -= 1
	if i < 0 || i >= Rufnummernart(len(_Rufnummernart_index)-1) {
		return "Rufnummernart(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Rufnummernart_name[_Rufnummernart_index[i]:_Rufnummernart_index[i+1]]
}
