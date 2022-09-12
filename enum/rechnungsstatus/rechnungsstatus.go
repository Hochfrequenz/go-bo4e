package rechnungsstatus

// Rechnungsstatus ist eine Abbildung verschiedener Zustände, die im Rahmen der Rechnungsbearbeitung durchlaufen werden
//
//go:generate stringer --type Rechnungsstatus
//go:generate jsonenums --type Rechnungsstatus
type Rechnungsstatus int

const (
	// UNGEPRUEFT means the invoice has not been checked yet / Die Rechnung wurde erstellt bzw ist eingegangen, wurde aber noch nicht geprüft.
	UNGEPRUEFT Rechnungsstatus = iota + 1
	// GEPRUEFT_OK means the invoice has been successfully checked / Die Rechnung wurde geprüft und als korrekt befunden.
	GEPRUEFT_OK
	// GEPRUEFT_FEHLERHAFT means the invoice check returned errors / Bei der Prüfung der Rechnung hat sich herausgestellt, dass die Rechnung Fehler aufweist.
	GEPRUEFT_FEHLERHAFT
	// GEBUCHT means the invoice has been booked / Die Buchhaltung hat die Rechnung aufgenommen und berücksichtigt.
	GEBUCHT
	// BEZAHLT means the invoice has been paid / Die Rechnung wurde beglichen.
	BEZAHLT
)
