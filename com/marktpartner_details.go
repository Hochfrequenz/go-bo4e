package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
)

// MarktpartnerDetails represents market partner details (e.g. Marktrolle)
type MarktpartnerDetails struct {
	Rollencodenummer   *string                `json:"rollencodenummer,omitempty"`
	Code               *string                `json:"code,omitempty"`
	Marktrolle         *marktrolle.Marktrolle `json:"marktrolle,omitempty"`
	Weiterverpflichtet *bool                  `json:"weiterverpflichtet,omitempty"`
}
