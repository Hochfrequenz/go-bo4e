package gasqualitaet

// Gasqualitaet differentiates between high and low caloric gas
//go:generate stringer --type Gasqualitaet
//go:generate jsonenums --type Gasqualitaet
type Gasqualitaet int

const (
	// H_GAS stands for High Caloric Gas
	H_GAS Gasqualitaet = iota + 1
	L_GAS              // L_GAS stands for Low Caloric Gas
)
