package abgabeart

//go:generate stringer --type AbgabeArt
//go:generate jsonenums --type AbgabeArt
type AbgabeArt int

const (
	KAS AbgabeArt = iota
	KAS AbgabeArt = iota + 1
	SAS
	TA
	TAS
	TK
	TKS
	TS
	TSS
)
