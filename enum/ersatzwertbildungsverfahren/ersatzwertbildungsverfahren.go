package ersatzwertbildungsverfahren

// Ersatzwertbildungsverfahren describes the method used to generate replacement values
//go:generate stringer --type Ersatzwertbildungsverfahren
//go:generate jsonenums --type Ersatzwertbildungsverfahren
type Ersatzwertbildungsverfahren int

const (
	// Z88_VERGLEICHSMESSUNGGEEICHT is "Vergleichsmessung (geeicht)"
	Z88_VERGLEICHSMESSUNGGEEICHT Ersatzwertbildungsverfahren = iota + 1

	// Z89_VERGLEICHSMESSUNGNICHTGEEICHT is "Vergleichmessung (nicht geeicht)
	Z89_VERGLEICHSMESSUNGNICHTGEEICHT

	// Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN is
	// "Messwertnachbildung aus geeichten Werten"
	Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN

	// Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN is
	// "Messwertnachbildung aus nicht geeichten Werten"
	Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN

	// Z92_INTERPOLATION is "Interpolation"
	Z92_INTERPOLATION

	// Z93_HALTEWERT is "Haltewert"
	Z93_HALTEWERT

	// Z94_BILANZIERUNGNETZABSCHNITT is "Bilanzierungsabschnitt"
	Z94_BILANZIERUNGNETZABSCHNITT

	// Z95_HISTORISCHEMESSWERTE is "Historische Messwerte"
	Z95_HISTORISCHEMESSWERTE

	// ZQ8_AUFTEILUNG is "Aufteilung"
	ZQ8_AUFTEILUNG

	// ZQ9_VERWENDUNGVONWERTENDESSTOERMENGENZAEHLWERKS is "Verwendung von Werten des Störmengenzählwerks"
	ZQ9_VERWENDUNGVONWERTENDESSTOERMENGENZAEHLWERKS

	// ZR0_UMGANGSUNDKORREKTURMENGEN is "Umgangs- und Korrekturmengen"
	ZR0_UMGANGSUNDKORREKTURMENGEN

	// ZJ2_STATISCHEMETHODE is "Statische Methode"
	ZJ2_STATISCHEMETHODE
)
