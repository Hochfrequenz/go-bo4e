package bezeichnungsummenzeitreihe

//go:generate stringer --type BezeichnungSummenzeitreihe
//go:generate jsonenums --type BezeichnungSummenzeitreihe
type BezeichnungSummenzeitreihe int

const (
	BG_SZR_B BezeichnungSummenzeitreihe = iota + 1
	BG_SZR_C
	BK_SZR_A
	BK_SZR_B_RZ
	BK_SZR_B_BG
	BK_SZR_C
	LF_SZR_A
	LF_SZR_B_RZ
	LF_SZR_B_BG
	DZUE
	NZR
	ASZR
	NGZ
)
