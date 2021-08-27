package messwertstatuszusatz

//go:generate stringer --type Messwertstatuszusatz
//go:generate jsonenums --type Messwertstatuszusatz
type Messwertstatuszusatz int // Aufzählung von zusätzlichen Informationen zum Status, beispielsweise in Lastgängen oder Zählwerkständen.

const (
	Z84_LEERSTAND                                      Messwertstatuszusatz = iota // Leerstand
	Z85_REALERZAEHLERUEBERLAUFGEPRUEFT                                             // Realer Zaehlerüberlauf geprueft
	Z86_PLAUSIBELWGKONTROLLABLESUNG                                                // Plausibel wg. Kontrollablesung
	Z87_PLAUSIBELWGKUNDENHINWEIS                                                   // Plausibel wg. Kundenhinweis
	ZC3_AUSTAUSCHDESERSATZWERTES                                                   // Austausch des Ersatzwertes
	Z88_VERGLEICHSMESSUNG_GEEICHT                                                  // Vergleichsmessung (geeicht)
	Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT                                            // Vergleichsmessung (nicht geeicht)
	Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN                                      // Messwertnachbildung aus geeichten Werten
	Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN                                 // Messwertnachbildung aus nicht geeichten Werten
	Z92_INTERPOLATION                                                              // Interpolation
	Z93_HALTEWERT                                                                  // Haltewert
	Z94_BILANZIERUNGNETZABSCHNITT                                                  // Bilanzierung Netzabschnitt
	Z95_HISTORISCHEMESSWERTE                                                       // Historische Messwerte
	ZJ2_STATISTISCHEMETHODE                                                        // Statistische Methode
	Z74_KEINZUGANG                                                                 // kein Zugang
	Z75_KOMMUNIKATIONSSTOERUNG                                                     // Kommunikationsstörung
	Z76_NETZAUSFALL                                                                // Netzausfall
	Z77_SPANNUNGSAUSFALL                                                           // Spannungsausfall
	Z78_GERAETEWECHSEL                                                             // Gerätewechsel
	Z79_KALIBRIERUNG                                                               // Kalibrierung
	Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN                             // Gerät arbeitet ausserhalb der Betriebsbedingungen
	Z81_MESSEINRICHTUNGGESTOERT_DEFEKT                                             // Messeinrichtung gestört/defekt
	Z82_UNSICHERHEITMESSUNG                                                        // Unsicherheit Messung
	Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK                                      // Berücksichtigung Störmengenzählwerk
	Z99_MENGENUMWERTUNGUNVOLLSTAENDIG                                              // Mengenumwertung unvollständig
	ZA0_UHRZEITGESTELLT_SYNCHRONISATION                                            // Uhrzeit gestellt /Synchronisation
	ZA1_MESSWERTUNPLAUSIBEL                                                        // Messwert unplausibel
	ZC2_TARIFSCHALTGERAETDEFEKT                                                    // Tarifschaltgeraet defekt
	ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND                                           // Impulswertigkeit nicht ausreichend
	ZA3_FALSCHERWANDLERFAKTOR                                                      // Falscher Wandlerfaktor
	ZA4_FEHLERHAFTEABLESUNG                                                        // Fehlerhafte Ablesung
	ZA5_AENDERUNGDERBERECHNUNG                                                     // Änderung der Berechnung
	ZA6_UMBAUDERMESSLOKATION                                                       // Umbau der Messlokation
	ZA7_DATENBEARBEITUNGSFEHLER                                                    // Datenbearbeitungsfehler
	ZA8_BRENNWERTKORREKTUR                                                         // Brennwertkorrektur
	ZA9_ZZAHL_KORREKTUR                                                            // Z-Zahl-Korrektur
	ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG                                             // Störung / Defekt Messeinrichtung
	ZB9_AENDERUNGTARIFSCHALTZEITEN                                                 // Änderung Tarifschaltzeiten
	ZG3_UMSTELLUNGGASQUALITAET                                                     // Umstellung Gasqualität
)
