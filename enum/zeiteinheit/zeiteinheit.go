package zeiteinheit

// Zeiteinheit ist eine Auflistung möglicher Einheiten zur Verwendung in zeitbezogenen Angaben.
//
//go:generate stringer --type Zeiteinheit
//go:generate jsonenums --type Zeiteinheit
type Zeiteinheit int

const (
	// SEKUNDE is 1 second
	SEKUNDE        Zeiteinheit = iota + 1
	MINUTE                     // MINUTE is 1 minute
	STUNDE                     // STUNDE is 1 hour
	VIERTELSTUNDE              // Deprecated: VIERTELSTUNDE is quarter-hour (15 MINUTE) but the C# lib serializes as "VIERTEL_STUNDE"
	TAG                        // TAG is 1 day
	WOCHE                      // WOCHE is 1 week
	MONAT                      // MONAT is 1 month
	QUARTAL                    // QUARTAL is 1/4 year
	HALBJAHR                   // HALBJAHR is 1/2 half year
	JAHR                       // JAHR is 1 year
	VIERTEL_STUNDE             // VIERTEL_STUNDE is quarter-hour
)
