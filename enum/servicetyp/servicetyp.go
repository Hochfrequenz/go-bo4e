package servicetyp

// Servicetyp Typ eines Services (Hochfrequenz-eigene Erweiterung des BO4E-Standards.
//
//go:generate stringer --type Servicetyp
//go:generate jsonenums --type Servicetyp
type Servicetyp int

const (
	STROM_NB Servicetyp = iota + 1
	STROM_MSB
	STROM_LIEF
	GAS_NB
	GAS_MSB
	GAS_LIEF
)
