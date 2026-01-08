package com

// GenericStringStringInfo ist ein generisches Objekt zum Abrufen von Werten aus einer Dummy Core Data Service View.
type GenericStringStringInfo struct {
	// KeyColumn ist der Schlüssel (anders benannt, da "key" ein reserviertes Schlüsselwort ist).
	KeyColumn *string `json:"keyColumn,omitempty"`

	// Value ist der Wert.
	Value *string `json:"value,omitempty"`
}
