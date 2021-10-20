package marktrolle

// The Marktrolle describes the roles of a bo.Marktteilnehmer
//go:generate stringer --type Marktrolle
//go:generate jsonenums --type Marktrolle
type Marktrolle int

const (
	// NB is a Netzbetreiber
	NB              Marktrolle = iota + 1
	LF                         // LF is a Lieferant
	MSB                        // MSB is a Messstellenbetreiber
	DL                         // DL is a Dienstleister
	BKV                        // BKV is a Bilanzkreisverantwortlicher
	BKO                        // BKO is a Bilanzkoordinator/Marktgebietsverantwortlicher
	UENB                       // UENB is a Übertragungsnetzbetreiber
	KUNDE_SELBST_NN            // KUNDE_SELBST_NN is a Kunden die NN-Entgelte selbst zahlen
	MGV                        // MGV is a Marktgebietsverantwortlicher
	EIV                        // EIV is a Einsatzverantwortlicher
	RB                         // RB is a Registerbetreiber Beispielsweise Herkunftsnachweisregister
	KUNDE                      // KUNDE is a Kunde
	INTERESSENT                // INTERESSENT is a Interessent
)
