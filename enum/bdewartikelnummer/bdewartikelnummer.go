package bdewartikelnummer

// BDEWArtikelnummer is a readable representation of the article numbers of BDEW
//go:generate stringer --type BDEWArtikelnummer
//go:generate jsonenums --type BDEWArtikelnummer
type BDEWArtikelnummer int

const (
	LEISTUNG                               BDEWArtikelnummer = iota + 1 // LEISTUNG has the BDEW article number 9990001000053
	LEISTUNGPAUSCHAL                                                    // LEISTUNGPAUSCHAL has the BDEW article number 9990001000079
	GRUNDPREIS                                                          // GRUNDPREIS has the BDEW article number 9990001000087
	REGELENERGIEPREIS                                                   // REGELENERGIEPREIS has the BDEW article number 9990001000128
	REGELENERGIELEISTUNG                                                // REGELENERGIELEISTUNG has the BDEW article number 9990001000136
	NOTSTROMLIEFERUNGARBEIT                                             // NOTSTROMLIEFERUNGARBEIT has the BDEW article number 9990001000144
	NOTSTROMLIEFERUNGLEISTUNG                                           // NOTSTROMLIEFERUNGLEISTUNG has the BDEW article number 9990001000152
	RESERVENETZKAPAZITAET                                               // RESERVENETZKAPAZITAET has the BDEW article number 9990001000160
	RESERVELEISTUNG                                                     // RESERVELEISTUNG has the BDEW article number 9990001000178
	ZUSAETZLICHEABLESUNG                                                // ZUSAETZLICHEABLESUNG has the BDEW article number 9990001000186
	PRUEFGEBUEHRENAUSSERPLANMAESSIG                                     // PRUEFGEBUEHRENAUSSERPLANMAESSIG has the BDEW article number 9990001000219
	WIRKARBEIT                                                          // WIRKARBEIT has the BDEW article number 9990001000269
	SINGULAERGENUTZTEBETRIEBSMITTEL                                     // SINGULAERGENUTZTEBETRIEBSMITTEL has the BDEW article number 9990001000285
	ABGABEKWKG                                                          // ABGABEKWKG has the BDEW article number 9990001000334
	ABSCHLAG                                                            // ABSCHLAG has the BDEW article number 9990001000376
	KONZESSIONSABGABE                                                   // KONZESSIONSABGABE has the BDEW article number 9990001000417
	ENTGELT_FERNAUSLESUNG                                               // ENTGELT_FERNAUSLESUNG has the BDEW article number 9990001000433
	UNTERMESSUNG                                                        // UNTERMESSUNG has the BDEW article number 9990001000475
	BLINDMEHRARBEIT                                                     // BLINDMEHRARBEIT has the BDEW article number 9990001000508
	ENTGELT_ABRECHNUNG                                                  // ENTGELT_ABRECHNUNG has the BDEW article number 9990001000532
	SPERRKOSTEN                                                         // SPERRKOSTEN has the BDEW article number 9990001000540
	ENTSPERRKOSTEN                                                      // ENTSPERRKOSTEN has the BDEW article number 9990001000558
	MAHNKOSTEN                                                          // MAHNKOSTEN has the BDEW article number 9990001000566
	MEHRMMINDERMENGEN                                                   // MEHRMMINDERMENGEN has the BDEW article number 9990001000574
	INKASSOKOSTEN                                                       // INKASSOKOSTEN has the BDEW article number 9990001000582
	BLINDMEHRLEISTUNG                                                   // BLINDMEHRLEISTUNG has the BDEW article number 9990001000590
	ENTGELTMESSUNGABLESUNG                                              // ENTGELTMESSUNGABLESUNG has the BDEW article number 9990001000615
	ENTGELTEINBAUBETRIEBWARTUNGMESSTECHNIK                              // ENTGELTEINBAUBETRIEBWARTUNGMESSTECHNIK has the BDEW article number 9990001000623
	AUSGLEICHSENERGIE                                                   // AUSGLEICHSENERGIE has the BDEW article number 9990001000631
	ZAEHLEINRICHTUNG                                                    // ZAEHLEINRICHTUNG has the BDEW article number 9990001000649
	WANDLERMENGENUMWERTER                                               // WANDLERMENGENUMWERTER has the BDEW article number 9990001000657
	KOMMUNIKATIONSEINRICHTUNG                                           // KOMMUNIKATIONSEINRICHTUNG has the BDEW article number 9990001000665
	TECHNISCHESTEUEREINRICHTUNG                                         // TECHNISCHESTEUEREINRICHTUNG has the BDEW article number 9990001000673
	PARAGRAF19STROMNEVUMLAGE                                            // PARAGRAF19STROMNEVUMLAGE has the BDEW article number 9990001000681
	BEFESTIGUNGSEINRICHTUNG                                             // BEFESTIGUNGSEINRICHTUNG has the BDEW article number 9990001000699
	OFFSHOREHAFTUNGSUMLAGE                                              // OFFSHOREHAFTUNGSUMLAGE has the BDEW article number 9990001000706
	FIXEARBEITSENTGELTKOMPONENTE                                        // FIXEARBEITSENTGELTKOMPONENTE has the BDEW article number 9990001000714
	FIXELEISTUNGSENTGELTKOMPONENTE                                      // FIXELEISTUNGSENTGELTKOMPONENTE has the BDEW article number 9990001000722
	UMLAGEABSCHALTBARELASTEN                                            // UMLAGEABSCHALTBARELASTEN has the BDEW article number 9990001000730
	MEHRMENGE                                                           // MEHRMENGE has the BDEW article number 9990001000748
	MINDERMENGE                                                         // MINDERMENGE has the BDEW article number 9990001000756
	ENERGIESTEUER                                                       // ENERGIESTEUER has the BDEW article number 9990001000764
	SMARTMETERGATEWAY                                                   // SMARTMETERGATEWAY has the BDEW article number 9990001000772
	STEUERBOX                                                           // STEUERBOX has the BDEW article number 9990001000780
	MSBINKLMESSUNG                                                      // MSBINKLMESSUNG has the BDEW article number 9990001000798
)
