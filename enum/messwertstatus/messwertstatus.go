package messwertstatus

// The Messwertstatus is the status of a meter reading
//
//go:generate stringer --type Messwertstatus
//go:generate jsonenums --type Messwertstatus
type Messwertstatus int

const (
	// ABGELESEN corresponds "Wahrer Wert: 220"
	ABGELESEN Messwertstatus = iota + 1

	// ERSATZWERT corresponds to"Ersatzwert: 67"
	ERSATZWERT

	// VOLAEUFIGERWERT corresponds to "Vorläufiger Wert: Z18"
	VOLAEUFIGERWERT

	// ANGABE_FUER_LIEFERSCHEIN corresponds to "Angabe für Lieferschein: Z31"
	ANGABE_FUER_LIEFERSCHEIN

	// VORSCHLAGSWERT corresponds to "Vorschlagswert: 201"
	VORSCHLAGSWERT

	// NICHT_VERWENDBAR corresponds to "Nicht verwendbarer Wert: 20"
	NICHT_VERWENDBAR

	// PROGNOSEWERT corresponds to "Prognosewert: 187"
	PROGNOSEWERT

	// ENERGIEMENGESUMMIERT corresponds to "Energiemenge summiert: 79"
	ENERGIEMENGESUMMIERT

	// FEHLT corresponds to "Fehlender Wert: Z30"
	FEHLT

	// GRUNDLAGE_POG_ERMITTLUNG indicates that the value is used as basis for PoG determination
	GRUNDLAGE_POG_ERMITTLUNG
)
