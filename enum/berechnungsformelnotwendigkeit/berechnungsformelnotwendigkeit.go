package berechnungsformelnotwendigkeit

//go:generate stringer --type BerechnungsformelNotwendigkeit
//go:generate jsonenums --type BerechnungsformelNotwendigkeit
type BerechnungsformelNotwendigkeit int

const (
	BERECHNUNGSFORMEL_NOTWENDIG BerechnungsformelNotwendigkeit = iota + 1
	BERECHNUNGSFORMEL_TRIVIAL
	BERECHNUNGSFORMEL_NICHT_NOTWENDIG
)
