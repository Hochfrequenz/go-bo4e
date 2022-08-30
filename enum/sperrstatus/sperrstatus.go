package sperrstatus

//go:generate stringer --type Sperrstatus
//go:generate jsonenums --type Sperrstatus

// Der Sperrstatus beschreibt, ob ein Zähler gesperrt ist oder nicht
type Sperrstatus int

const (
	ENTSPERRT Sperrstatus = iota + 1 // ENTSPERRT heißt, dass der Zähler nicht gesperrt ist
	GESPERRT                         // GESPERRT beschreibt einen gesperrten Zähler
)
