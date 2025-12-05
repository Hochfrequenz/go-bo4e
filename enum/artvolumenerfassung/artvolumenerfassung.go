package artvolumenerfassung

// ArtVolumenerfassung Auflistung möglicher Arten der Volumenerfassung (bzgl. Mengenumwerter)
//
//go:generate stringer --type ArtVolumenerfassung
//go:generate jsonenums --type ArtVolumenerfassung
type ArtVolumenerfassung int

const (
	KENNLINIENKORREKTUR ArtVolumenerfassung = iota + 1
	SCHLEICHMENGENUNTERDRÜCKUNG
)
