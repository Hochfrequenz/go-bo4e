package grundlageenergiemenge

// GrundlageEnergiemenge gives hints about the foundation of an energy amount
//go:generate stringer --type GrundlageEnergiemenge
//go:generate jsonenums --type  GrundlageEnergiemenge
type GrundlageEnergiemenge int

const (
	// ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT
	// corresponds to edi@energy qualifier Z36 "Zählerstand zum Beginn der angegebenen Energiemenge
	// vorhanden und kommuniziert"
	ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT GrundlageEnergiemenge = iota + 1

	// ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT
	// corresponds to edi@energy qualifier Z37 "Zählerstand zum Ende der angegebenen Energiemenge
	// vorhanden und kommuniziert
	ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT

	// ZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG
	// corresponds to edi@energy qualifier Z38 "Zählerstand zum Beginn der angegebenen Energiemenge
	// nicht vorhanden da Mengenabgrenzung"
	ZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG

	// ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG
	// corresponds to edi@energy qualifier Z39 "Zählerstand zum Ende der angegebenen Energiemenge
	// nicht vorhanden da Mengenabgrenzung"
	ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG
)
