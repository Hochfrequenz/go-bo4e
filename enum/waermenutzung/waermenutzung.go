package waermenutzung

// Stromverbrauchsart/Wärmenutzung Marktlokation

//go:generate stringer --type Waermenutzung
//go:generate jsonenums --type Waermenutzung
type Waermenutzung int

const (
	SPEICHERHEIZUNG Waermenutzung = iota + 1
	WAERMEPUMPE
	DIREKTHEIZUNG

	WAERMEPUMPE_WAERME_KAELTE // WAERMEPUMPE_WAERME_KAELTE is ZV5 Wärmepumpe (Wärme+Kälte)
	WAERMEPUMPE_KAELTE        // WAERMEPUMPE_KAELTE is ZV6 Wärmepumpe (Kälte)
	WAERMEPUMPE_WAERME        // WAERMEPUMPE_WAERME is ZV7 Wärmepumpe (Wärme)
)
