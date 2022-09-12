package handelsunstimmigkeitstyp

// Handelsunstimmigkeitstyp are the possible types of a bo.Handelsunstimmigkeit
//
//go:generate stringer --type Handelsunstimmigkeitstyp
//go:generate jsonenums --type Handelsunstimmigkeitstyp
type Handelsunstimmigkeitstyp int

const (
	// HANDELSRECHNUNG represents SG2 DOC 1001 code "380"
	HANDELSRECHNUNG Handelsunstimmigkeitstyp = iota + 1

	// LIEFERSCHEIN_HANDELSUNSTIMMIGKEITSTYP represents SG2 DOC 1001 code "270"
	LIEFERSCHEIN_HANDELSUNSTIMMIGKEITSTYP

	// LIEFERSCHEIN_GRUND_ARBEITSPREIS represents SG2 DOC 1001 code "Z41"
	LIEFERSCHEIN_GRUND_ARBEITSPREIS

	// LIEFERSCHEIN_ARBEITS_LEISTUNGSPREIS represents SG2 DOC 1001 code "Z42"
	LIEFERSCHEIN_ARBEITS_LEISTUNGSPREIS
)
