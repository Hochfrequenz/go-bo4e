package artikelid

// ArtikelId is a readable representation of the BNetzA article Ids which will replace BDEWArtikelnummer consecutively
//go:generate stringer --type ArtikelId
//go:generate jsonenums --type ArtikelId
type ArtikelId int

const (
	UNTERBRECHUNG_REGULAER          ArtikelId = iota + 1 // 2-01-7-001	Unterbrechung der Anschlussnutzung in der regulären Arbeitszeit (€/Auftrag)
	WIEDERHERSTELLUNG_REGULAER                           // 2-01-7-002	Wiederherstellung der Anschlussnutzung in der regulären Arbeitszeit (€/Auftrag)
	UNTERBRECHUNG_OHNE_ERFOLG                            // 2-01-7-003	Erfolglose Unterbrechung (€/Auftrag)
	STORNO_UNTERBRECHUNG_BIS_VORTAG                      // 2-01-7-004	Stornierung eines Auftrages zur Unterbrechung der Anschlussnutzung bis zum Vortag der Sperrung (€/Auftrag)
	STORNO_UNTERBRECHUNG_SPERRTAG                        // 2-01-7-005	Stornierung eines Auftrages zur Unterbrechung der Anschlussnutzung am Tag der Sperrung (€/Auftrag)
	WIEDERHERSTELLUNG_IRREGULAER                         // 2-01-7-006	Wiederherstellung der Anschlussnutzung außerhalb der regulären Arbeitszeit (€/Auftrag)
	VERZUGPAUSCHAL                                       // 2-02-0-001	Verzugskosten pauschal (€/Fall)
	VERZUGVARIABEL                                       // 2-02-0-002	Verzugskosten variabel (€)
)
