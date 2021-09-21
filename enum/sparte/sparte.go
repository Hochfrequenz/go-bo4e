package sparte

// Sparte contains different divisions of typical utilities.
//go:generate stringer --type Sparte
//go:generate jsonenums --type Sparte
type Sparte int

const (
	Strom      Sparte = iota + 1 // power/electricity
	Gas                          // gas
	Fernwaerme                   // long-distance heat
	Nahwaerme                    // local heat
	Wasser                       // water
	Abwasser                     // waste water
)
