package botyp

//go:generate stringer --type BOTyp
//go:generate jsonenums --type BOTyp

// BOTyp is an enumeration that contains all available Business Objects
type BOTyp int

const (
	Angebot           BOTyp = iota + 1 // offer
	Ansprechpartner                    // contact person
	Ausschreibung                      // tender offer
	Energiemenge                       // energy amount
	Geschaeftspartner                  // business partner
	Kosten                             // costs
	Lastgang
	Marktlokation        // market location / "MaLo"
	Messlokation         // metering location / "MeLo"
	Marktteilnehmer      // market partner
	Netznutzungsrechnung //
	Preisblatt           // price sheet
	PreisblattDienstleistung
	PreisblattKonzessionsabgabe
	PreisblattMessung
	PreisblattUmlagen
	Rechnung // Invoice
	Tarifinfo
	TarifPreisblatt
	Vertrag // contract
	Zaehler // meter
	Zeitreihe
)
