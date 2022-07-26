package berichtstatus

// Gibt den Status eines Statusberichts an
//go:generate stringer --type BerichtStatus
//go:generate jsonenums --type BerichtStatus
type BerichtStatus int

const (
	ERFOLGREICH BerichtStatus = iota + 1 //
	FEHLER                               //
)
