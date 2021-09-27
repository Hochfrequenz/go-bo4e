package rechnungsstatus

// Rechnungsstatus ist eine Abbildung verschiedener Zustände, die im Rahmen der Rechnungsbearbeitung durchlaufen werden
//go:generate stringer --type Rechnungsstatus
//go:generate jsonenums --type Rechnungsstatus
type Rechnungsstatus int

const (
	// Ungeprueft means the invoice has not been checked yet
	Ungeprueft Rechnungsstatus = iota + 1
	// GeprueftOk means the invoice has been successfully checked
	GeprueftOk
	// GeprueftFehlerhaft means the invoice check returned errors
	GeprueftFehlerhaft
	// Gebucht means the invoice has been booked
	Gebucht
	// Bezahlt means the invoice has been paid
	Bezahlt
)
