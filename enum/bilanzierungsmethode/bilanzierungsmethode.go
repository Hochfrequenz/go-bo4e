package bilanzierungsmethode

// Bilanzierungsmethode describes the different accounting methods and basics
//go:generate stringer --type Bilanzierungsmethode
//go:generate jsonenums --type Bilanzierungsmethode
type Bilanzierungsmethode int

const (
	// RLM is a Registrierende Leistungsmessung
	RLM           Bilanzierungsmethode = iota + 1
	SLP                           // SLP is a Standard Lastprofil
	TLP_GEMEINSAM                 // TLP_GEMEINSAM is a TLP gemeinsame Messung
	TLP_GETRENNT                  // TLP_GETRENNT is a TLP getrennte Messung
	PAUSCHAL                      // PAUSCHAL is a Pauschale Betrachtung (Band)
)
