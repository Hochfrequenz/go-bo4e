// Code generated by "stringer --type Anfragekategorie"; DO NOT EDIT.

package anfragekategorie

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PROZESSDATENBERICHT-1]
	_ = x[GERAETEUEBERNAHME-2]
	_ = x[WEITERVERPFLICHTUNG_BETRIEB_MELO-3]
	_ = x[AENDERUNG_MELO-4]
	_ = x[STAMMDATEN_MALO_ODER_MELO-5]
	_ = x[BILANZIERTE_MENGE_MEHR_MINDER_MENGEN-6]
	_ = x[ALLOKATIONSLISTE_MEHR_MINDER_MENGEN-7]
	_ = x[ENERGIEMENGE_UND_LEISTUNGSMAXIMUM-8]
	_ = x[ABRECHNUNG_MESSSTELLENBETRIEB_MSB_AN_LF-9]
	_ = x[AENDERUNG_PROGNOSEGRUNDLAGE_GERAETEKONFIGURATION-10]
	_ = x[AENDERUNG_GERAETEKONFIGURATION-11]
	_ = x[REKLAMATION_VON_WERTEN-12]
	_ = x[LASTGANG_MALO_TRANCHE-13]
	_ = x[SPERRUNG-14]
	_ = x[ENTSPERRUNG-15]
}

const _Anfragekategorie_name = "PROZESSDATENBERICHTGERAETEUEBERNAHMEWEITERVERPFLICHTUNG_BETRIEB_MELOAENDERUNG_MELOSTAMMDATEN_MALO_ODER_MELOBILANZIERTE_MENGE_MEHR_MINDER_MENGENALLOKATIONSLISTE_MEHR_MINDER_MENGENENERGIEMENGE_UND_LEISTUNGSMAXIMUMABRECHNUNG_MESSSTELLENBETRIEB_MSB_AN_LFAENDERUNG_PROGNOSEGRUNDLAGE_GERAETEKONFIGURATIONAENDERUNG_GERAETEKONFIGURATIONREKLAMATION_VON_WERTENLASTGANG_MALO_TRANCHESPERRUNGENTSPERRUNG"

var _Anfragekategorie_index = [...]uint16{0, 19, 36, 68, 82, 107, 143, 178, 211, 250, 298, 328, 350, 371, 379, 390}

func (i Anfragekategorie) String() string {
	i -= 1
	if i < 0 || i >= Anfragekategorie(len(_Anfragekategorie_index)-1) {
		return "Anfragekategorie(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Anfragekategorie_name[_Anfragekategorie_index[i]:_Anfragekategorie_index[i+1]]
}
