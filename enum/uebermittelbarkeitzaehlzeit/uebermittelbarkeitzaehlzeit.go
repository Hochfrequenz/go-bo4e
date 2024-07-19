package uebermittelbarkeitzaehlzeit

//go:generate stringer --type UebermittelbarkeitZaehlzeit
//go:generate jsonenums --type UebermittelbarkeitZaehlzeit
type UebermittelbarkeitZaehlzeit int

const (
	ELEKTRONISCH UebermittelbarkeitZaehlzeit = iota + 1
	NICHT_ELEKTRONISCH
)
