// Code generated by "stringer --type GrundlageEnergiemenge"; DO NOT EDIT.

package grundlageenergiemenge

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT-1]
	_ = x[ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT-2]
	_ = x[ZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG-3]
	_ = x[ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG-4]
}

const _GrundlageEnergiemenge_name = "ZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERTZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERTZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNGZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG"

var _GrundlageEnergiemenge_index = [...]uint16{0, 78, 155, 242, 328}

func (i GrundlageEnergiemenge) String() string {
	i -= 1
	if i < 0 || i >= GrundlageEnergiemenge(len(_GrundlageEnergiemenge_index)-1) {
		return "GrundlageEnergiemenge(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _GrundlageEnergiemenge_name[_GrundlageEnergiemenge_index[i]:_GrundlageEnergiemenge_index[i+1]]
}
