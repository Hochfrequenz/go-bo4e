package tarifart

// Tarifart wird verwendet zur Charakterisierung von Zählern und daraus resultierenden Tarifen.
//go:generate stringer --type Tarifart
//go:generate jsonenums --type Tarifart
type Tarifart int

const (
	// EINTARIF = single tariff
	EINTARIF          Tarifart = iota + 1
	ZWEITARIF                  // ZWEITARIF = two tariffs
	MEHRTARIF                  // MEHRTARIF = multiple (>2) tariffs
	SMARTMETER                 // SMARTMETER means Smart Meter Tariff
	LEISTUNGSGEMESSEN          // LEISTUNGSGEMESSEN means Leistungsgemessener Tarif
)
