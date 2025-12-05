package auftragsstornogrund

// Auftragsstornogrund is an enum
//
//go:generate stringer --type Auftragsstornogrund
//go:generate jsonenums --type Auftragsstornogrund
type Auftragsstornogrund int

const (
	STORNIERUNG_DER_ENTSPERRUNG_ERFOLGT Auftragsstornogrund = iota + 1
	STORNIERUNG_DER_SPERRUNG_ERFOLGT
)
