package com

import (
	"time"
)

type AusgerollteZaehlzeit struct {
	Aenderungszeitpunkt time.Time `json:"aenderungszeitpunkt"`
	Register            *string   `json:"register"`
}
