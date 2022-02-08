package plausibilisierungshinweis

// Plausibilisierungshinweis gives hints on plausibility
//go:generate stringer --type Plausibilisierungshinweis
//go:generate jsonenums --type  Plausibilisierungshinweis
type Plausibilisierungshinweis int

const (
	// KUNDENSELBSTABLESUNG corresponds to edi@energy qualifier Z83 "Kundenselbstablesung"
	KUNDENSELBSTABLESUNG Plausibilisierungshinweis = iota + 1

	// LEERSTAND corresponds to edi@energy qualifier Z84 "Leerstand"
	LEERSTAND

	// REALER_ZAEHLERUEBERLAUF_GEPRUEFT corresponds to edi@energy qualifier Z85 "Realer Zaehlerüberlauf geprueft"
	REALER_ZAEHLERUEBERLAUF_GEPRUEFT

	// PLAUSIBEL_WG_KONTROLLABLESUNG corresponds to edi@energy qualifier Z86 "Plausibel wg. Kontrollablesung"
	PLAUSIBEL_WG_KONTROLLABLESUNG

	// PLAUSIBEL_WG_KUNDENHINWEIS corresponds to edi@energy qualifier Z87 "Plausibel wg. Kundenhinweis"
	PLAUSIBEL_WG_KUNDENHINWEIS

	// AUSTAUSCH_DES_ERSATZWERTES corresponds to edi@energy qualifier ZC3 "Austausch des Ersatzwertes"
	AUSTAUSCH_DES_ERSATZWERTES

	// RECHENWERT corresponds to edi@energy qualifier ZR5 "Rechenwert"
	RECHENWERT
)
