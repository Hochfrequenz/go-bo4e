package abgabeart

//go:generate stringer --type AbgabeArt
//go:generate jsonenums --type AbgabeArt
type AbgabeArt int

const (
	AbgabeArtKAS AbgabeArt = iota
	AbgabeArtSA
	AbgabeArtSAS
	AbgabeArtTA
	AbgabeArtTAS
	AbgabeArtTK
	AbgabeArtTKS
	AbgabeArtTS
	AbgabeArtTSS
)