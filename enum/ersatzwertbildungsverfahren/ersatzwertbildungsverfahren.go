package ersatzwertbildungsverfahren

// Ersatzwertbildungsverfahren describes the method used to generate replacement values
//go:generate stringer --type Ersatzwertbildungsverfahren
//go:generate jsonenums --type Ersatzwertbildungsverfahren
type Ersatzwertbildungsverfahren int

const (
	// VERGLEICHSMESSUNG_GEEICHT corresponds to edi@energy qualifier Z88 "Vergleichsmessung (geeicht)"
	VERGLEICHSMESSUNG_GEEICHT Ersatzwertbildungsverfahren = iota + 1

	// VERGLEICHSMESSUNG_NICHT_GEEICHT corresponds to edi@energy qualifier Z89 "Vergleichmessung (nicht geeicht)
	VERGLEICHSMESSUNG_NICHT_GEEICHT

	// MESSWERTNACHBILDUNG_AUS_GEEICHTE_NWERTEN corresponds to edi@energy qualifier
	// Z90 "Messwertnachbildung aus geeichten Werten"
	MESSWERTNACHBILDUNG_AUS_GEEICHTE_NWERTEN

	// MESSWERTNACHBILDUNG_AUS_NICHT_GEEICHTEN_WERTEN corresponds to edi@energy qualifier
	// Z91 "Messwertnachbildung aus nicht geeichten Werten"
	MESSWERTNACHBILDUNG_AUS_NICHT_GEEICHTEN_WERTEN

	// INTERPOLATION corresponds to edi@energy qualifier Z92 "Interpolation"
	INTERPOLATION

	// HALTEWERT corresponds to edi@energy qualifier Z93 "Haltewert"
	HALTEWERT

	// BILANZIERUNGNETZABSCHNITT corresponds to edi@energy qualifier Z94 "Bilanzierungsabschnitt"
	BILANZIERUNGNETZABSCHNITT

	// HISTORISCHE_MESSWERTE corresponds to edi@energy qualifier Z95 "Historische Messwerte"
	HISTORISCHE_MESSWERTE

	// AUFTEILUNG corresponds to edi@energy qualifier ZQ8 "Aufteilung"
	AUFTEILUNG

	// VERWENDUNG_VON_WERTEN_DES_STOERMENGENZAEHLWERKS corresponds to edi@energy qualifier
	// ZQ9 "Verwendung von Werten des Störmengenzählwerks"
	VERWENDUNG_VON_WERTEN_DES_STOERMENGENZAEHLWERKS

	// UMGANGS_UND_KORREKTURMENGEN corresponds to edi@energy qualifier ZR0 "Umgangs- und Korrekturmengen"
	UMGANGS_UND_KORREKTURMENGEN

	// STATISCHE_METHODE corresponds to edi@energy qualifier ZJ2 "Statische Methode"
	STATISCHE_METHODE
)
