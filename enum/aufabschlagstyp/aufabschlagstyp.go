package aufabschlagstyp

// AufAbschlagstyp defines , if the Auf- or Abschlag is relative or absolute
//
//go:generate stringer --type AufAbschlagstyp
//go:generate jsonenums --type AufAbschlagstyp
type AufAbschlagstyp int

const (
	RELATIV AufAbschlagstyp = iota + 1 // prozentualer AufAbschlag
	ABSOLUT                            //Absoluter AufAbschlag
)
