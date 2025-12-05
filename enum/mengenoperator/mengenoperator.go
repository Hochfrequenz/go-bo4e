package mengenoperator

//go:generate stringer --type Mengenoperator
//go:generate jsonenums --type Mengenoperator
type Mengenoperator int

const (
	KLEINER_ALS Mengenoperator = iota + 1
	GROESSER_ALS
	GLEICH
)
