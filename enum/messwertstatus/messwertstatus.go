package messwertstatus

//go:generate stringer --type Messwertstatus
//go:generate jsonenums --type Messwertstatus
type Messwertstatus int // Der Status eines Zählerstandes

const (
	ABGELESEN            Messwertstatus = iota // Abgelesener Wert (abrechnungsrelevant)
	ERSATZWERT                                 // Ersatzwert - geschätzt, veranschlagt (abrechnungsrelevant)
	VORSCHLAGSWERT                             // Vorschlagswert (nicht abrechnungsrelevant)
	PROGNOSEWERT                               // Ein prognostizierter Wert
	VORLAUFIGERWERT                            // Ein vorläufiger Wert, dieser kann später angepasst werden.
	ENERGIEMENGESUMMIERT                       // Summenwert, Bilanzsumme
	FEHLT                                      // Explizite Kennzeichnung eines fehlendes Wertes
)
