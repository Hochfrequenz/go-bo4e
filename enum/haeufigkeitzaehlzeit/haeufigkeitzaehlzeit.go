package haeufigkeitzaehlzeit

//go:generate stringer --type HaeufigkeitZaehlzeit
//go:generate jsonenums --type HaeufigkeitZaehlzeit
type HaeufigkeitZaehlzeit int

const (
	EINMALIG HaeufigkeitZaehlzeit = iota + 1
	JAEHRLICH
)
