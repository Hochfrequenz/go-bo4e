package statusart

// Statusart represents the Type of Wert (MSCONS SG10 STS 9015)
//go:generate stringer --type Statusart
//go:generate jsonenums --type Statusart
type Statusart int

const (
	// VERTRAG corresponds Vertrag: 6
	VERTRAG Statusart = iota + 1

	// MESSWERTQUALITAET corresponds to "Messwertqualität: 8"
	MESSWERTQUALITAET

	// MESSKLASSIFIZIERUNG corresponds to "Messklassifizierung: 10"
	MESSKLASSIFIZIERUNG

	// PLAUSIBILISIERUNGSHINWEIS corresponds to "Z33"
	PLAUSIBILISIERUNGSHINWEIS

	// ERSATZWERTBILDUNGSVERFAHREN corresponds to "Z32"
	ERSATZWERTBILDUNGSVERFAHREN

	// KORREKTURGRUND corresponds to "Z34"
	KORREKTURGRUND

	// GASQUALITAET corresponds to "Z31"
	GASQUALITAET
)
