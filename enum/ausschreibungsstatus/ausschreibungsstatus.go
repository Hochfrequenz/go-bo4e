package ausschreibungsstatus

// Ausschreibungsstatus is an enum
//
//go:generate stringer --type Ausschreibungsstatus
//go:generate jsonenums --type Ausschreibungsstatus
type Ausschreibungsstatus int

const (
	PHASE1 Ausschreibungsstatus = iota + 1
	PHASE2
	PHASE3
	PHASE4
)
