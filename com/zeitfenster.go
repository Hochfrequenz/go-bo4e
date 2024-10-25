package com

import (
	"github.com/hochfrequenz/go-bo4e/internal"
)

type Zeitfenster struct {
	Startzeit *internal.TimeOnly `json:"startzeit,omitempty"  example:"08:00:00"`
	Endzeit   *internal.TimeOnly `json:"endzeit,omitempty"  example:"18:00:00"`
}
