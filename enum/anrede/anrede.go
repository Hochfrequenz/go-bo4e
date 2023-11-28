package anrede

// Anrede ist eine Auflistung möglicher abzurechnender Dienstleistungen.
//
//go:generate stringer --type Anrede
//go:generate jsonenums --type Anrede
type Anrede int

const (
	// HERR (male)
	HERR                    Anrede = iota + 1
	FRAU                           // FRAU (female)
	DIVERS                         // DIVERS (non-binary) - Unofficial / Customer requirement, not part of the BO4E standard yet.
	EHELEUTE                       // EHELEUTE (married couple)
	FIRMA                          // FIRMA (company)
	WOHNGEMEINSCHAFT               // WOHNGEMEINSCHAFT is a share flat/WG - Unofficial / Customer requirement, not part of the BO4E standard yet.
	INDIVIDUELL                    // INDIVIDUELL festgelegt
	FAMILIE                        // FAMILIE (family) - Unofficial / Customer requirement, not part of the BO4E standard yet.
	ERBENGEMEINSCHAFT              // ERBENGEMEINSCHAFT (community of heirs, community of co-heirs, community of joint heirs) - Unofficial / Customer requirement, not part of the BO4E standard yet.
	GRUNDSTUECKGEMEINSCHAFT        // GRUNDSTUECKGEMEINSCHAFT - Unofficial / Customer requirement, not part of the BO4E standard yet.
	DR                             // DR (doctor) - Unofficial / Customer requirement, not part of the BO4E standard yet.
)
