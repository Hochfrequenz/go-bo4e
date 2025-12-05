package schwachlastfaehig

// Schwachlastfaehig is an enum
//
//go:generate stringer --type Schwachlastfaehig
//go:generate jsonenums --type Schwachlastfaehig
type Schwachlastfaehig int

const (
	NICHT_SCHWACHLASTFAEHIG Schwachlastfaehig = iota + 1
	SCHWACHLASTFAEHIG
)
