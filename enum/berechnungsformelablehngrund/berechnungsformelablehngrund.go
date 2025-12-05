package berechnungsformelablehngrund

//go:generate stringer --type BerechnungsformelAblehngrund
//go:generate jsonenums --type BerechnungsformelAblehngrund
type BerechnungsformelAblehngrund int

const (
	START_DATUM_UNPLAUSIBEL BerechnungsformelAblehngrund = iota + 1
	ZU_VIELE_MESSLOKATIONEN
	FEHLENDE_MESSLOKATIONEN
	FALSCHE_MESSLOKATION
)
