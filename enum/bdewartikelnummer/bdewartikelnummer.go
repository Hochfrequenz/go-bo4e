package bdewartikelnummer

// BDEWArtikelnummer is a readable representation of the article numbers of BDEW
//go:generate stringer --type BDEWArtikelnummer
//go:generate jsonenums --type BDEWArtikelnummer
type BDEWArtikelnummer int

const (
	// LEISTUNG has the BDEW article number 9990001000053
	LEISTUNG                                   BDEWArtikelnummer = iota + 1
	LEISTUNG_PAUSCHAL                                            // LEISTUNG_PAUSCHAL has the BDEW article number 9990001000079
	GRUNDPREIS                                                   // GRUNDPREIS has the BDEW article number 9990001000087
	REGELENERGIE_ARBEIT                                          // REGELENERGIE_ARBEIT has the BDEW article number 9990001000128
	REGELENERGIE_LEISTUNG                                        // REGELENERGIE_LEISTUNG has the BDEW article number 9990001000136
	NOTSTROMLIEFERUNG_ARBEIT                                     // NOTSTROMLIEFERUNG_ARBEIT has the BDEW article number 9990001000144
	NOTSTROMLIEFERUNG_LEISTUNG                                   // NOTSTROMLIEFERUNG_LEISTUNG has the BDEW article number 9990001000152
	RESERVENETZKAPAZITAET                                        // RESERVENETZKAPAZITAET has the BDEW article number 9990001000160
	RESERVELEISTUNG                                              // RESERVELEISTUNG has the BDEW article number 9990001000178
	ZUSAETZLICHE_ABLESUNG                                        // ZUSAETZLICHE_ABLESUNG has the BDEW article number 9990001000186
	PRUEFGEBUEHREN_AUSSERPLANMAESSIG                             // PRUEFGEBUEHREN_AUSSERPLANMAESSIG has the BDEW article number 9990001000219
	WIRKARBEIT                                                   // WIRKARBEIT has the BDEW article number 9990001000269
	SINGULAER_GENUTZTE_BETRIEBSMITTEL                            // SINGULAER_GENUTZTE_BETRIEBSMITTEL has the BDEW article number 9990001000285
	ABGABE_KWKG                                                  // ABGABE_KWKG has the BDEW article number 9990001000334
	ABSCHLAG                                                     // ABSCHLAG has the BDEW article number 9990001000376
	KONZESSIONSABGABE                                            // KONZESSIONSABGABE has the BDEW article number 9990001000417
	ENTGELT_FERNAUSLESUNG                                        // ENTGELT_FERNAUSLESUNG has the BDEW article number 9990001000433
	UNTERMESSUNG                                                 // UNTERMESSUNG has the BDEW article number 9990001000475
	BLINDMEHRARBEIT                                              // BLINDMEHRARBEIT has the BDEW article number 9990001000508
	ENTGELT_ABRECHNUNG                                           // ENTGELT_ABRECHNUNG has the BDEW article number 9990001000532
	SPERRKOSTEN                                                  // SPERRKOSTEN has the BDEW article number 9990001000540
	ENTSPERRKOSTEN                                               // ENTSPERRKOSTEN has the BDEW article number 9990001000558
	MAHNKOSTEN                                                   // MAHNKOSTEN has the BDEW article number 9990001000566
	MEHR_MINDERMENGEN                                            // MEHR_MINDERMENGEN has the BDEW article number 9990001000574
	INKASSOKOSTEN                                                // INKASSOKOSTEN has the BDEW article number 9990001000582
	BLINDMEHRLEISTUNG                                            // BLINDMEHRLEISTUNG has the BDEW article number 9990001000590
	ENTGELT_MESSUNG_ABLESUNG                                     // ENTGELT_MESSUNG_ABLESUNG has the BDEW article number 9990001000615
	ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK                   // ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK has the BDEW article number 9990001000623
	AUSGLEICHSENERGIE                                            // AUSGLEICHSENERGIE has the BDEW article number 9990001000631
	ZAEHLEINRICHTUNG                                             // ZAEHLEINRICHTUNG has the BDEW article number 9990001000649
	WANDLER_MENGENUMWERTER                                       // WANDLER_MENGENUMWERTER has the BDEW article number 9990001000657
	KOMMUNIKATIONSEINRICHTUNG                                    // KOMMUNIKATIONSEINRICHTUNG has the BDEW article number 9990001000665
	TECHNISCHE_STEUEREINRICHTUNG                                 // TECHNISCHE_STEUEREINRICHTUNG has the BDEW article number 9990001000673
	PARAGRAF_19_STROM_NEV_UMLAGE                                 // PARAGRAF_19_STROM_NEV_UMLAGE has the BDEW article number 9990001000681
	BEFESTIGUNGSEINRICHTUNG                                      // BEFESTIGUNGSEINRICHTUNG has the BDEW article number 9990001000699
	OFFSHORE_HAFTUNGSUMLAGE                                      // OFFSHORE_HAFTUNGSUMLAGE has the BDEW article number 9990001000706
	FIXE_ARBEITSENTGELTKOMPONENTE                                // FIXE_ARBEITSENTGELTKOMPONENTE has the BDEW article number 9990001000714
	FIXE_LEISTUNGSENTGELTKOMPONENTE                              // FIXE_LEISTUNGSENTGELTKOMPONENTE has the BDEW article number 9990001000722
	UMLAGE_ABSCHALTBARE_LASTEN                                   // UMLAGE_ABSCHALTBARE_LASTEN has the BDEW article number 9990001000730
	MEHRMENGE                                                    // MEHRMENGE has the BDEW article number 9990001000748
	MINDERMENGE                                                  // MINDERMENGE has the BDEW article number 9990001000756
	ENERGIESTEUER                                                // ENERGIESTEUER has the BDEW article number 9990001000764
	SMARTMETER_GATEWAY                                           // SMARTMETER_GATEWAY has the BDEW article number 9990001000772
	STEUERBOX                                                    // STEUERBOX has the BDEW article number 9990001000780
	MSB_INKL_MESSUNG                                             // MSB_INKL_MESSUNG has the BDEW article number 9990001000798
	AUSGLEICHSENERGIE_UNTERDECKUNG                               // AUSGLEICHSENERGIE_UNTERDECKUNG has the BDEW article number 9990001000805
)
