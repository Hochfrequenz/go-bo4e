package technischeressourceverbrauchsart

// Verbrauchsart der Technischen Ressource
//
//go:generate stringer --type TechnischeRessourceVerbrauchsart
//go:generate jsonenums --type TechnischeRessourceVerbrauchsart

type TechnischeRessourceVerbrauchsart int

const (
	KRAFT_LICHT         TechnischeRessourceVerbrauchsart = iota + 1 //Z64: Kraft/Licht
	WAERME                                                          //Z65: Wärme
	E_MOBILITAET                                                    //ZE5: E-Mobilität
	STRASSENBELEUCHTUNG                                             //ZA8: Straßenbeleuchtung
)
