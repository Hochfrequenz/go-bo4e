package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/emobilitaetsart"
	"github.com/hochfrequenz/go-bo4e/enum/unterbrechbarkeit"
	"github.com/hochfrequenz/go-bo4e/enum/verbrauchsart"
	"github.com/hochfrequenz/go-bo4e/enum/waermenutzung"
)

type Messprodukt struct {
	Emobilitaetsart        *emobilitaetsart.EMobilitaetsart     `json:"emobilitaetsart,omitempty"`
	MessproduktID          *string                              `json:"messproduktId,omitempty"`
	Unterbrechbarkeit      *unterbrechbarkeit.Unterbrechbarkeit `json:"unterbrechbarkeit,omitempty"`
	Verbrauchsart          *verbrauchsart.Verbrauchsart         `json:"verbrauchsart,omitempty"`
	Verwendungszwecke      []Verwendungszweck                   `json:"verwendungszwecke,omitempty"`
	Waermenutzung          *waermenutzung.Waermenutzung         `json:"waermenutzung,omitempty"`
	WerteuebermittlungAnNB *bool                                `json:"werteuebermittlungAnNB,omitempty"`
	Zaehlzeiten            *Zaehlzeitregister                   `json:"zaehlzeiten,omitempty"`
	ZweiteMessung          *bool                                `json:"zweiteMessung,omitempty"`
}
