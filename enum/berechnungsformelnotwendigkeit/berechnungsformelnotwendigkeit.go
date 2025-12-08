package berechnungsformelnotwendigkeit

//go:generate stringer --type BerechnungsformelNotwendigkeit
//go:generate jsonenums --type BerechnungsformelNotwendigkeit
type BerechnungsformelNotwendigkeit int

const (
	BERECHNUNGSFORMEL_NOTWENDIG             BerechnungsformelNotwendigkeit = iota + 1 // Z33: Berechnungsformel angefuegt
	BERECHNUNGSFORMEL_TRIVIAL                                                         // Z40: Berechnungsformel besitzt keine Rechenoperation
	BERECHNUNGSFORMEL_NICHT_NOTWENDIG                                                 // Berechnungsformel ist nicht erforderlich
	BERECHNUNGSFORMEL_MUSS_ANGEFRAGT_WERDEN                                           // Z34: Berechnungsformel muss beim Absender angefragt werden
)
