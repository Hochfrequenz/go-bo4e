package netznutzungsvertragsart

// Netznutzungsvertragsart bildet ab, zwischen wem ein Netznutzungsvertrag besteht
// Note that this is not official BO4E standard (yet).
//go:generate stringer --type Netznutzungsvertragsart
//go:generate jsonenums --type Netznutzungsvertragsart
type Netznutzungsvertragsart int

const (
	// KUNDEN_NB beschreibt einen direkten Vertrag zwischen Kunden und NB (EDIFACT Z08)
	KUNDEN_NB      Netznutzungsvertragsart = iota + 1
	LIEFERANTEN_NB                         // LIEFERANTEN_NB beschreibt einen Vertrag zwischen Lieferanten und NB (EDIFACT Z09)
)
