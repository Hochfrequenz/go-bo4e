package aggregationsverantwortung

// Aggregationsverantwortung describes who is responsible for aggregations
//go:generate stringer --type Aggregationsverantwortung
//go:generate jsonenums --type Aggregationsverantwortung
type Aggregationsverantwortung int

const (
	// UENB means Übertragungsnetzbetreiber / long distance grid operator
	UENB Aggregationsverantwortung = iota + 1
	VNB                            // VNB means Verteilnetzbetreiber / local grid operator
)
