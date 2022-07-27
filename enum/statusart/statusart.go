package statusart

//go:generate stringer --type StatusArt
//go:generate jsonenums --type StatusArt

// StatusArt represents the Type of Wert (MSCONS SG10 STS 9015).
type StatusArt int

const (
	// VERTRAG represents MSCONS SG10 STS 9015 value "6"
	VERTRAG StatusArt = iota + 1

	// MESSWERTQUALITAET represents MSCONS SG10 STS 9015 value "8"
	MESSWERTQUALITAET

	// MESSKLASSIFIZIERUNG represents MSCONS SG10 STS 9015 value "10"
	MESSKLASSIFIZIERUNG

	// PLAUSIBILISIERUNGSHINWEIS represents MSCONS SG10 STS 9015 value "Z33"
	PLAUSIBILISIERUNGSHINWEIS

	// ERSATZWERTBILDUNGSVERFAHREN represents MSCONS SG10 STS 9015 value "Z32"
	ERSATZWERTBILDUNGSVERFAHREN

	// KORREKTURGRUND represents MSCONS SG10 STS 9015 value "Z34"
	KORREKTURGRUND

	// GASQUALITAET represents MSCONS SG10 STS 9015 value "Z31"
	GASQUALITAET
)
