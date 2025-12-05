package kategorietechnischeressource

// KategorieTechnischeRessource Kategorisierung der technischen Ressource in Hinblick auf §14a EnWG.
//
//go:generate stringer --type KategorieTechnischeRessource
//go:generate jsonenums --type KategorieTechnischeRessource
type KategorieTechnischeRessource int

const (
	FAELLT_UNTER_14A KategorieTechnischeRessource = iota + 1
	FAELLT_NICHT_UNTER_14A
)
