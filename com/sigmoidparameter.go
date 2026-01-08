package com

import "github.com/shopspring/decimal"

// Sigmoidparameter contains parameters for sigmoid pricing functions
type Sigmoidparameter struct {
	// A ist die Briefmarke Ortsverteilnetz.
	A decimal.Decimal `json:"A" validate:"required"`

	// B ist der Wendepunkt für die bepreiste Menge.
	B decimal.Decimal `json:"B" validate:"required"`

	// C ist der Exponent.
	C decimal.Decimal `json:"C" validate:"required"`

	// D ist die Briefmarke Transportnetz.
	D decimal.Decimal `json:"D" validate:"required"`
}
