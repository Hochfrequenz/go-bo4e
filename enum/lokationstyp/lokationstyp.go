package lokationstyp

//go:generate stringer --type Lokationstyp
//go:generate jsonenums --type Lokationstyp
type Lokationstyp int

const (
	MaLo Lokationstyp = iota // Marktlokation / market location
	MeLo                     // Messlokation / metering location
)
