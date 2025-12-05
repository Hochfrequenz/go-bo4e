package schwachlastfaehig

// Schwachlastfaehig Schwachlastfähigkeit Marktlokation
//
//go:generate stringer --type Schwachlastfaehig
//go:generate jsonenums --type Schwachlastfaehig
type Schwachlastfaehig int

const (
	NICHT_SCHWACHLASTFAEHIG Schwachlastfaehig = iota + 1
	SCHWACHLASTFAEHIG
)
