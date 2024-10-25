package com

type Bankverbindung struct {
	IBAN         string `json:"iban,omitempty"`
	Kontoinhaber string `json:"kontoinhaber,omitempty"`
	Bankkennung  string `json:"bankkennung,omitempty"`
	Bankname     string `json:"bankname,omitempty"`
}
