package oekozertifikat

// Oekozertifikat is an enum
//
//go:generate stringer --type Oekozertifikat
//go:generate jsonenums --type Oekozertifikat
type Oekozertifikat int

const (
	CMS_EE02 Oekozertifikat = iota + 1
	EECS
	FRAUNHOFER
	BET
	KLIMA_INVEST
	LGA
	FREIBERG
	RECS
	REGS_EGL
	TUEV
	TUEV_HESSEN
	TUEV_NORD
	TUEV_RHEINLAND
	TUEV_SUED
	TUEV_SUED_EE01
	TUEV_SUED_EE02
)
