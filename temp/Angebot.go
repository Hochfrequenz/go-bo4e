// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    angebot, err := UnmarshalAngebot(bytes)
//    bytes, err = angebot.Marshal()

package main

import "encoding/json"

func UnmarshalAngebot(data []byte) (Angebot, error) {
	var r Angebot
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Angebot) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Angebot struct {
	Description          string     `json:"description"`
	Title                string     `json:"title"`
	Type                 string     `json:"type"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Properties           Properties `json:"properties"`
}

type Properties struct {
	ID                          ID            `json:"_id"`
	Typ                         Typ           `json:"_typ"`
	Version                     ID            `json:"_version"`
	Anfragereferenz             ID            `json:"anfragereferenz"`
	Angebotsdatum               Angebotsdatum `json:"angebotsdatum"`
	Angebotsgeber               Typ           `json:"angebotsgeber"`
	Angebotsnehmer              Typ           `json:"angebotsnehmer"`
	Angebotsnummer              ID            `json:"angebotsnummer"`
	Bindefrist                  Angebotsdatum `json:"bindefrist"`
	Sparte                      Typ           `json:"sparte"`
	UnterzeichnerAngebotsgeber  Typ           `json:"unterzeichnerAngebotsgeber"`
	UnterzeichnerAngebotsnehmer Typ           `json:"unterzeichnerAngebotsnehmer"`
	Varianten                   Varianten     `json:"varianten"`
	ZusatzAttribute             Varianten     `json:"zusatzAttribute"`
}

type ID struct {
	Title   string    `json:"title"`
	Default *string   `json:"default"`
	AnyOf   []IDAnyOf `json:"anyOf"`
}

type IDAnyOf struct {
	Type string `json:"type"`
}

type Angebotsdatum struct {
	Title   string               `json:"title"`
	Default interface{}          `json:"default"`
	AnyOf   []AngebotsdatumAnyOf `json:"anyOf"`
}

type AngebotsdatumAnyOf struct {
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
}

type Typ struct {
	Default *string    `json:"default"`
	AnyOf   []TypAnyOf `json:"anyOf"`
}

type TypAnyOf struct {
	Ref  *string `json:"$ref,omitempty"`
	Type *string `json:"type,omitempty"`
}

type Varianten struct {
	Title   string           `json:"title"`
	Default interface{}      `json:"default"`
	AnyOf   []VariantenAnyOf `json:"anyOf"`
}

type VariantenAnyOf struct {
	Type  string `json:"type"`
	Items *Items `json:"items,omitempty"`
}

type Items struct {
	Ref string `json:"$ref"`
}
