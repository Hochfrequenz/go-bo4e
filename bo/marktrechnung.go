package bo

// Marktrechnung ist eine Rechnung, die zwischen Marktteilnehmer n ausgetauscht wird.
type Marktrechnung struct {
	Rechnung
	Rechnungsersteller  Marktteilnehmer `json:"rechnungsersteller" validate:"required"`  // Der Aussteller der Rechnung
	Rechnungsempfaenger Marktteilnehmer `json:"rechnungsempfaenger" validate:"required"` // Der Empfänger der Rechnung
}
