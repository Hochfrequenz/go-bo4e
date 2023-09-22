package abwicklungsmodell

// Abwicklungsmodell (E-Mob)
//
//go:generate stringer --type Abwicklungsmodell
//go:generate jsonenums --type Abwicklungsmodell
type Abwicklungsmodell int

const (
	// Modell 1 "Bilanzierung an der Marktlokation"
	MODELL_1 Abwicklungsmodell = iota + 1
	MODELL_2                   // Modell 2 "Bilanzierung im Bilanzierungsgebiet (BG) des LPB"
)
