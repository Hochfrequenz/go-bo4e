package abweichungsgrund

// Gibt einen Abweichungsgrund bei Ablehung einer COMDIS an. (REMADV SG7 AJT 4465)
//go:generate stringer --type AbweichungsGrund
//go:generate jsonenums --type AbweichungsGrund
type AbweichungsGrund int

const (
	PREIS_RECHENREGEL_FALSCH                                                                   AbweichungsGrund = iota + 1 // 5
	FALSCHER_ABRECHNUNGSZEITRAUM                                                                                           // 9
	UNBEKANNTE_MARKTLOKATION_MESSLOKATION                                                                                  // 14
	SONSTIGER_ABWEICHUNGSGRUND                                                                                             // 28
	DOPPELTE_RECHNUNG                                                                                                      // 53
	ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN                                                                              // Z01
	ABRECHNUNGSENDE_UNGLEICH_VERTRAGSENDE                                                                                  // Z02
	BETRAG_DER_ABSCHLAGSRECHNUNG_FALSCH                                                                                    // Z03
	VORAUSBEZAHLTER_BETRAG_FALSCH                                                                                          // Z04
	ARTIKEL_NICHT_VEREINBART                                                                                               // Z06
	NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FEHLEN                                                                             // Z07
	RECHNUNGSNUMMER_BEREITS_ERHALTEN                                                                                       // Z08
	NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FALSCH                                                                             // Z10
	ZEITLICHE_MENGENANGABE_FEHLERHAFT                                                                                      // Z33
	FALSCHER_BILANZIERUNGSBEGINN                                                                                           // Z35
	FALSCHES_NETZNUTZUNGSENDE                                                                                              // Z36
	BILANZIERTE_MENGE_FEHLT                                                                                                // Z37
	BILANZIERTE_MENGE_FALSCH                                                                                               // Z38
	NETZNUTZUNGSABRECHNUNG_FEHLT                                                                                           // Z39
	REVERSE_CHARGE_ANWENDUNG_FEHLT_ODER_FEHLERHAFT                                                                         // Z40
	ALLOKATIONSLISTE_FEHLT                                                                                                 // Z41
	MEHR_MINDERMENGE_FALSCH                                                                                                // Z42
	UNGUELTIGES_RECHNUNGSDATUM                                                                                             // Z43
	ZEITINTERVALL_DER_BILANZIERTEN_MENGE_INKONSISTENT                                                                      // Z44
	RECHNUNGSEMPFAENGER_WIDERSPRICHT_DER_STEUERRECHTLICHEN_EINSCHAETZUNG_DES_RECHNUNGSSTELLERS                             // Z45
	ANGEGEBENE_QUOTES_AN_MARKTLOKATION_NICHT_VORHANDEN                                                                     // Z52
	RECHNUNGSABWICKLUNG_NICHT_VEREINBART                                                                                   // Z53
	COMDIS_WIRD_ABGELEHNT                                                                                                  // Z63
)
