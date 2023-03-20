package waermenutzung

// Stromverbrauchsart/Wärmenutzung Marktlokation

//go:generate stringer --type Waermenutzung
//go:generate jsonenums --type Waermenutzung
type Waermenutzung int

const (
	SPEICHERHEIZUNG Waermenutzung = iota + 1
	WAERMEPUMPE
	DIREKTHEIZUNG
)
