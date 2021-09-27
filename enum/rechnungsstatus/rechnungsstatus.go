package rechnungsstatus

// Rechnungsstatus ist eine Abbildung verschiedener Zustände, die im Rahmen der Rechnungsbearbeitung durchlaufen werden
//go:generate stringer --type Rechnungsstatus
//go:generate jsonenums --type Rechnungsstatus
type Rechnungsstatus int

const (
	// Ungeprueft means the invoice has not been checked yet / Die Rechnung wurde erstellt bzw ist eingegangen, wurde aber noch nicht geprüft.
	Ungeprueft Rechnungsstatus = iota + 1
	// GeprueftOk means the invoice has been successfully checked / Die Rechnung wurde geprüft und als korrekt befunden.
	GeprueftOk
	// GeprueftFehlerhaft means the invoice check returned errors / Bei der Prüfung der Rechnung hat sich herausgestellt, dass die Rechnung Fehler aufweist.
	GeprueftFehlerhaft
	// Gebucht means the invoice has been booked / Die Buchhaltung hat die Rechnung aufgenommen und berücksichtigt.
	Gebucht
	// Bezahlt means the invoice has been paid / Die Rechnung wurde beglichen.
	Bezahlt
)
