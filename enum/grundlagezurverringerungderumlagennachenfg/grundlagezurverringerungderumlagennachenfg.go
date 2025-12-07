package grundlagezurverringerungderumlagennachenfg

//go:generate stringer --type GrundlageZurVerringerungDerUmlagenNachEnfg
//go:generate jsonenums --type GrundlageZurVerringerungDerUmlagenNachEnfg
type GrundlageZurVerringerungDerUmlagenNachEnfg int

const (
	KUNDE_ERFUELLT_VORAUSSETZUNG GrundlageZurVerringerungDerUmlagenNachEnfg = iota + 1
	KUNDE_ERFUELLT_VORAUSSETZUNG_NICHT
	KEINE_ANGABE
)
