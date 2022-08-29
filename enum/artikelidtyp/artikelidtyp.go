package artikelidtyp

// ArtikelIdTyp ermöglicht die Unterscheidung zwischen ArtikelId und GruppenArtikelId
//go:generate stringer --type ArtikelIdTyp
//go:generate jsonenums --type ArtikelIdTyp
type ArtikelIdTyp int

const (
	ARTIKELID                                                                   ArtikelIdTyp = iota + 1 // Übertragungsnetzbetreiber
	GRUPPENARTIKELID // Verteilnetzbetreiber
)
