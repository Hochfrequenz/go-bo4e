package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
	"time"
)

// Messlokationszuordnung. Mit dieser Komponente werden Messlokationen zu Marktlokationen zugeordnet. Dabei kann eine arithmetische Operation (Addition, Subtraktion) angegeben werden, mit der die Messlokation zum Verbrauch der Marklokation beiträgt.
type Messlokationszuordnung struct {
	MesslokationsId *string                                        `json:"messlokationsId,omitempty" validate:"alphanum,len=33"` // MesslokationsId, früher die Zählpunktbezeichnung.
	Arithmetik      *arithmetischeoperation.ArithmetischeOperation `json:"arithmetik,omitempty"`                                 // Arithmetik ist die Operation, mit der eine Messung an dieser Lokation für den Gesamtverbrauch der Marktlokation verrechnet wird. Beispielsweise bei einer Untermessung, wird der Verbauch der Untermessung subtrahiert.
	GueltigSeit     *time.Time                                     `json:"gueltigSeit,omitempty"`                                // GueltigSeit ist der inklusive Zeitpunkt, ab dem die Messlokation zur Marktlokation gehört
	GueltigBis      *time.Time                                     `json:"gueltigBis,omitempty"`                                 // GueltigBis ist der exklusive Zeitpunkt, bis zu dem die Messlokation zur Marktlokation gehört
	// dropped the GueltigSeit/GueltigBis lte/gtefield validation when they became optional
	// https://github.com/go-playground/validator/issues/972
}
