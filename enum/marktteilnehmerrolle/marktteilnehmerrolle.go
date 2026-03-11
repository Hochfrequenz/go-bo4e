package marktteilnehmerrolle

// Marktteilnehmerrolle: Rolle eines Marktteilnehmers im Nachrichtenkontext.
//
//go:generate stringer --type Marktteilnehmerrolle
//go:generate jsonenums --type Marktteilnehmerrolle
type Marktteilnehmerrolle int

const (
	// ANDERE_PARTEI: Andere Partei - ein am Vorgang beteiligter Marktpartner.
	ANDERE_PARTEI Marktteilnehmerrolle = iota + 1
	// EMPFAENGER: Empfaenger einer Nachricht
	EMPFAENGER
)
