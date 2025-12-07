package verbrauchsmengetyp

// Verbrauchsmengetyp type von Verbrauchsmenge
//
//go:generate stringer --type Verbrauchsmengetyp
//go:generate jsonenums --type Verbrauchsmengetyp
type Verbrauchsmengetyp int

const (
	ARBEITLEISTUNGTAGESPARAMETERABHMALO Verbrauchsmengetyp = iota + 1
	VERANSCHLAGTEJAHRESMENGE
	TUMKUNDENWERT
)
