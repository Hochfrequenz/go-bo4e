package reklamationsgrund

// Reklamationsgrund gibt den Grund einer bo.Energiemenge n- bo.Reklamation
//go:generate stringer --type Reklamationsgrund
//go:generate jsonenums --type Reklamationsgrund
type Reklamationsgrund int

const (
	// WERTE_ZU_HOCH beschreibt, dass die bo.Energiemenge ist größer als die Erwartung
	WERTE_ZU_HOCH    Reklamationsgrund = iota + 1
	WERTE_ZU_NIEDRIG                   // WERTE_ZU_NIEDRIG beschreibt, dass die bo.Energiemenge ist kleiner als die Erwartung
	WERTE_FEHLEN                       // WERTE_FEHLEN beschreibt, dass die erwartete bo.Energiemenge fehlt
)
