package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
	"github.com/hochfrequenz/go-bo4e/enum/energieflussrichtung"
	"github.com/shopspring/decimal"
)

// Rechenschritt ist Teil einer Berechnungsformel.
// Entspricht der SG9 in UTILTS.
type Rechenschritt struct {
	// RechenschrittBestandteilId ist die BestandteilID des Rechenschritts (1-99999).
	RechenschrittBestandteilId int `json:"rechenschrittBestandteilId,omitempty"`

	// ReferenzRechenschrittId ist die ReferenzID des Rechenschritts (1-99999).
	ReferenzRechenschrittId int `json:"referenzRechenschrittID,omitempty"`

	// Operation ist die Rechenoperation dieses Schrittes.
	Operation arithmetischeoperation.ArithmetischeOperation `json:"operation,omitempty"`

	// VerlustfaktorTrafo ist ein möglicher Trafoverlust.
	VerlustfaktorTrafo *decimal.Decimal `json:"verlustfaktorTrafo,omitempty"`

	// VerlustfaktorLeitung ist ein möglicher Leitungsverlust.
	VerlustfaktorLeitung *decimal.Decimal `json:"verlustfaktorLeitung,omitempty"`

	// MesslokationId ist der Verweis auf eine MesslokationsId.
	MesslokationId *string `json:"messlokationId,omitempty"`

	// Energieflussrichtung gibt an, ob die gemessene Energie an der Messlokation zum Netz fließt (Erzeugung)
	// oder vom Netz wegfließt (Verbrauch).
	Energieflussrichtung *energieflussrichtung.Energieflussrichtung `json:"energieflussrichtung,omitempty"`

	// WeitererRechenschritt ermöglicht rekursive Verschachtelung weiterer Rechenschritte.
	WeitererRechenschritt *Rechenschritt `json:"weitererRechenschritt,omitempty"`

	// AufteilungsfaktorEnergiemenge ist der Aufteilungsfaktor für die Energiemenge.
	AufteilungsfaktorEnergiemenge *decimal.Decimal `json:"aufteilungsfaktorEnergiemenge,omitempty"`
}
