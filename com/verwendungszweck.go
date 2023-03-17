package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/verwendungszweck"
)

type Verwendungszweck struct {
	Marktrolle marktrolle.Marktrolle               `json:"marktrolle,omitempty"` // rollencodenummer von Marktrolle
	Zweck      []verwendungszweck.Verwendungszweck `json:"zweck,omitempty"`      // code von Marktrolle
}
