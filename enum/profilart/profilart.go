package profilart

// Profilart describes whether the profile is temperaturabhängig or a Standardlastprofil (SLP)
//
//go:generate stringer --type Profilart
//go:generate jsonenums --type Profilart
type Profilart int

// todo: how and why does this differ from profiltyp.Profiltyp?

const (
	ART_STANDARDLASTPROFIL                        Profilart = iota + 1 //ART_STANDARDLASTPROFIL is SLP (Z02 in EDIFACT)
	ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL                           //ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL is TLP (Z03 in EDIFACT)
	ART_LASTPROFIL                                                     // ART_LASTPROFIL is ??? (Z12 in EDIFACT)
	ART_STANDARDEINSPEISEPROFIL                                        // ART_STANDARDEINSPEISEPROFIL is a Standardeinspeiseprofil (Z04 in EDIFACT)
	ART_TAGESPARAMETERABHAENGIGES_EINSPEISEPROFIL                      // ART_TAGESPARAMETERABHAENGIGES_EINSPEISEPROFIL is a tagesparameterabhängiges Einspeiseprofil (Z05 in EDIFACT)
)
