// Code generated by "stringer --type GrundDerPrivilegierungNachEnFG"; DO NOT EDIT.

package grundderprivilegierungnachenfg

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[STROMSPEICHER_UND_VERLUSTENERGIE-1]
	_ = x[ELEKTRISCH_ANGETRIEBENE_WAERMEPUMPEN-2]
	_ = x[UMLAGEERHEBUNG_BEI_ANLAGEN_ZUR_VERSTROMUNG_VON_KUPPELGASEN-3]
	_ = x[HERSTELLUNG_VON_GRUENEN_WASSERSTOFF-4]
	_ = x[STROMKOSTENINTENSIVE_UNTERNEHMEN-5]
	_ = x[HERSTELLUNG_VON_WASSERSTOFF_IN_STROMKOSTENINTENSIVEN_UNTERNEHMEN-6]
	_ = x[SCHIENENBAHNEN-7]
	_ = x[ELEKTRISCHE_BETRIEBENE_BUSSEN_IM_LINIENVERKEHR-8]
	_ = x[LANDSTROMANLAGEN-9]
}

const _GrundDerPrivilegierungNachEnFG_name = "STROMSPEICHER_UND_VERLUSTENERGIEELEKTRISCH_ANGETRIEBENE_WAERMEPUMPENUMLAGEERHEBUNG_BEI_ANLAGEN_ZUR_VERSTROMUNG_VON_KUPPELGASENHERSTELLUNG_VON_GRUENEN_WASSERSTOFFSTROMKOSTENINTENSIVE_UNTERNEHMENHERSTELLUNG_VON_WASSERSTOFF_IN_STROMKOSTENINTENSIVEN_UNTERNEHMENSCHIENENBAHNENELEKTRISCHE_BETRIEBENE_BUSSEN_IM_LINIENVERKEHRLANDSTROMANLAGEN"

var _GrundDerPrivilegierungNachEnFG_index = [...]uint16{0, 32, 68, 126, 161, 193, 257, 271, 317, 333}

func (i GrundDerPrivilegierungNachEnFG) String() string {
	i -= 1
	if i < 0 || i >= GrundDerPrivilegierungNachEnFG(len(_GrundDerPrivilegierungNachEnFG_index)-1) {
		return "GrundDerPrivilegierungNachEnFG(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _GrundDerPrivilegierungNachEnFG_name[_GrundDerPrivilegierungNachEnFG_index[i]:_GrundDerPrivilegierungNachEnFG_index[i+1]]
}
