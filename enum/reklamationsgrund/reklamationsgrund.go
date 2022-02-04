package reklamationsgrund

// Reklamationsgrund gibt den Grund einer bo.Energiemenge n- bo.Reklamation
//go:generate stringer --type Reklamationsgrund
//go:generate jsonenums --type Reklamationsgrund
type Reklamationsgrund int

const (
	// die erwartete bo.Energiemenge fehlt
	FEHLT      Reklamationsgrund = iota + 1
	ZU_NIEDRIG                   // die bo.Energiemenge ist kleiner als die Erwartung
	ZU_HOCH                      // die bo.Energiemenge ist größer als die Erwartung
)
