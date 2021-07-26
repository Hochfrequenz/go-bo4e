package botyp

//go:generate stringer --type BOTyp
//go:generate jsonenums --type BOTyp

type BOTyp int

const (
	Angebot              BOTyp = iota // offer
	Ansprechpartner                   // contact person
	Ausschreibung                     // tender offer
	Energiemenge                      // energy amount
	Geschaeftspartner                 // business partner
	Kosten                            // costs
	Marktlokation                     // market location / "MaLo"
	Messlokation                      // metering location / "MeLo"
	Marktteilnehmer                   // market partner
	Netznutzungsrechnung              //
	Preisblatt                        // price sheet
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
