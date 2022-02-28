package bo_test

import (
	"reflect"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/abweichungsgrund"
	"github.com/hochfrequenz/go-bo4e/enum/avistyp"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

// TestFailedAvisValidation verifies that the validators of Avis work
func (s *Suite) Test_Failed_AvisValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(bo.AvisStructLevelValidation, bo.Avis{})

	invalidAvisMap := map[string][]interface{}{
		"required": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				AvisNummer:     "",
				AvisTyp:        0,
				AvisPositionen: []com.AvisPosition{},
				ZuZahlen: com.Betrag{
					Wert:     decimal.Decimal{},
					Waehrung: waehrungscode.EUR,
				},
			},
		},
		"min": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.AVIS,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				AvisNummer:     "1",
				AvisTyp:        avistyp.ZAHLUNGSAVIS,
				AvisPositionen: []com.AvisPosition{}, // len 0
				ZuZahlen: com.Betrag{
					Wert:     decimal.Decimal{},
					Waehrung: waehrungscode.EUR,
				},
			},
		},
		"Avis.ZuZahlen.Wert==sum(Avis.AvisPositionen.ZuZahlen[].Wert)": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.AVIS,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				AvisNummer: "1",
				AvisTyp:    avistyp.ZAHLUNGSAVIS,
				AvisPositionen: []com.AvisPosition{
					{
						RechnungsNummer: "2",
						RechnungsDatum:  time.Now(),
						Storno:          false,
						GesamtBrutto: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.EUR,
						},
						ZuZahlen: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.EUR,
						},
						Abweichung: com.Abweichung{
							Referenz:         "A",
							AbweichungsGrund: abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
						},
					},
				},
				ZuZahlen: com.Betrag{
					Wert:     decimal.New(1, 0),
					Waehrung: waehrungscode.EUR,
				},
			},
		},
		"Avis.ZuZahlen.Waehrung==Avis.AvisPositionen.ZuZahlen[j].Waehrung": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.AVIS,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				AvisNummer: "1",
				AvisTyp:    avistyp.ZAHLUNGSAVIS,
				AvisPositionen: []com.AvisPosition{
					{
						RechnungsNummer: "2",
						RechnungsDatum:  time.Now(),
						Storno:          false,
						GesamtBrutto: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.EUR,
						},
						ZuZahlen: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.AFN,
						},
						Abweichung: com.Abweichung{
							Referenz:         "A",
							AbweichungsGrund: abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
						},
					},
				},
				ZuZahlen: com.Betrag{
					Wert:     decimal.New(1, 0),
					Waehrung: waehrungscode.EUR,
				},
			},
		},
		"Avis.ZuZahlen.Wert==0": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.AVIS,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				AvisNummer: "1",
				AvisTyp:    avistyp.ABGELEHNTE_FORDERUNG,
				AvisPositionen: []com.AvisPosition{
					{
						RechnungsNummer: "2",
						RechnungsDatum:  time.Now(),
						Storno:          false,
						GesamtBrutto: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.EUR,
						},
						ZuZahlen: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.AFN,
						},
						Abweichung: com.Abweichung{
							Referenz:         "A",
							AbweichungsGrund: abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
						},
					},
				},
				ZuZahlen: com.Betrag{
					Wert:     decimal.New(1, 0),
					Waehrung: waehrungscode.EUR,
				},
			},
		},
		"Avis.AvisPositionen.ZuZahlen.Wert==0": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.AVIS,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				AvisNummer: "1",
				AvisTyp:    avistyp.ABGELEHNTE_FORDERUNG,
				AvisPositionen: []com.AvisPosition{
					{
						RechnungsNummer: "2",
						RechnungsDatum:  time.Now(),
						Storno:          false,
						GesamtBrutto: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.EUR,
						},
						ZuZahlen: com.Betrag{
							Wert:     decimal.New(1, 0),
							Waehrung: waehrungscode.AFN,
						},
						Abweichung: com.Abweichung{
							Referenz:         "A",
							AbweichungsGrund: abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
						},
					},
				},
				ZuZahlen: com.Betrag{
					Wert:     decimal.Zero,
					Waehrung: waehrungscode.EUR,
				},
			},
		},
		"Avis.AvisPositionen.Abweichung->ZahlungsavisRequired": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.AVIS,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				AvisNummer: "1",
				AvisTyp:    avistyp.ZAHLUNGSAVIS,
				AvisPositionen: []com.AvisPosition{
					{
						RechnungsNummer: "2",
						RechnungsDatum:  time.Now(),
						Storno:          false,
						Abweichung: com.Abweichung{
							Referenz:                  "A",
							AbweichungsGrund:          abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
							AbweichungsGrundBemerkung: "B",
						},
						GesamtBrutto: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.EUR,
						},
						ZuZahlen: com.Betrag{
							Wert:     decimal.New(1, 0),
							Waehrung: waehrungscode.AFN,
						},
					},
				},
				ZuZahlen: com.Betrag{
					Wert:     decimal.New(1, 0),
					Waehrung: waehrungscode.EUR,
				},
			},
		},
		"Avis.AvisPositionen.Abweichung->AbgelehnteForderungRequired": {
			bo.Avis{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.AVIS,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				AvisNummer: "1",
				AvisTyp:    avistyp.ABGELEHNTE_FORDERUNG,
				AvisPositionen: []com.AvisPosition{
					{
						RechnungsNummer: "2",
						RechnungsDatum:  time.Now(),
						Storno:          false,
						GesamtBrutto: com.Betrag{
							Wert:     decimal.Zero,
							Waehrung: waehrungscode.EUR,
						},
						ZuZahlen: com.Betrag{
							Wert:     decimal.New(1, 0),
							Waehrung: waehrungscode.AFN,
						},
					},
				},
				ZuZahlen: com.Betrag{
					Wert:     decimal.New(1, 0),
					Waehrung: waehrungscode.EUR,
				},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidAvisMap)
}

//  TestSuccessfulMesslokationValidation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_AvisValidation() {
	validate := validator.New()
	avisPositive := com.AvisPosition{
		RechnungsNummer: "P",
		RechnungsDatum:  time.Now(),
		Storno:          false,
		GesamtBrutto: com.Betrag{
			Wert:     decimal.New(10, 0),
			Waehrung: waehrungscode.EUR,
		},
		ZuZahlen: com.Betrag{
			Wert:     decimal.New(10, 0),
			Waehrung: waehrungscode.EUR,
		},
	}
	avisNegative := com.AvisPosition{
		RechnungsNummer: "N",
		RechnungsDatum:  time.Now(),
		Storno:          false,
		GesamtBrutto: com.Betrag{
			Wert:     decimal.New(0, 0),
			Waehrung: waehrungscode.EUR,
		},
		ZuZahlen: com.Betrag{
			Wert:     decimal.New(15, 0),
			Waehrung: waehrungscode.EUR,
		},
		Abweichung: com.Abweichung{
			Referenz:                  "B",
			AbweichungsGrund:          abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
			AbweichungsGrundBemerkung: "C",
		},
	}
	validAvises := []bo.BusinessObject{
		bo.Avis{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.AVIS,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			AvisNummer:     "1",
			AvisTyp:        avistyp.ABGELEHNTE_FORDERUNG,
			AvisPositionen: []com.AvisPosition{avisNegative},
			ZuZahlen:       avisNegative.GesamtBrutto,
		},
		bo.Avis{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.AVIS,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			AvisNummer:     "2",
			AvisTyp:        avistyp.ZAHLUNGSAVIS,
			AvisPositionen: []com.AvisPosition{avisPositive},
			ZuZahlen:       avisPositive.GesamtBrutto,
		},
	}
	VerfiySuccessfulValidations(s, validate, validAvises)
}

func (s *Suite) Test_Empty_Avis_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.AVIS)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Avis{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.AVIS))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Avis_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.AVIS))
}
