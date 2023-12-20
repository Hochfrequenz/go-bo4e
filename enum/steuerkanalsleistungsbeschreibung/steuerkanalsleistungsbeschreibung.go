package steuerkanalsleistungsbeschreibung

//go:generate stringer --type Steuerkanalsleistungsbeschreibung
//go:generate jsonenums --type Steuerkanalsleistungsbeschreibung
type Steuerkanalsleistungsbeschreibung int

const (
	AN_AUS  Steuerkanalsleistungsbeschreibung = iota + 1 // AN/AUS
	GESTUFT                                              // GESTUFT
)
