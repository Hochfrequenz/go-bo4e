package erzeugungsart

// Auflistung der Erzeugungsarten von Energie
//
//go:generate stringer --type Erzeugungsart
//go:generate jsonenums --type Erzeugungsart

type Erzeugungsart int

const (
	KWK          Erzeugungsart = iota + 1 //Kraft-Waerme-Koppelung
	WIND                                  //Windkraft
	SOLAR                                 //Solarenergie
	KERNKRAFT                             //Kernkraft
	WASSER                                //Wasserkraft
	GEOTHERMIE                            //Geothermie
	BIOMASSE                              //Biomasse
	KOHLE                                 //Kohle
	GAS                                   //Erdgas
	SONSTIGE                              //Sonstige
	SONSTIGE_EEG                          //Sonstige nach EEG
)
