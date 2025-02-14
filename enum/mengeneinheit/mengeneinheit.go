package mengeneinheit

// Mengeneinheit is a unit used for measurements
//
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
	ANZAHL                   // ANZAHL means the same as STUECK (it has been renamed in the official docs, so this is just a fallback, use STUECK if you can); see also https://github.com/Hochfrequenz/BO4E-python/issues/359
	KUBIKMETER               // KUBIKMETER is m^3
	STUNDE                   // STUNDE is 1 hour
	TAG                      // TAG is 1 day
	MONAT                    // MONAT is 1 month
	JAHR                     // JAHR is year
	PROZENT                  // PROZENT is percent
	ZERO                     // ZERO should not be used and is just here for full-compatability with the bo4e.net lib
)
