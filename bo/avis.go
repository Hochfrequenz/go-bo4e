package bo

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/avistyp"
	"github.com/shopspring/decimal"
)

type Avis struct {
	Geschaeftsobjekt
	AvisNummer     string             `json:"avisNummer,omitempty" validate:"required"`       // Eine im Verwendungskontext eindeutige Nummer für das Avis
	AvisTyp        avistyp.AvisTyp    `json:"avisTyp,omitempty" validate:"required"`          // Gibt den Typ des Avis an
	AvisPositionen []com.AvisPosition `json:"positionen,omitempty" validate:"required,min=1"` // Avispositionen
	ZuZahlen       com.Betrag         `json:"zuZahlen,omitempty" validate:"required"`         // Summenbetrag
}

// AvisStructLevelValidation combines all the single validators
func AvisStructLevelValidation(sl validator.StructLevel) {
	AvisStructLevelValidationZuZahlen(sl)
	AvisStructLevelValidationZuZahlenAbgelehnteForderung(sl)
	AvisStructLevelValidationAbweichung(sl)
}

// AvisStructLevelValidationGesamtNetto verifies that the sum of all Rechnungsposition.Netto equals the Rechnung.GesamtNetto
func AvisStructLevelValidationZuZahlen(sl validator.StructLevel) {
	avis := sl.Current().Interface().(Avis)
	expectedZuZahlen := decimal.Zero
	for _, avisPosition := range avis.AvisPositionen {
		if avisPosition.ZuZahlen.Waehrung != avis.ZuZahlen.Waehrung {
			// the waehrung has to be consistent over all avispositionen
			sl.ReportError(avisPosition.ZuZahlen.Waehrung, "Waehrung", "AvisPositionen", "Avis.ZuZahlen.Waehrung==Avis.AvisPositionen.ZuZahlen[j].Waehrung", "")
			return
		}
		expectedZuZahlen = expectedZuZahlen.Add(avis.ZuZahlen.Wert)
	}
	if avis.ZuZahlen.Wert != expectedZuZahlen {
		// the waehrung has to be consistent over all avispositionen
		sl.ReportError(avis.ZuZahlen, "Wert", "ZuZahlen", "Avis.ZuZahlen.Wert==sum(Avis.AvisPositionen.ZuZahlen[].Wert)", "")
		return
	}
}

func AvisStructLevelValidationZuZahlenAbgelehnteForderung(sl validator.StructLevel) {
	avis := sl.Current().Interface().(Avis)
	// nothing to check for ZAHLUNGSAVIS as any value is valid there
	if avis.AvisTyp == avistyp.ABGELEHNTE_FORDERUNG {
		if avis.ZuZahlen.Wert != decimal.Zero {
			// ABGELEHNTE_FORDERUNG must have zero as the amount to pay
			sl.ReportError(avis.ZuZahlen, "Wert", "ZuZahlen", "Avis.ZuZahlen.Wert==0", "")
			return
		}
		for _, avisPosition := range avis.AvisPositionen {
			if avisPosition.ZuZahlen.Wert != decimal.Zero {
				// ABGELEHNTE_FORDERUNG must have zero as the amount to pay
				sl.ReportError(avisPosition.ZuZahlen, "Wert", "ZuZahlen", "Avis.AvisPositionen.ZuZahlen.Wert==0", "")
				return
			}
		}
	}
}

func AvisStructLevelValidationAbweichung(sl validator.StructLevel) {
	avis := sl.Current().Interface().(Avis)
	declineAvis := avis.AvisTyp == avistyp.ABGELEHNTE_FORDERUNG
	for _, avisPosition := range avis.AvisPositionen {
		if declineAvis && (len(avisPosition.Abweichungen) == 0 || avisPosition.Abweichungen[0].AbweichungsGrund == nil || *avisPosition.Abweichungen[0].AbweichungsGrund == 0) {
			sl.ReportError(avisPosition.Abweichungen, "AbweichungsGrund", "Abweichung", "Avis.AvisPositionen.Abweichung->AbgelehnteForderungRequired", "")
			return
		}
		if len(avisPosition.Abweichungen) != 0 && !declineAvis {
			sl.ReportError(avisPosition.Abweichungen, "AbweichungsGrund", "Abweichung", "Avis.AvisPositionen.Abweichung->ZahlungsavisRequired", "")
			return
		}
	}
}
