package prioritaet

//go:generate stringer --type Prioritaet
//go:generate jsonenums --type Prioritaet
type Prioritaet int

const (
	SEHR_NIEDRIG Prioritaet = iota + 1
	NIEDRIG
	NORMAL
	HOCH
	SEHR_HOCH
)
