package profilart

// Profilart describes whether the profil is temperaturabhängig or a standardlastprofil
//go:generate stringer --type Profilart
//go:generate jsonenums --type Profilart
type Profilart int

const (
	ART_STANDARDLASTPROFIL Profilart = iota + 1	//Z02
	ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL    //Z03
	ART_LASTPROFIL           //Z12
)
