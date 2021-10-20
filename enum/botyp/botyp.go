package botyp

// BOTyp is an enumeration that contains all available Business Objects
//go:generate stringer --type BOTyp
//go:generate jsonenums --type BOTyp
type BOTyp int

const (
	// ANGEBOT is an offer
	ANGEBOT                     BOTyp = iota + 1
	ANSPRECHPARTNER                   // ANSPRECHPARTNER is the type of a contact person, see bo.ANSPRECHPARTNER
	AUSSCHREIBUNG                     // AUSSCHREIBUNG is a tender offer
	ENERGIEMENGE                      // ENERGIEMENGE is the tpye of an energy amount, see bo.ENERGIEMENGE
	GESCHAEFTSPARTNER                 // GESCHAEFTSPARTNER is a business partner
	KOSTEN                            // KOSTEN are costs
	LASTGANG                          // LASTGANG is the type of bo.LASTGANG
	MARKTLOKATION                     // MARKTLOKATION is a market location / "MaLo"
	MESSLOKATION                      // MESSLOKATION is a metering location / "MeLo"
	MARKTTEILNEHMER                   // MARKTTEILNEHMER is a market partner
	NETZNUTZUNGSRECHNUNG              // NETZNUTZUNGSRECHNUNG is a grid usage bill
	PREISBLATT                        // PREISBLATT is a price sheet
	PREISBLATTDIENSTLEISTUNG          // PREISBLATTDIENSTLEISTUNG is a PREISBLATT for services
	PREISBLATTKONZESSIONSABGABE       // PREISBLATTKONZESSIONSABGABE is a PREISBLATT for concession fees
	PREISBLATTMESSUNG                 // PREISBLATTMESSUNG is a PREISBLATT for metering
	PREISBLATTUMLAGEN                 // PREISBLATTUMLAGEN is a PREISBLATT for allocations
	RECHNUNG                          // RECHNUNG is a Invoice
	TARIFINFO                         // TARIFINFO Are information about tariffs
	TARIFPREISBLATT                   // TARIFPREISBLATT is a PREISBLATT about Tariffs
	VERTRAG                           // VERTRAG is contract
	ZAEHLER                           // ZAEHLER is a power, gas, heat or water meter
	ZEITREIHE                         // ZEITREIHE is data over time (e.g. consumption, prognosis etc.)
)
