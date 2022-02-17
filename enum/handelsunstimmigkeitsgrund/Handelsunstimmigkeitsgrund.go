package handelsunstimmigkeitsgrund

// Handelsunstimmigkeitsgrund are the possible types of reasons for the correctness of "Rechnung" or "Lieferscheine"
// (COMDIS SG3 AJT 4465)
//go:generate stringer --type Handelsunstimmigkeitsgrund
//go:generate jsonenums --type Handelsunstimmigkeitsgrund
type Handelsunstimmigkeitsgrund int

const (
	// ANMELDUNG_BESTAETIGT represents COMDIS SG3 AJT 4465 code "Z58"
	ANMELDUNG_BESTAETIGT Handelsunstimmigkeitsgrund = iota + 1

	// ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN represents COMDIS SG3 AJT 4465 code "Z59"
	ABRECHNUNGSBEGINN_GLEICH_BESTAETIGTEM_VERTRAGSBEGINN

	// ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE represents COMDIS SG3 AJT 4465 code "Z60"
	ABRECHNUNGSENDE_GLEICH_BESTAETIGTEM_VERTRAGSENDE

	// NN_MSCONS_UEBERSENDET represents COMDIS SG3 AJT 4465 code "Z61"
	NN_MSCONS_UEBERSENDET

	// RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET represents COMDIS SG3 AJT 4465 code "Z62"
	RICHTIGE_MESSWERTE_ENERGIEMENGEN_UEBERSENDET

	// SONSTIGES_SIEHE_BEGRUENDUNG represents COMDIS SG3 AJT 4465 code "28"
	SONSTIGES_SIEHE_BEGRUENDUNG
)
