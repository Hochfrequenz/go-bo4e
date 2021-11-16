package com

import "github.com/hochfrequenz/go-bo4e/enum/themengebiet"

// Zustaendigkeit ordnet einem bo.Ansprechpartner Abteilungen und Zuständigkeiten zu
type Zustaendigkeit struct {
	Jobtitel     string                    `json:"jobtitel,omitempty"`                         // Jobtitel ist die berufliche Rolle des Ansprechpartners
	Abteilung    string                    `json:"abteilung,omitempty"`                        // Abteilung ist die Abteilung, in der der bo.Ansprechpartner tätig ist
	Themengebiet themengebiet.Themengebiet `json:"themengebiet,omitempty" validate:"required"` // Themengebiet erlaubt eine thematische Zuordnung des bo.Ansprechpartner
}
