package mengeneinheit

//go:generate stringer --type Mengeneinheit
//go:generate jsonenums --type Mengeneinheit
type Mengeneinheit int

const (
	W          Mengeneinheit = iota // Watt,
	WH                              // Wattstunde
	KW                              // Kilowatt
	KWH                             // Kilowattstunde
	KVARH                           // Kilovarstunde
	MW                              // Megawatt
	MWH                             // Megawattstunde
	Stueck                          // number of pieces / Stückzahl
	KUBIKMETER                      // m^3
	Stunde                          // hour
	Tag                             // day
	Monat                           // month
	Jahr                            // year
	Prozent                         // percent

)
