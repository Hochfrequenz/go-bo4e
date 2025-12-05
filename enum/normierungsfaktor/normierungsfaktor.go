package normierungsfaktor

// Normierungsfaktor is an enum
//
//go:generate stringer --type Normierungsfaktor
//go:generate jsonenums --type Normierungsfaktor
type Normierungsfaktor int

const (
	NORMIERUNGSFAKTOR_1_000_000_KWH_A Normierungsfaktor = iota + 1
	NORMIERUNGSFAKTOR_300_KWH_K
	NORMIERUNGSFAKTOR_1_000_000_KW
)
