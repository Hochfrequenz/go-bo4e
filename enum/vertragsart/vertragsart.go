package vertragsart

// Vertragsart describes different kinds of contracts
//
//go:generate stringer --type Vertragsart
//go:generate jsonenums --type Vertragsart
type Vertragsart int

const (
	// ENERGIELIEFERVERTRAG is an energy supply contract
	ENERGIELIEFERVERTRAG       Vertragsart = iota + 1
	NETZNUTZUNG                            // Deprecated: Netznutzungsvertrag (use NETZNUTZUNGSVERTRAG instead)
	BILANZIERUNGSVERTRAG                   // BILANZIERUNGSVERTRAG
	MESSSTELLENBETRIEBSVERTRAG             // Messstellenabetriebsvertrag
	BUENDELVERTRAG                         // Bündelvertrag
	NETZNUTZUNGSVERTRAG                    // NETZNUTZUNGSVERTRAG means Netznutzungsvertrag
	// because we like good docstrings
)
