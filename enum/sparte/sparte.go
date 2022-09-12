package sparte

// Sparte contains different divisions of typical utilities.
//
//go:generate stringer --type Sparte
//go:generate jsonenums --type Sparte
type Sparte int

const (
	// STROM means power/electricity
	STROM      Sparte = iota + 1
	GAS               // GAS means gas
	FERNWAERME        // FERNWAERME means long-distance heat
	NAHWAERME         // NAHWAERME means local heat
	WASSER            // WASSER means water
	ABWASSER          // ABWASSER means waste water
)
