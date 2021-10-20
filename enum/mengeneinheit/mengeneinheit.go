package mengeneinheit

// Mengeneinheit is a unit used for measurements
//go:generate stringer --type Mengeneinheit
//go:generate jsonenums --type Mengeneinheit
type Mengeneinheit int

const (
	// W is for Watt
	W          Mengeneinheit = iota + 1
	WH                       // WH is for Watt hour
	KW                       // KW is for Kilo watt
	KWH                      // KWH is for Kilo watt hour
	KVARH                    // KVARH is the abbreviation for Kilovarstunde
	MW                       // MW is for Megawatt
	MWH                      // MWH is for Mega watt hour
	STUECK                   // STUECK is for number of pieces / Stückzahl
	KUBIKMETER               // KUBIKMETER is m^3
	STUNDE                   // STUNDE is 1 hour
	TAG                      // TAG is 1 day
	MONAT                    // MONAT is 1 month
	JAHR                     // JAHR is year
	PROZENT                  // PROZENT is percent
)
