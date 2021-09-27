package tarifart

// Tarifart wird verwendet zur Charakterisierung von Zählern und daraus resultierenden Tarifen.
//go:generate stringer --type Tarifart
//go:generate jsonenums --type Tarifart
type Tarifart int

const (
	Eintarif          Tarifart = iota + 1 // Eintarif
	Zweitarif                             // Zweitarif
	Mehrtarif                             // Mehrtarif
	SmartMeter                            // Smart Meter Tarif
	Leistungsgemessen                     // Leistungsgemessener Tarif
)
