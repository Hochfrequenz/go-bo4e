package geraetetyp

//go:generate stringer --type Geraetetyp
//go:generate jsonenums --type Geraetetyp
type Geraetetyp int

const (
	WECHSELSTROMZAEHLER             Geraetetyp = iota // Wechselstromzähler
	DREHSTROMZAEHLER                                  // Drehstromzähler
	ZWEIRICHTUNGSZAEHLER                              // Zweirichtungszähler
	RLM_ZAEHLER                                       // RLM-Zähler
	BALGENGASZAEHLER                                  // Balgengaszähler
	MAXIMUMZAEHLER                                    // Maximumzähler (Schleppzähler)
	MULTIPLEXANLAGE                                   // Multiplexeranlage
	PAUSCHALANLAGE                                    // Pauschalanlagen
	VERSTAERKERANLAGE                                 // Verstärkeranlage
	SUMMATIONSGERAET                                  // Summationsgerät
	IMPULSGEBER                                       // Impulsgeber
	EDL_21_ZAEHLERAUFSATZ                             // EDL 21 Zähleraufsatz für Zähler
	VIER_QUADRANTEN_LASTGANGZAEHLER                   // Vier-Quadranten-Lastgangzähler
	MENGENUMWERTER                                    // Mengenumwerter
	STROMWANDLER                                      // Stromwandler
	SPANNUNGSWANDLER                                  // Spannungs-Wandler
	KOMBIMESSWANDLER                                  // Kombimesswandler
	BLOCKSTROMWANDLER                                 // Blockstromwandler
	DATENLOGGER                                       // Datenlogger
	KOMMUNIKATIONSANSCHLUSS                           // Kommunikationsanschluss
	MODEM                                             // Modem
	TELEKOMMUNIKATIONSEINRICHTUNG                     // vom Messstellenbetreiber beigestellte Telekommunikationseinrichtung (Telefonanschluss)
	DREHKOLBENGASZAEHLER                              // Drehkolbengaszähler
	TURBINENRADGASZAEHLER                             // Turbinenradgaszähler
	ULTRASCHALLZAEHLER                                // Ultraschallzähler
	WIRBELGASZAEHLER                                  // Wirbelgaszähler
	MODERNE_MESSEINRICHTUNG                           // moderne Messeinrichtung
	INTELLIGENTES_MESSYSTEM                           // Intelligentes Messystem
	ELEKTRONISCHER_HAUSHALTSZAEHLER                   // elektronischer Haushaltszähler
	STEUEREINRICHTUNG                                 // Steuereinrichtung
	TARIFSCHALTGERAET                                 // Tarifschaltgerät
	RUNDSTEUEREMPFAENGER                              // Rundsteuerempfänger
	OPTIONALE_ZUS_ZAEHLEINRICHTUNG                    // optionale zusätzliche Zähleinrichtung
	MESSWANDLERSATZ_IMS_MME                           // Messwandlersatz Strom iMS und mME, NSP
	KOMBIMESSWANDLER_IMS_MME                          // Kombimesswandlersatz (Strom u. Spg) iMS und mME
	TARIFSCHALTGERAET_IMS_MME                         // Tarifschaltung iMS und mME
	RUNDSTEUEREMPFAENGER_IMS_MME                      // Rundsteuerempfänger iMS und mME
	TEMPERATUR_KOMPENSATION                           // Temperaturkompensation
	HOECHSTBELASTUNGS_ANZEIGER                        // Höchsbelastungsanzeiger
	SONSTIGES_GERAET                                  // Sonstiges Gerät
	PREPAYMENTZAEHLER                                 // Prepaymentzähler
	EDL_21                                            // EDL21
	_96_H_ZAEHLER                                     // 96 h Zähler
	EDL_40_ZAEHLERAUFSATZ                             // EDL 40 Zähleraufsatz für Zähler
)
