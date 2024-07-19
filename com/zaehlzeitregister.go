package com

import "github.com/hochfrequenz/go-bo4e/enum/schwachlastfaehigkeit"

type Zaehlzeitregister struct {
	ZaehlzeitDefinition *string                                      `json:"zaehlzeitdefinition,omitempty"`
	Register            *string                                      `json:"register,omitempty"`
	Schwachlastfaehig   *schwachlastfaehigkeit.Schwachlastfaehigkeit `json:"schwachlastfaehig,omitempty"`
}
