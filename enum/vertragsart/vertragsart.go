package vertragsart

//go:generate stringer --type Vertragsart
//go:generate jsonenums --type Vertragsart
// Vertragsart describes different kinds of contracs
type Vertragsart int

const (
	ENERGIELIEFERVERTRAG       Vertragsart = iota + 1 // Energieliefervertrag
	NETZNUTZUNGSVERTRAG                               // Netznutzungsvertrag
	BILANZIERUNGSVERTRAG                              // Bilanzierungsvertrag
	MESSSTELLENBETRIEBSVERTRAG                        // Messstellenabetriebsvertrag
	BUENDELVERTRAG                                    // Bündelvertrag
)
