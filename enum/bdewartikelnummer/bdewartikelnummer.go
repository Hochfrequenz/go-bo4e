package bdewartikelnummer

// BDEWArtikelnummer is a readable representation of the article numbers of BDEW
//go:generate stringer --type BDEWArtikelnummer
//go:generate jsonenums --type BDEWArtikelnummer
type BDEWArtikelnummer int

const (
	Leistung                               BDEWArtikelnummer = iota + 1 // 9990001000053
	Leistungpauschal                                                    // 9990001000079
	Grundpreis                                                          // 9990001000087
	Regelenergiearbeit                                                  // 9990001000128
	Regelenergieleistung                                                // 9990001000136
	Notstromlieferungarbeit                                             // 9990001000144
	Notstromlieferungleistung                                           // 9990001000152
	Reservenetzkapazitaet                                               // 9990001000160
	Reserveleistung                                                     // 9990001000178
	Zusaetzlicheablesung                                                // 9990001000186
	Pruefgebuehrenausserplanmaessig                                     // 9990001000219
	Wirkarbeit                                                          // 9990001000269
	Singulaergenutztebetriebsmittel                                     // 9990001000285
	Abgabekwkg                                                          // 9990001000334
	Abschlag                                                            // 9990001000376
	Konzessionsabgabe                                                   // 9990001000417
	Entgelt_fernauslesung                                               // 9990001000433
	Untermessung                                                        // 9990001000475
	Blindmehrarbeit                                                     // 9990001000508
	Entgelt_abrechnung                                                  // 9990001000532
	Sperrkosten                                                         // 9990001000540
	Entsperrkosten                                                      // 9990001000558
	Mahnkosten                                                          // 9990001000566
	MehrMmindermengen                                                   // 9990001000574
	Inkassokosten                                                       // 9990001000582
	Blindmehrleistung                                                   // 9990001000590
	EntgeltmessungAblesung                                              // 9990001000615
	EntgelteinbauBetriebWartungMesstechnik                              // 9990001000623
	Ausgleichsenergie                                                   // 9990001000631
	Zaehleinrichtung                                                    // 9990001000649
	WandlerMengenumwerter                                               // 9990001000657
	Kommunikationseinrichtung                                           // 9990001000665
	TechnischeSteuereinrichtung                                         // 9990001000673
	Paragraf19StromNevUmlage                                            // 9990001000681
	Befestigungseinrichtung                                             // 9990001000699
	OffshoreHaftungsumlage                                              // 9990001000706
	FixeArbeitsentgeltkomponente                                        // 9990001000714
	FixeLeistungsentgeltkomponente                                      // 9990001000722
	UmlageAbschaltbareLasten                                            // 9990001000730
	Mehrmenge                                                           // 9990001000748
	Mindermenge                                                         // 9990001000756
	Energiesteuer                                                       // 9990001000764
	SmartmeterGateway                                                   // 9990001000772
	Steuerbox                                                           // 9990001000780
	MsbInklMessung                                                      // 9990001000798
)
