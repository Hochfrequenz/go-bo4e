package fehlertyp

// Typen von com.Fehler n
//
//go:generate stringer --type FehlerTyp
//go:generate jsonenums --type FehlerTyp
type FehlerTyp int

const (
	SYNTAX       FehlerTyp = iota + 1 // Modellfehler / Syntaxfehler
	VERARBEITUNG                      // Fehler in der Verarbeitung
)
