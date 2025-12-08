package arithmetischeoperation

// ArithmetischeOperation describes different arithmetic operations
//
//go:generate stringer --type ArithmetischeOperation
//go:generate jsonenums --type ArithmetischeOperation
type ArithmetischeOperation int

const (
	// ADDITION means addition
	ADDITION       ArithmetischeOperation = iota + 1
	SUBTRAKTION                           // SUBTRAKTION means subtraction
	MULTIPLIKATION                        // MULTIPLIKATION means multipication
	DIVISION                              // DIVISION means division (divisor)
	DIVIDEND                              // DIVIDEND means division (dividend)
	POSITIVWERT                           // POSITIVWERT means positive value
)
