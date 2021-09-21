package lokationstyp

// The Lokationstyp describe whether the object it's used in is a market or metering location
//go:generate stringer --type Lokationstyp
//go:generate jsonenums --type Lokationstyp
type Lokationstyp int

const (
	MaLo Lokationstyp = iota + 1 // Marktlokation / market location
	MeLo                         // Messlokation / metering location
)
