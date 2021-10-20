package messwertstatus

// The Messwertstatus is the status of a meter reading
//go:generate stringer --type Messwertstatus
//go:generate jsonenums --type Messwertstatus
type Messwertstatus int

const (
	// ABGELESEN ist ein Abgelesener Wert (abrechnungsrelevant)
	ABGELESEN            Messwertstatus = iota + 1
	ERSATZWERT                          // ERSATZWERT ist ein Ersatzwert - geschätzt, veranschlagt (abrechnungsrelevant)
	VORSCHLAGSWERT                      // VORSCHLAGSWERT ist ein Vorschlagswert (nicht abrechnungsrelevant)
	PROGNOSEWERT                        // PROGNOSEWERT ist ein prognostizierter Wert
	VORLAUFIGERWERT                     // VORLAUFIGERWERT ist ein vorläufiger Wert, dieser kann später angepasst werden.
	ENERGIEMENGESUMMIERT                // ENERGIEMENGESUMMIERT ist ein Summenwert, Bilanzsumme
	FEHLT                               // FEHLT ist eine explizite Kennzeichnung eines fehlendn Wertes
)
