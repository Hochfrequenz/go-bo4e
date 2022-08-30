package gebiettyp

// Gebiettyp is a list of all possible types of domains (Gebiete)
//
//go:generate stringer --type Gebiettyp
//go:generate jsonenums --type Gebiettyp
type Gebiettyp int

const (
	// REGELZONE is a Regelzone
	REGELZONE              Gebiettyp = iota + 1
	MARKTGEBIET                      // MARKTGEBIET is a Marktgebiet
	BILANZIERUNGSGEBIET              // BILANZIERUNGSGEBIET is a Bilanzierungsgebiet
	VERTEILNETZ                      // VERTEILNETZ is a Verteilnetz
	TRANSPORTNETZ                    // TRANSPORTNETZ is a Transportnetz
	REGIONALNETZ                     // REGIONALNETZ is a Regionalnetz
	AREALNETZ                        // AREALNETZ is a Arealnetz
	GRUNDVERSORGUNGSGEBIET           // GRUNDVERSORGUNGSGEBIET is a Grundversorgungsgebiet
	VERSORGUNGSGEBIET                // VERSORGUNGSGEBIET is a Versorgungsgebiet
)
