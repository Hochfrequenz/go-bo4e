package lokationstyp

//go:generate stringer --type Lokationstyp
//go:generate jsonenums --type Lokationstyp
type Lokationstyp int

const (
	MaLo Lokationstyp = iota + 1 // Marktlokation / market location
	MeLo                         // Messlokation / metering location
)
