package waehrungseinheit

// why tf is this a separate enum and not part of the waehrungscode.Waehrungscode?

// Waehrungseinheit models euro and euro cent. In diesem Enum werden die Währungen und ihre Untereinheiten definiert, beispielsweise für die Verwendung in Preisen.
//go:generate stringer --type Waehrungseinheit
//go:generate jsonenums --type Waehrungseinheit
type Waehrungseinheit int

const (
	// EUR is for Euro
	EUR Waehrungseinheit = iota + 1
	// CT is for Eurocent
	CT
)
