package com

type Erreichbarkeit struct {
	MontagErreichbarkeit     *Zeitfenster `json:"montagErreichbarkeit,omitempty"`
	DienstagErreichbarkeit   *Zeitfenster `json:"dienstagErreichbarkeit,omitempty"`
	MittwochErreichbarkeit   *Zeitfenster `json:"mittwochErreichbarkeit,omitempty"`
	DonnerstagErreichbarkeit *Zeitfenster `json:"donnerstagErreichbarkeit,omitempty"`
	FreitagErreichbarkeit    *Zeitfenster `json:"freitagErreichbarkeit,omitempty"`
	Mittagspause             *Zeitfenster `json:"mittagspause,omitempty"`
}
