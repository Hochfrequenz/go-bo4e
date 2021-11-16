package rufnummernart

// Rufnummernart bildet verschiedene Rufnummerntypen ab
//go:generate stringer --type Rufnummernart
//go:generate jsonenums --type Rufnummernart
type Rufnummernart int

const (
	// RUF_ZENTRALE bedeutet Rufnummer Zentrale
	RUF_ZENTRALE  Rufnummernart = iota + 1
	FAX_ZENTRALE                // FAX_ZENTRALE bedeutet Faxnummer Zentrale
	SAMMELRUF                   // SAMMELRUF bedeutet Sammelrufnummer
	SAMMELFAX                   // SAMMELFAX bedeutet Sammelfaxnummer
	ABTEILUNGRUF                // ABTEILUNGRUF bedeutet Rufnummer Abteilung
	ABTEILUNGFAX                // ABTEILUNGFAX bedeutet Faxnummer Abteilung
	RUF_DURCHWAHL               // RUF_DURCHWAHL bedeutet Rufnummer Durchwahl
	FAX_DURCHWAHL               // FAX_DURCHWAHL bedeutet Faxnummer Durchwahl
	MOBIL_NUMMER                // MOBIL_NUMMER bedeutet Nummer des mobilen Telefons
)
