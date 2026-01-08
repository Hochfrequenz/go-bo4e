package bo

// AuftragsStorno beschreibt die Stornierung eines Auftrags.
// Deprecated: In der C#-Implementierung wird stattdessen das enum Auftragsstornogrund verwendet.
// In C# ist dies eine abstrakte Klasse - in Go kann diese Struktur eingebettet werden.
type AuftragsStorno struct {
	Geschaeftsobjekt
	// AuftragsId ist die eindeutige Kennung des zu stornierenden Auftrags.
	AuftragsId string `json:"auftragsId,omitempty"`
}
