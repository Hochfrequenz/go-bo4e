package plausibilisierungshinweis

// Plausibilisierungshinweis gives hints on plausibility
//go:generate stringer --type Plausibilisierungshinweis
//go:generate jsonenums --type  Plausibilisierungshinweis
type Plausibilisierungshinweis int

const (
	// Z83_KUNDENSELBSTABLESUNG ist Kundenselbstablesung
	Z83_KUNDENSELBSTABLESUNG Plausibilisierungshinweis = iota + 1

	// Z84_LEERSTAND ist Leerstand
	Z84_LEERSTAND

	// Z85_REALERZAEHLERUEBERLAUFGEPRUEFT ist Realer Zaehlerüberlauf geprueft
	Z85_REALERZAEHLERUEBERLAUFGEPRUEFT

	// Z86_PLAUSIBELWGKONTROLLABLESUNG ist Plausibel wg. Kontrollablesung
	Z86_PLAUSIBELWGKONTROLLABLESUNG

	// Z87_PLAUSIBELWGKUNDENHINWEIS ist Plausibel wg. Kundenhinweis
	Z87_PLAUSIBELWGKUNDENHINWEIS

	// ZC3_AUSTAUSCHDESERSATZWERTES ist Austausch des Ersatzwertes
	ZC3_AUSTAUSCHDESERSATZWERTES

	// ZR5_RECHENWERT is "Rechenwert"
	ZR5_RECHENWERT
)
