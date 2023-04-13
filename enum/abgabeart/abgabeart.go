package abgabeart

//go:generate stringer --type AbgabeArt
//go:generate jsonenums --type AbgabeArt
type AbgabeArt int

const (
	KAS AbgabeArt = iota
	SA
	SAS
	TA
	TAS
	TK
	TKS
	TS
	TSS
)
