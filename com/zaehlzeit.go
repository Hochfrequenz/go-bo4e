package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/ermittlungleistungsmaximum"
	"github.com/hochfrequenz/go-bo4e/enum/haeufigkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/uebermittelbarkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlzeitdefinitiontyp"
)

type Zaehlzeit struct {
	Code                       *string                                                  `json:"code,omitempty"`
	Haeufigkeit                *haeufigkeitzaehlzeit.HaeufigkeitZaehlzeit               `json:"haeufigkeit,omitempty"`
	Uebermittelbarkeit         *uebermittelbarkeitzaehlzeit.UebermittelbarkeitZaehlzeit `json:"uebermittelbarkeit,omitempty"`
	ErmittlungLeistungsmaximum *ermittlungleistungsmaximum.ErmittlungLeistungsmaximum   `json:"ermittlungLeistungsmaximum,omitempty"`
	IstBestellbar              *bool                                                    `json:"istBestellbar,omitempty"`
	Typ                        *zaehlzeitdefinitiontyp.ZaehlzeitdefinitionTyp           `json:"typ,omitempty"`
	BeschreibungTyp            *string                                                  `json:"beschreibungTyp,omitempty"`
}
