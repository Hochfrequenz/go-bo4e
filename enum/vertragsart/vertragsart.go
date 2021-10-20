package vertragsart

// Vertragsart describes different kinds of contracts
//go:generate stringer --type Vertragsart
//go:generate jsonenums --type Vertragsart
type Vertragsart int

const (
	ENERGIELIEFERVERTRAG       Vertragsart = iota + 1 // ENERGIELIEFERVERTRAG
	NETZNUTZUNG                                       // Netznutzungsvertrag
	BILANZIERUNGSVERTRAG                              // BILANZIERUNGSVERTRAG
	MESSSTELLENBETRIEBSVERTRAG                        // Messstellenabetriebsvertrag
	BUENDELVERTRAG                                    // Bündelvertrag
)
