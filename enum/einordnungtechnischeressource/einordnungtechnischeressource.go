package einordnungtechnischeressource

// EinordnungTechnischeRessource Einordnung der verbrauchenden Technischen Ressource nach § 14a EnWG mit Inbetriebsetzung vor 2024
//
//go:generate stringer --type EinordnungTechnischeRessource
//go:generate jsonenums --type EinordnungTechnischeRessource
type EinordnungTechnischeRessource int

const (
	WECHSEL_IN_14A_EINMALIG_MOEGLICH EinordnungTechnischeRessource = iota + 1
	WECHSEL_IN_14A_NICHT_MOEGLICH
	BEFRISTET_ALTES_14A
	WECHSEL_DURCHGEFUEHRT
)
