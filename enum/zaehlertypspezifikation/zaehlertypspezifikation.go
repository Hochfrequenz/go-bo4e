package zaehlertypspezifikation

// ZaehlertypSpezifikation is an enum
//
//go:generate stringer --type ZaehlertypSpezifikation
//go:generate jsonenums --type ZaehlertypSpezifikation
type ZaehlertypSpezifikation int

const (
	EDL40 ZaehlertypSpezifikation = iota + 1
	EDL21
	SONSTIGER_EHZ
	MME_STANDARD
	MME_MEDA
)
