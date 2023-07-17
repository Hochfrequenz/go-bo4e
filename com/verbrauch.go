package com

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/tarifstufe"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
)

// Verbrauch is a a consumption during a specific time frame
type Verbrauch struct {
	// Startdatum is an inclusive start
	Startdatum time.Time `json:"startdatum,omitempty" validate:"omitempty"`

	// Enddatum is an exclusive end
	Enddatum time.Time `json:"enddatum,omitempty" validate:"omitempty,gtfield=Startdatum"`

	// Wertermittlungsverfahren is the type of the consumption
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren,omitempty" validate:"required"`

	// Messwertstatus includes the plausibility of the value
	Messwertstatus *messwertstatus.Messwertstatus `json:"messwertstatus,omitempty" validate:"omitempty"`

	// StatusZusatzInformationen enthält die Auflistung der STS Segmente
	// Plausibilisierungshinweis, Ersatzwertbildungsverfahren, Korrekturgrund,
	// Gasqualität, Tarif, Grundlage der Energiemenge.
	StatusZusatzInformationen []StatusZusatzInformation `json:"statuszusatzinformationen" validate:"omitempty"`

	// Obiskennzahl is the OBIS Kennzahl
	// (?<power>(1)-((?:[0-5]?[0-9])|(?:6[0-5])):((?:[1-8]|99))\.((?:6|8|9|29))\.([0-9]{1,2}))||(?<gas>)(7)-((?:[0-5]?[0-9])|(?:6[0-5])):(.{1,2})\.(.{1,2})\.([0-9]{1,2})
	Obiskennzahl string `json:"obiskennzahl,omitempty" validate:"required" example:"1-0:1.8.1"`

	// Wert is the value (of einheit)
	Wert decimal.Decimal `json:"wert,omitempty" validate:"required" example:"17"`

	// Einheit is the unit (associated to the wert)
	Einheit mengeneinheit.Mengeneinheit `json:"einheit,omitempty" validate:"required" example:"KWH"`

	Nutzungszeitpunkt     time.Time `json:"nutzungszeitpunkt"`     // Der Nutzungszeitpunkt wird verwendet, um einen Zählerstand eindeutig einem Prozesszeitpunkt zuzuordnen. Dieser Prozesszeitpunkt kann entweder ein Zeitpunkt einer Stammdatenänderung sein(z. B.bei einem Gerätewechsel, in der die Änderung vor dem Versand des Zählerstandes übermittelt wurde) oder die Bestellung eines Wertes aufgrund eines eingetretenen Ereignisses(z.B. Lieferantenwechsel). Der  Nutzungszeitpunkt ist für den Zählerstand der Zeitpunkt der für die weitere Verarbeitung relevant ist(z.B.Zuordnung bei Empfänger anhand der Zuordnungstupel)
	Ausfuehrungszeitpunkt time.Time `json:"ausfuehrungszeitpunkt"` // Der Ausfuehrungszeitpunkt wird verwendet, um einen Zählerstand eindeutig einer tatsächlichen Änderung zuzuordnen, z.B.bei einem Gerätewechsel oder Geräteparameteränderung der tatsächliche Zeitpunkt an dem die Änderung an der Messlokation durchgeführt wurde. Der Nutzungszeitpunkt ist für den Zählerstand der Zeitpunkt der für die weitere Verarbeitung relevant ist(z.B. Zuordnung bei Empfänger anhand der Zuordnungstupel).

	Tarifstufe tarifstufe.Tarifstufe `json:"tarifstufe,omitempty" validate:"omitempty" example:"TARIFSTUFE_0"`
}
