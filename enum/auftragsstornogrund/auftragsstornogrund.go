package auftragsstornogrund

// Auftragsstornogrund Auftragsstornogrund für Sperr- bzw. Entsperrauftragsstornoantworten for 19128 and 19129
//
//go:generate stringer --type Auftragsstornogrund
//go:generate jsonenums --type Auftragsstornogrund
type Auftragsstornogrund int

const (
	STORNIERUNG_DER_ENTSPERRUNG_ERFOLGT             Auftragsstornogrund = iota + 1 // A01
	STORNIERUNG_DER_SPERRUNG_ERFOLGT                                               // A04
	STORNIERUNG_DER_ENTSPERRUNG_NICHT_MEHR_MOEGLICH                                // A02
	STORNIERUNG_DER_SPERRUNG_NICHT_MEHR_MOEGLICH                                   // A05
	STORNIERUNG_DER_SPERRUNG_BIS_ZUM_VORTAG_ERFOLGT                                // A03
)
