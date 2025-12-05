package auftragsstornogrund

// Auftragsstornogrund Auftragsstornogrund für Sperr- bzw. Entsperrauftragsstornoantworten for 19128 and 19129
//
//go:generate stringer --type Auftragsstornogrund
//go:generate jsonenums --type Auftragsstornogrund
type Auftragsstornogrund int

const (
	STORNIERUNG_DER_ENTSPERRUNG_ERFOLGT Auftragsstornogrund = iota + 1
	STORNIERUNG_DER_SPERRUNG_ERFOLGT
)
