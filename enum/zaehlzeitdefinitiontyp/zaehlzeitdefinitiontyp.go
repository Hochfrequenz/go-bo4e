package zaehlzeitdefinitiontyp

//go:generate stringer --type ZaehlzeitdefinitionTyp
//go:generate jsonenums --type ZaehlzeitdefinitionTyp
type ZaehlzeitdefinitionTyp int

const (
	WAERMEPUMPE ZaehlzeitdefinitionTyp = iota + 1
	NACHTSPEICHERHEIZUNG
	SCHWACHLASTZEITFENSTER
	SONSTIGE
	HOCHLASTZEITFENSTER
)
