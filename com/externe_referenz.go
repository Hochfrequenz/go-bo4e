package com

// ExterneReferenz models the existence of an object in another system and defines a way to reference it there
type ExterneReferenz struct {
	ExRefName string `json:"exRefName" validate:"required"` // Name of the external reference (e.g. "SAP IS-U" or "BelVis ID"
	ExRefWert string `json:"exRefWert" validate:"required"` // Value of the external reference, usually a unique value
}
