// Code generated by "stringer --type Marktrolle"; DO NOT EDIT.

package marktrolle

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NB-1]
	_ = x[LF-2]
	_ = x[MSB-3]
	_ = x[DL-4]
	_ = x[BKV-5]
	_ = x[BKO-6]
	_ = x[UENB-7]
	_ = x[KUNDE_SELBST_NN-8]
	_ = x[MGV-9]
	_ = x[EIV-10]
	_ = x[RB-11]
	_ = x[KUNDE-12]
	_ = x[INTERESSENT-13]
	_ = x[GMSB-14]
}

const _Marktrolle_name = "NBLFMSBDLBKVBKOUENBKUNDE_SELBST_NNMGVEIVRBKUNDEINTERESSENTGMSB"

var _Marktrolle_index = [...]uint8{0, 2, 4, 7, 9, 12, 15, 19, 34, 37, 40, 42, 47, 58, 62}

func (i Marktrolle) String() string {
	i -= 1
	if i < 0 || i >= Marktrolle(len(_Marktrolle_index)-1) {
		return "Marktrolle(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Marktrolle_name[_Marktrolle_index[i]:_Marktrolle_index[i+1]]
}
