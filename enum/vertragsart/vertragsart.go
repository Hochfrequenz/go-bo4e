package vertragsart

// Vertragsart describes different kinds of contracts
//go:generate stringer --type Vertragsart
//go:generate jsonenums --type Vertragsart
type Vertragsart int

const (
	Energieliefervertrag       Vertragsart = iota + 1 // Energieliefervertrag
	Netznutzung                                       // Netznutzungsvertrag
	Bilanzierungsvertrag                              // Bilanzierungsvertrag
	Messstellenbetriebsvertrag                        // Messstellenabetriebsvertrag
	Buendelvertrag                                    // Bündelvertrag
)
