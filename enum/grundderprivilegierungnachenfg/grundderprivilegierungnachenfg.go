package grundderprivilegierungnachenfg

// GrundDerPrivilegierungNachEnFG is the Grund der Privilegierung nach EnFG (UTILMD Strom)
//
//go:generate stringer --type GrundDerPrivilegierungNachEnFG
//go:generate jsonenums --type GrundDerPrivilegierungNachEnFG
type GrundDerPrivilegierungNachEnFG int

const (
	// STROMSPEICHER_UND_VERLUSTENERGIE - § 21 EnFG Stromspeicher und Verlustenergie
	STROMSPEICHER_UND_VERLUSTENERGIE GrundDerPrivilegierungNachEnFG = iota + 1

	// ELEKTRISCH_ANGETRIEBENE_WAERMEPUMPEN - § 22 EnFG elektrisch angetriebene Wärmepumpen
	ELEKTRISCH_ANGETRIEBENE_WAERMEPUMPEN

	// UMLAGEERHEBUNG_BEI_ANLAGEN_ZUR_VERSTROMUNG_VON_KUPPELGASEN - § 23 EnFG Umlageerhebung bei Anlagen zur Verstromung von Kuppelgasen
	UMLAGEERHEBUNG_BEI_ANLAGEN_ZUR_VERSTROMUNG_VON_KUPPELGASEN

	// HERSTELLUNG_VON_GRUENEN_WASSERSTOFF - § 24 EnFG Herstellung von Grünen Wasserstoff
	HERSTELLUNG_VON_GRUENEN_WASSERSTOFF

	// STROMKOSTENINTENSIVE_UNTERNEHMEN - §§ 30 - 35 EnFG stromkostenintensive Unternehmen
	STROMKOSTENINTENSIVE_UNTERNEHMEN

	// HERSTELLUNG_VON_WASSERSTOFF_IN_STROMKOSTENINTENSIVEN_UNTERNEHMEN - § 36 EnFG Herstellung von Wasserstoff in stromkostenintensiven Unternehmen
	HERSTELLUNG_VON_WASSERSTOFF_IN_STROMKOSTENINTENSIVEN_UNTERNEHMEN

	// SCHIENENBAHNEN - § 37 EnFG Schienenbahnen
	SCHIENENBAHNEN

	// ELEKTRISCHE_BETRIEBENE_BUSSEN_IM_LINIENVERKEHR - § 38 EnFG elektrische betriebene Bussen im Linienverkehr
	ELEKTRISCHE_BETRIEBENE_BUSSEN_IM_LINIENVERKEHR

	// LANDSTROMANLAGEN - § 39 EnFG Landstromanlagen
	LANDSTROMANLAGEN
)
