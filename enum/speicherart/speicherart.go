package speicherart

// Im Falle der Speicher bei <see cref="TechnischeRessourceNutzung"/>, eine genauere Angabe über die Art der Speicher zu definieren
//
//go:generate stringer --type Speicherart
//go:generate jsonenums --type Speicherart

type Speicherart int

const (
	WASSERSTOFFSPEICHER  Speicherart = iota + 1 //ZF7: Wasserstoffspeicher
	PUMPSPEICHER                                //ZF8: Pumpspeicher
	BATTERIESPEICHER                            //ZF9: Batteriespeicher
	SONSTIGE_SPEICHERART                        //ZG6: Sonstige Speicherart
)
