package bo

import (
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/shopspring/decimal"
)

type Tranche struct {
	Geschaeftsobjekt
	TrancheId        string              `json:"trancheId" validate:"required"`        // Identifikationsnummer einer Tranche, an der Energie entweder verbraucht, oder erzeugt wird (Like MarktlokationsId <see cref="Marktlokation"/>)
	Sparte           *sparte.Sparte      `json:"sparte" validate:"required"`           // Sparte der Tranche, z.B. Gas oder Strom. It is nullable because in the future it should be that way, although it is not nullable in BO4E-dotnet
	Aufteilungsmenge decimal.NullDecimal `json:"aufteilungsmenge" validate:"required"` // Prozentualer Anteil der Tranche an der erzeugenden Marktlokation in Prozent mit 2 Nachkommastellen
	ObisKennzahl     *string             `json:"obisKennzahl" validate:"required"`     // Die OBIS-Kennzahl für die Tranche, die festlegt, welche auf die gemessene Größe mit dem Stand gemeldet wird.
}
