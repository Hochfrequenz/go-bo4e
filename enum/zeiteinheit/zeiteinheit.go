package zeiteinheit

//go:generate stringer --type Zeiteinheit
//go:generate jsonenums --type Zeiteinheit
// Zeiteinheit ist eine Auflistung möglicher Einheiten zur Verwendung in zeitbezogenen Angaben.
type Zeiteinheit int

const (
	Sekunde       Zeiteinheit = iota + 1 // second
	Minute                               // minute
	Stunde                               // hour
	ViertelStunde                        // quarter-hour
	Tag                                  // day
	Woche                                // week
	Monat                                // month
	Quartal                              // 1/4 year
	Halbjahr                             // 1/2 half year
	Jahr                                 // year
)
