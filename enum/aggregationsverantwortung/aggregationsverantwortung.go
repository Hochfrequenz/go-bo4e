package aggregationsverantwortung

// Aggregationsverantwortung describes who is responsible for aggregations
// Note that this enum is not official BO4E standard (yet)!
//
//go:generate stringer --type Aggregationsverantwortung
//go:generate jsonenums --type Aggregationsverantwortung
type Aggregationsverantwortung int

const (
	// UENB means Übertragungsnetzbetreiber / long distance grid operator
	UENB Aggregationsverantwortung = iota + 1
	VNB                            // VNB means Verteilnetzbetreiber / local grid operator
)
