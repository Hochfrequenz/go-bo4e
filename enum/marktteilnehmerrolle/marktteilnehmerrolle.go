package marktteilnehmerrolle

// Marktteilnehmerrolle describes the role of a market participant in the context of a message
//
//go:generate stringer --type Marktteilnehmerrolle
//go:generate jsonenums --type Marktteilnehmerrolle
type Marktteilnehmerrolle int

const (
	ANDERE_PARTEI Marktteilnehmerrolle = iota + 1 // ANDERE_PARTEI is another party involved in the process
	EMPFAENGER                                     // EMPFAENGER is the recipient of a message
)
