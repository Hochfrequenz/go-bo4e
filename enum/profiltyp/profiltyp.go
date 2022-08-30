package profiltyp

// Profiltyp describes whether the profil is temperaturabhängig or a standardlastprofil
// Note that this enum is not official BO4E standard (yet)!
//
//go:generate stringer --type Profiltyp
//go:generate jsonenums --type Profiltyp
type Profiltyp int

const (
	// SLS_SEP means SLS/SEP
	SLP_SEP Profiltyp = iota + 1
	TLP_TEP           // TLP_TEP means TLS/TEP
)
