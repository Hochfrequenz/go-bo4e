package lokationstyp

//go:generate stringer --type Wertermittlungsverfahren
//go:generate jsonenums --type Wertermittlungsverfahren
type Wertermittlungsverfahren int

const (
	PROGNOSE Wertermittlungsverfahren = iota // Prognose / prognosis
	MESSUNG                                  // Messung / measurement
)
