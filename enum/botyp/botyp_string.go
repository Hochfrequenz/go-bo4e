// Code generated by "stringer --type BOTyp"; DO NOT EDIT.

package botyp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ANGEBOT-1]
	_ = x[ANSPRECHPARTNER-2]
	_ = x[AUSSCHREIBUNG-3]
	_ = x[BILANZIERUNG-4]
	_ = x[ENERGIEMENGE-5]
	_ = x[GESCHAEFTSPARTNER-6]
	_ = x[KOSTEN-7]
	_ = x[LASTGANG-8]
	_ = x[MARKTLOKATION-9]
	_ = x[MESSLOKATION-10]
	_ = x[MARKTTEILNEHMER-11]
	_ = x[NETZNUTZUNGSRECHNUNG-12]
	_ = x[PREISBLATT-13]
	_ = x[PREISBLATTDIENSTLEISTUNG-14]
	_ = x[PREISBLATTKONZESSIONSABGABE-15]
	_ = x[PREISBLATTMESSUNG-16]
	_ = x[PREISBLATTUMLAGEN-17]
	_ = x[RECHNUNG-18]
	_ = x[REKLAMATION-19]
	_ = x[TARIFINFO-20]
	_ = x[TARIFPREISBLATT-21]
	_ = x[VERTRAG-22]
	_ = x[ZAEHLER-23]
	_ = x[ZEITREIHE-24]
	_ = x[HANDELSUNSTIMMIGKEIT-25]
	_ = x[AVIS-26]
}

const _BOTyp_name = "ANGEBOTANSPRECHPARTNERAUSSCHREIBUNGBILANZIERUNGENERGIEMENGEGESCHAEFTSPARTNERKOSTENLASTGANGMARKTLOKATIONMESSLOKATIONMARKTTEILNEHMERNETZNUTZUNGSRECHNUNGPREISBLATTPREISBLATTDIENSTLEISTUNGPREISBLATTKONZESSIONSABGABEPREISBLATTMESSUNGPREISBLATTUMLAGENRECHNUNGREKLAMATIONTARIFINFOTARIFPREISBLATTVERTRAGZAEHLERZEITREIHEHANDELSUNSTIMMIGKEITAVIS"

var _BOTyp_index = [...]uint16{0, 7, 22, 35, 47, 59, 76, 82, 90, 103, 115, 130, 150, 160, 184, 211, 228, 245, 253, 264, 273, 288, 295, 302, 311, 331, 335}

func (i BOTyp) String() string {
	i -= 1
	if i < 0 || i >= BOTyp(len(_BOTyp_index)-1) {
		return "BOTyp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _BOTyp_name[_BOTyp_index[i]:_BOTyp_index[i+1]]
}
