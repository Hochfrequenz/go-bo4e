package sperrauftragsart

//go:generate stringer --type Sperrauftragsart
//go:generate jsonenums --type Sperrauftragsart
type Sperrauftragsart int

const (
	SPERREN Sperrauftragsart = iota + 1
	ENTSPERREN
)
