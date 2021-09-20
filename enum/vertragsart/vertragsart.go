package vertragsart

//go:generate stringer --type Vertragsart
//go:generate jsonenums --type Vertragsart
// Vertragsart describes different kinds of contracs
type Vertragsart int

const (
	Energieliefervertrag       Vertragsart = iota + 1 // Energieliefervertrag
	Netznutzung                                       // Netznutzungsvertrag
	Bilanzierungsvertrag                              // Bilanzierungsvertrag
	Messstellenbetriebsvertrag                        // Messstellenabetriebsvertrag
	Buendelvertrag                                    // Bündelvertrag
)
