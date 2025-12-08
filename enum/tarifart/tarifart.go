package tarifart

// Tarifart wird verwendet zur Charakterisierung von Zählern und daraus resultierenden Tarifen.
//
//go:generate stringer --type Tarifart
//go:generate jsonenums --type Tarifart
type Tarifart int

const (
	// EINTARIF = single tariff
	EINTARIF          Tarifart = iota + 1
	ZWEITARIF                  // ZWEITARIF = two tariffs
	MEHRTARIF                  // MEHRTARIF = multiple (>2) tariffs
	SMARTMETER                 // Deprecated: SMARTMETER means Smart Meter Tariff (use SMART_METER instead)
	LEISTUNGSGEMESSEN          // LEISTUNGSGEMESSEN means Leistungsgemessener Tarif
	SMART_METER                // SMART_METER means Smart Meter Tariff
)
