package messtechnischeeinordnung

//go:generate stringer --type MesstechnischeEinordnung
//go:generate jsonenums --type MesstechnischeEinordnung
type MesstechnischeEinordnung int

const (
	IMS MesstechnischeEinordnung = iota + 1
	KME_MME
	KEINE_MESSUNG
)
