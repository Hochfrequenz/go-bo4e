package unterbrechbarkeit

// Stromverbrauchsart/Unterbrechbarkeit Marktlokation

//go:generate stringer --type Unterbrechbarkeit
//go:generate jsonenums --type Unterbrechbarkeit
type Unterbrechbarkeit int

const (
	UV Unterbrechbarkeit = iota + 1
	NUV
)
