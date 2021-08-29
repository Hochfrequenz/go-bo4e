package lokationstyp

//go:generate stringer --type Lokationstyp
//go:generate jsonenums --type Lokationstyp
// The Lokationstyp describe whether the object it's used in is a market or metering location
type Lokationstyp int

const (
	MaLo Lokationstyp = iota + 1 // Marktlokation / market location
	MeLo                         // Messlokation / metering location
)
