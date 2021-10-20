package lokationstyp

// The Lokationstyp describe whether the object it's used in is a market or metering location
//go:generate stringer --type Lokationstyp
//go:generate jsonenums --type Lokationstyp
type Lokationstyp int

const (
	MALO Lokationstyp = iota + 1 // MALO is short for Marktlokation / market location
	MELO                         // MELO is short for Messlokation / metering location
)
