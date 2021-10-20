package messwertstatuszusatz

// Messwertstatuszusatz ist eine Aufzählung von zusätzlichen Informationen zum Status, beispielsweise in Lastgängen oder Zählwerkständen.
//go:generate stringer --type Messwertstatuszusatz
//go:generate jsonenums --type Messwertstatuszusatz
type Messwertstatuszusatz int

const (
	// Z84_LEERSTAND ist Leerstand
	Z84_LEERSTAND                                      Messwertstatuszusatz = iota + 1
	Z85_REALERZAEHLERUEBERLAUFGEPRUEFT                                      // Z85_REALERZAEHLERUEBERLAUFGEPRUEFT ist Realer Zaehlerüberlauf geprueft
	Z86_PLAUSIBELWGKONTROLLABLESUNG                                         // Z86_PLAUSIBELWGKONTROLLABLESUNG ist Plausibel wg. Kontrollablesung
	Z87_PLAUSIBELWGKUNDENHINWEIS                                            // Z87_PLAUSIBELWGKUNDENHINWEIS ist Plausibel wg. Kundenhinweis
	ZC3_AUSTAUSCHDESERSATZWERTES                                            // ZC3_AUSTAUSCHDESERSATZWERTES ist Austausch des Ersatzwertes
	Z88_VERGLEICHSMESSUNG_GEEICHT                                           // Z88_VERGLEICHSMESSUNG_GEEICHT ist Vergleichsmessung (geeicht)
	Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT                                     // Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT ist Vergleichsmessung (nicht geeicht)
	Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN                               // Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN ist Messwertnachbildung aus geeichten Werten
	Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN                          // Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN ist Messwertnachbildung aus nicht geeichten Werten
	Z92_INTERPOLATION                                                       // Z92_INTERPOLATION ist Interpolation
	Z93_HALTEWERT                                                           // Z93_HALTEWERT ist Haltewert
	Z94_BILANZIERUNGNETZABSCHNITT                                           // Z94_BILANZIERUNGNETZABSCHNITT ist Bilanzierung Netzabschnitt
	Z95_HISTORISCHEMESSWERTE                                                // Z95_HISTORISCHEMESSWERTE ist Historische Messwerte
	ZJ2_STATISTISCHEMETHODE                                                 // ZJ2_STATISTISCHEMETHODE ist Statistische Methode
	Z74_KEINZUGANG                                                          // Z74_KEINZUGANG ist kein Zugang
	Z75_KOMMUNIKATIONSSTOERUNG                                              // Z75_KOMMUNIKATIONSSTOERUNG ist Kommunikationsstörung
	Z76_NETZAUSFALL                                                         // Z76_NETZAUSFALL ist Netzausfall
	Z77_SPANNUNGSAUSFALL                                                    // Z77_SPANNUNGSAUSFALL ist Spannungsausfall
	Z78_GERAETEWECHSEL                                                      // Z78_GERAETEWECHSEL ist Gerätewechsel
	Z79_KALIBRIERUNG                                                        // Z79_KALIBRIERUNG ist Kalibrierung
	Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN                      // Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN ist Gerät arbeitet ausserhalb der Betriebsbedingungen
	Z81_MESSEINRICHTUNGGESTOERT_DEFEKT                                      // Z81_MESSEINRICHTUNGGESTOERT_DEFEKT ist Messeinrichtung gestört/defekt
	Z82_UNSICHERHEITMESSUNG                                                 // Z82_UNSICHERHEITMESSUNG ist Unsicherheit Messung
	Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK                               // Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK ist Berücksichtigung Störmengenzählwerk
	Z99_MENGENUMWERTUNGUNVOLLSTAENDIG                                       // Z99_MENGENUMWERTUNGUNVOLLSTAENDIG ist Mengenumwertung unvollständig
	ZA0_UHRZEITGESTELLT_SYNCHRONISATION                                     // ZA0_UHRZEITGESTELLT_SYNCHRONISATION ist Uhrzeit gestellt /Synchronisation
	ZA1_MESSWERTUNPLAUSIBEL                                                 // ZA1_MESSWERTUNPLAUSIBEL ist Messwert unplausibel
	ZC2_TARIFSCHALTGERAETDEFEKT                                             // ZC2_TARIFSCHALTGERAETDEFEKT ist Tarifschaltgeraet defekt
	ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND                                    // ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND ist Impulswertigkeit nicht ausreichend
	ZA3_FALSCHERWANDLERFAKTOR                                               // ZA3_FALSCHERWANDLERFAKTOR ist Falscher Wandlerfaktor
	ZA4_FEHLERHAFTEABLESUNG                                                 // ZA4_FEHLERHAFTEABLESUNG ist Fehlerhafte Ablesung
	ZA5_AENDERUNGDERBERECHNUNG                                              // ZA5_AENDERUNGDERBERECHNUNG ist Änderung der Berechnung
	ZA6_UMBAUDERMESSLOKATION                                                // ZA6_UMBAUDERMESSLOKATION ist Umbau der Messlokation
	ZA7_DATENBEARBEITUNGSFEHLER                                             // ZA7_DATENBEARBEITUNGSFEHLER ist Datenbearbeitungsfehler
	ZA8_BRENNWERTKORREKTUR                                                  // ZA8_BRENNWERTKORREKTUR ist Brennwertkorrektur
	ZA9_ZZAHL_KORREKTUR                                                     // ZA9_ZZAHL_KORREKTUR ist Z-Zahl-Korrektur
	ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG                                      // ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG ist Störung / Defekt Messeinrichtung
	ZB9_AENDERUNGTARIFSCHALTZEITEN                                          // ZB9_AENDERUNGTARIFSCHALTZEITEN ist Änderung Tarifschaltzeiten
	ZG3_UMSTELLUNGGASQUALITAET                                              // ZG3_UMSTELLUNGGASQUALITAET ist Umstellung Gasqualität
)
