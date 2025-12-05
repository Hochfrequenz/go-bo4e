package inbetriebsetzungtechnischeressource

// InbetriebsetzungTechnischeRessource is an enum
//
//go:generate stringer --type InbetriebsetzungTechnischeRessource
//go:generate jsonenums --type InbetriebsetzungTechnischeRessource
type InbetriebsetzungTechnischeRessource int

const (
	NACH_2023 InbetriebsetzungTechnischeRessource = iota + 1
	VOR_2024
)
