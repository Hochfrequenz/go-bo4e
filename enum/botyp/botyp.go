package botyp

// BOTyp is an enumeration that contains all available Business Objects
//go:generate stringer --type BOTyp
//go:generate jsonenums --type BOTyp
type BOTyp int

const (
	ANGEBOT              BOTyp = iota + 1 // offer
	ANSPRECHPARTNER                       // ANSPRECHPARTNER is the type of a contact person, see bo.ANSPRECHPARTNER
	AUSSCHREIBUNG                         // tender offer
	ENERGIEMENGE                          // ENERGIEMENGE is the tpye of an energy amount, see bo.ENERGIEMENGE
	GESCHAEFTSPARTNER                     // business partner
	KOSTEN                                // costs
	LASTGANG                              // LASTGANG is the type of bo.LASTGANG
	MARKTLOKATION                         // market location / "MaLo"
	MESSLOKATION                          // metering location / "MeLo"
	MARKTTEILNEHMER                       // market partner
	NETZNUTZUNGSRECHNUNG                  //
	PREISBLATT                            // price sheet
	PREISBLATTDIENSTLEISTUNG
	PREISBLATTKONZESSIONSABGABE
	PREISBLATTMESSUNG
	PREISBLATTUMLAGEN
	RECHNUNG // Invoice
	TARIFINFO
	TARIFPREISBLATT
	VERTRAG // contract
	ZAEHLER // meter
	ZEITREIHE
)
