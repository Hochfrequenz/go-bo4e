package grundlageenergiemenge

// GrundlageEnergiemenge gives hints about the foundation of an energy amount
//go:generate stringer --type GrundlageEnergiemenge
//go:generate jsonenums --type  GrundlageEnergiemenge
type GrundlageEnergiemenge int

const (
	// Z36_ZAEHLERSTANDZUMBEGINDERANGEGEBENENENERGIEMENGEVORHANDENUNDKOMMUNIZIERT
	// is "Zählerstand zum Beginn der angegebenen Energiemenge vorhanden und kommuniziert"
	Z36_ZAEHLERSTANDZUMBEGINDERANGEGEBENENENERGIEMENGEVORHANDENUNDKOMMUNIZIERT GrundlageEnergiemenge = iota + 1

	// Z37_ZAEHLERSTANDZUMENDEDERANGEGEBENENENERGIEMENGEVORHANDENUNDKOMMUNIZIERT
	// is "Zählerstand zum Ende der angegebenen Energiemenge vorhanden und kommuniziert
	Z37_ZAEHLERSTANDZUMENDEDERANGEGEBENENENERGIEMENGEVORHANDENUNDKOMMUNIZIERT

	// Z37_ZAEHLERSTANDZUMENDEDERANGEGEBENENENERGIEMENGEVORHANDENUNDKOMMUNIZIERT
	// is "Zählerstand zum Beginn der angegebenen Energiemenge nicht vorhanden da Mengenabgrenzung"
	Z38_ZAEHLERSTANDZUMBEGINDERANGEGEBENENENERGIEMENGENICHTVORHANDENDAMENGENABGRENZUNG

	// Z39_ZAEHLERSTANDZUMENDEDERANGEGEBENENENERGIEMENGENICHTVORHANDENDAMENGENABGRENZUNG
	// is "Zählerstand zum Ende der angegebenen Energiemenge nicht vorhanden da Mengenabgrenzung"
	Z39_ZAEHLERSTANDZUMENDEDERANGEGEBENENENERGIEMENGENICHTVORHANDENDAMENGENABGRENZUNG
)
