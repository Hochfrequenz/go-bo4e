package sparte

//go:generate stringer --type Sparte

type Sparte int

const (
	Strom      Sparte = iota // power/electricity
	Gas                      // gas
	Fernwaerme               // long-distance heat
	Nahwaerme                // local heat
	Wasser                   // water
	Abwasser                 // waste water
)
