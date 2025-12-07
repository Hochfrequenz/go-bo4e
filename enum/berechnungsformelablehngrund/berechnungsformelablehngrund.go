package berechnungsformelablehngrund

//go:generate stringer --type BerechnungsformelAblehngrund
//go:generate jsonenums --type BerechnungsformelAblehngrund
type BerechnungsformelAblehngrund int

const (
	START_DATUM_UNPLAUSIBEL                               BerechnungsformelAblehngrund = iota + 1 // A02: Gueltig-Ab Datum der Berechnungsformel ist unplausibel
	ZU_VIELE_MESSLOKATIONEN                                                                       // A04: Es sind zu viele Messlokationen in der Berechnungformel vorhanden
	FEHLENDE_MESSLOKATIONEN                                                                       // A05: Es fehlen Messlokationen in der Berechnungsformel
	FALSCHE_MESSLOKATION                                                                          // A06: ID der Messlokationen stimmen nicht ueberein
	FALSCHE_LIEFERRICHTUNG_DER_MARKTLOKATION                                                      // A01: Lieferrichtung der Marktlokation ist nicht korrekt
	FALSCHE_FLUSSRICHTUNG_DER_MESSLOKATION                                                        // A07: Die Flussrichtung mindestens einer Messlokation ist nicht korrekt angegeben
	MARKTLOKATION_IST_NICHT_GENAU_DER_MESSLOKATION_ZUGEORDNET                                     // A10: Der Marktlokation ist nicht genau eine Messlokation zugeordnet
)
