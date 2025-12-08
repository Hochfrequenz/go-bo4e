package dienstleistungstyp

// Dienstleistungstyp ist eine Auflistung möglicher abzurechnender Dienstleistungen.
//
//go:generate stringer --type Dienstleistungstyp
//go:generate jsonenums --type Dienstleistungstyp
type Dienstleistungstyp int

const (
	DATENBEREITSTELLUNG_TAEGLICH                Dienstleistungstyp = iota + 1 // Datenbereitstellung täglich
	DATENBEREITSTELLUNG_WOECHENTLICH                                          // Datenbereitstellung wöchentlich
	DATENBEREITSTELLUNG_MONATLICH                                             // Datenbereitstellung monatlich
	DATENBEREITSTELLUNG_JAEHRLICH                                             // Datenbereitstellung jährlich
	DATENBEREITSTELLUNG_HISTORISCHE_LG                                        // Datenbereitstellung historischer Lastgänge
	DATENBEREITSTELLUNG_STUENDLICH                                            // Datenbereitstellung stündlich
	DATENBEREITSTELLUNG_VIERTELJAEHRLICH                                      // Datenbereitstellung vierteljährlich
	DATENBEREITSTELLUNG_HALBJAEHRLICH                                         // Datenbereitstellung halbjährlich
	DATENBEREITSTELLUNG_MONATLICH_ZUSAETZLICH                                 // Datenbereitstellung monatlich zusätzlich
	DATENBEREITSTELLUNG_EINMALIG                                              // Datenbereitstellung einmalig
	AUSLESUNG_2X_TAEGLICH_FERNAUSLESUNG                                       // Auslesung 2x täglich mittels Fernauslesung
	AUSLESUNG_3X_TAEGLICH_FERNAUSLESUNG                                       // Auslesung 3x täglich mittels Fernauslesung
	AUSLESUNG_TAEGLICH_FERNAUSLESUNG                                          // Auslesung täglich mittels Fernauslesung
	AUSLESUNG_MANUELL_MSB                                                     // Auslesung, manuell vom Messstellenbetreiber vorgenommen
	AUSLESUNG_MONATLICH_FERNAUSLESUNG                                         // Auslesung monatlich bei mittels Fernauslesung
	AUSLESUNG_JAEHRLICH_FERNAUSLESUNG                                         // Auslesung jährlich bei SLP mittels Fernauslesung
	AUSLESUNG_MDE                                                             // Auslesung mit mobiler Daten Erfassung (MDE)
	ABLESUNG_MONATLICH                                                        // Ablesung monatlich
	ABLESUNG_VIERTELJAEHRLICH                                                 // Ablesung vierteljährlich
	ABLESUNG_HALBJAEHRLICH                                                    // Ablesung halbjährlich
	ABLESUNG_JAEHRLICH                                                        // Ablesung jährlich
	AUSLESUNG_FERNAUSLESUNG                                                   // Auslesung mittels Fernauslesung
	ABLESUNG_ZUSAETZLICH_MSB                                                  // Ablesung, zusätzlich vom Messstellenbetreiber vorgenommen
	ABLESUNG_ZUSAETZLICH_KUNDE                                                // Ablesung SLP, zusätzlich vom Kunden vorgenommen
	AUSLESUNG_FERNAUSLESUNG_ZUSAETZLICH_MSB                                   // Auslesung, mittels Fernauslesung, zusätzlich vom Messstellenbetreiber vorgenommen
	AUSLESUNG_MOATLICH_FERNAUSLESUNG                                          // Auslesung monatlich mittels Fernauslesung
	AUSLESUNG_STUENDLICH_FERNAUSLESUNG                                        // Auslesung stündlich mittels Fernauslesung
	AUSLESUNG_TEMPERATURMENGENUMWERTER                                        // Auslesung Temperaturmengenumwerter
	AUSLESUNG_ZUSTANDSMENGENUMWERTER                                          // Auslesung Zustandsmengenumwerter
	AUSLESUNG_SYSTEMMENGENUMWERTER                                            // Auslesung Systemmengenumwerter
	AUSLESUNG_VORGANG                                                         // Auslesung je Vorgang
	AUSLESUNG_KOMPAKTMENGENUMWERTER                                           // Auslesung Kompaktmengenumwerter
	SPERRUNG                                                                  // Sperrung einer Abnahmestelle
	ENTSPERRUNG                                                               // Entsperrung einer Abnahmestelle
	MAHNKOSTEN                                                                // Mahnkosten
	INKASSOKOSTEN                                                             // Inkassokosten
	AUSLESUNG_LGK_MANUELL_MSB                                                 // Auslesung LGK, manuell vom MSB vorgenommen
	AUSLESUNG_MONATLICH_SLP_FERNAUSLESUNG                                     // Auslesung monatlich SLP mittels Fernauslesung
	AUSLESUNG_JAEHRLICH_SLP_FERNAUSLESUNG                                     // Auslesung jaehrlich SLP mittels Fernauslesung
	AUSLESUNG_MDE_SLP                                                         // Auslesung SLP mittels MDE
	ABLESUNG_MONATLICH_SLP                                                    // Ablesung monatlich SLP
	ABLESUNG_VIERTELJAEHRLICH_SLP                                             // Ablesung vierteljaehrlich SLP
	ABLESUNG_HALBJAEHRLICH_SLP                                                // Ablesung halbjaehrlich SLP
	ABLESUNG_JAEHRLICH_SLP                                                    // Ablesung jaehrlich SLP
	AUSLESUNG_SLP_FERNAUSLESUNG                                               // Auslesung SLP mittels Fernauslesung
	ABLESUNG_SLP_ZUSAETZLICH_MSB                                              // Ablesung SLP, zusaetzlich vom MSB vorgenommen
	ABLESUNG_SLP_ZUSAETZLICH_KUNDE                                            // Ablesung SLP, zusaetzlich vom Kunden vorgenommen
	AUSLESUNG_LGK_FERNAUSLESUNG_ZUSAETZLICH_MSB                               // Auslesung LGK Fernauslesung, zusaetzlich vom MSB
	ABLESUNG_MONATLICH_LGK                                                    // Ablesung monatlich LGK
	AUSLESUNG_TEMERATURMENGENUMWERTER                                         // Auslesung Temperaturmengenumwerter (typo in C# lib)
	AUSLESUNG_VORGANG_SLP                                                     // Auslesung SLP je Vorgang
	AUSLESUUNG_KOMPAKTMENGENUMWERTER                                          // Auslesung Kompaktmengenumwerter (typo in C# lib)
	AUSLESUNG_MDE_LGK                                                         // Auslesung LGK mittels MDE
	SPERRUNG_SLP                                                              // Sperrung SLP
	ENTSPERRUNG_SLP                                                           // Entsperrung SLP
	SPERRUNG_RLM                                                              // Sperrung RLM
	ENTSPERRUNG_RLM                                                           // Entsperrung RLM
)
