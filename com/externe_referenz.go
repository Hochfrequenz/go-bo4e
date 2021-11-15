package com

// ExterneReferenz models the existence of an object in another system and defines a way to reference it there
type ExterneReferenz struct {
	ExRefName string `json:"exRefName" validate:"required"` // ExRefName is the name of the external reference (e.g. "SAP IS-U" or "BelVis ID"
	ExRefWert string `json:"exRefWert" validate:"required"` // ExRefWert is the value of the external reference, usually a unique value
}
