package verbrauchsmengetyp

// Verbrauchsmengetyp is an enum
//
//go:generate stringer --type Verbrauchsmengetyp
//go:generate jsonenums --type Verbrauchsmengetyp
type Verbrauchsmengetyp int

const (
	ARBEITLEISTUNGTAGESPARAMETERABHMALO Verbrauchsmengetyp = iota + 1
	VERANSCHLAGTEJAHRESMENGE
	TUMKUNDENWERT
)
