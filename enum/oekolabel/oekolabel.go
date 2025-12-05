package oekolabel

// Oekolabel is an enum
//
//go:generate stringer --type Oekolabel
//go:generate jsonenums --type Oekolabel
type Oekolabel int

const (
	GASGREEN_GRUENER_STROM Oekolabel = iota + 1
	GASGREEN
	GRUENER_STROM_GOLD
	GRUENER_STROM_SILBER
	GRUENER_STROM
	GRUENES_GAS
	NATURWATT_STROM
	OK_POWER
	RENEWABLE_PLUS
	WATERGREEN
	WATERGREEN_PLUS
)
