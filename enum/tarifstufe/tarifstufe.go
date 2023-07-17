package tarifstufe

// Tarifstufe entspricht dem angegebenen Tarif aus der OBIS-Kennzahl mit der beanstandeten Energiemenge
//
//go:generate stringer --type Tarifstufe
//go:generate jsonenums --type Tarifstufe
type Tarifstufe int

const (
	TARIFSTUFE_0 Tarifstufe = iota + 20 // Z20
	TARIFSTUFE_1                        // Z21
	TARIFSTUFE_2                        // Z22
	TARIFSTUFE_3                        // Z23
	TARIFSTUFE_4                        // Z24
	TARIFSTUFE_5                        // Z25
	TARIFSTUFE_6                        // Z26
	TARIFSTUFE_7                        // Z27
	TARIFSTUFE_8                        // Z28
	TARIFSTUFE_9                        // Z29
)
