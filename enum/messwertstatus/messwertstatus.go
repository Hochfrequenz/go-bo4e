package messwertstatus

// The Messwertstatus is the status of a meter reading
//go:generate stringer --type Messwertstatus
//go:generate jsonenums --type Messwertstatus
type Messwertstatus int

const (
	ABGELESEN            Messwertstatus = iota + 1 // Abgelesener Wert (abrechnungsrelevant)
	ERSATZWERT                                     // Ersatzwert - geschätzt, veranschlagt (abrechnungsrelevant)
	VORSCHLAGSWERT                                 // Vorschlagswert (nicht abrechnungsrelevant)
	PROGNOSEWERT                                   // Ein prognostizierter Wert
	VORLAUFIGERWERT                                // Ein vorläufiger Wert, dieser kann später angepasst werden.
	ENERGIEMENGESUMMIERT                           // Summenwert, Bilanzsumme
	FEHLT                                          // Explizite Kennzeichnung eines fehlendes Wertes
)
