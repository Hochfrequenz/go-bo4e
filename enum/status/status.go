package status

//go:generate stringer --type Status
//go:generate jsonenums --type Status

type Status int

const (

	// TARIF_1 represents MSCONS SG10 STS 4405 value T1
	TARIF_1 Status = iota + 1

	// TARIF_2 represents MSCONS SG10 STS 4405 value T2
	TARIF_2

	// TARIF_3 represents MSCONS SG10 STS 4405 value T3
	TARIF_3

	// TARIF_4 represents MSCONS SG10 STS 4405 value T4
	TARIF_4

	// TARIF_5 represents MSCONS SG10 STS 4405 value T5
	TARIF_5

	// TARIF_6 represents MSCONS SG10 STS 4405 value T6
	TARIF_6

	// TARIF_7 represents MSCONS SG10 STS 4405 value T7
	TARIF_7

	// TARIF_8 represents MSCONS SG10 STS 4405 value T8
	TARIF_8

	// TARIF_9 represents MSCONS SG10 STS 4405 value T9
	TARIF_9

	// ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT
	// represents MSCONS SG10 STS 4405 value Z36
	ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT

	// ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT
	// represents MSCONS SG10 STS 4405 value Z37
	ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT

	// ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG
	// represents MSCONS SG10 STS 4405 value Z38
	ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG

	// ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG
	// represents MSCONS SG10 STS 4405 value Z39
	ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG

	// KEIN_ZUGANG represents MSCONS SG10 STS 4405 value Z74
	KEIN_ZUGANG

	// KOMMUNIKATIONSSTOERUNG represents MSCONS SG10 STS 4405 value Z75
	KOMMUNIKATIONSSTOERUNG

	// NETZAUSFALL represents MSCONS SG10 STS 4405 value Z76
	NETZAUSFALL

	// SPANNUNGSAUSFALL represents MSCONS SG10 STS 4405 value Z77
	SPANNUNGSAUSFALL

	// STATUS_GERAETEWECHSEL represents MSCONS SG10 STS 4405 value Z78.
	// Needs to have a prefix, because this is also an enum value in "Themengebiet".
	STATUS_GERAETEWECHSEL

	// KALIBRIERUNG represents MSCONS SG10 STS 4405 value Z79
	KALIBRIERUNG

	// GERAET_ARBEITET_AUSSERHALB_DER_BETRIEBSBEDINGUNGEN represents MSCONS SG10 STS 4405 value Z80
	GERAET_ARBEITET_AUSSERHALB_DER_BETRIEBSBEDINGUNGEN

	// MESSEINRICHTUNG_GESTOERT_DEFEKT represents MSCONS SG10 STS 4405 value Z81
	MESSEINRICHTUNG_GESTOERT_DEFEKT

	// UNSICHERHEIT_MESSUNG represents MSCONS SG10 STS 4405 value Z82
	UNSICHERHEIT_MESSUNG

	// KUNDENSELBSTABLESUNG represents MSCONS SG10 STS 4405 value Z83
	KUNDENSELBSTABLESUNG

	// LEERSTAND represents MSCONS SG10 STS 4405 value Z84
	LEERSTAND

	// REALER_ZAEHLERUEBERLAUF_GEPRUEFT represents MSCONS SG10 STS 4405 value Z85
	REALER_ZAEHLERUEBERLAUF_GEPRUEFT

	// PLAUSIBEL_WG_KONTROLLABLESUNG represents MSCONS SG10 STS 4405 value Z86
	PLAUSIBEL_WG_KONTROLLABLESUNG

	// PLAUSIBEL_WG_KUNDENHINWEIS represents MSCONS SG10 STS 4405 value Z87
	PLAUSIBEL_WG_KUNDENHINWEIS

	// VERGLEICHSMESSUNG_GEEICHT represents MSCONS SG10 STS 4405 value Z88
	VERGLEICHSMESSUNG_GEEICHT

	// VERGLEICHSMESSUNG_NICHT_GEEICHT represents MSCONS SG10 STS 4405 value Z89
	VERGLEICHSMESSUNG_NICHT_GEEICHT

	// MESSWERTNACHBILDUNG_AUS_GEEICHTEN_WERTEN represents MSCONS SG10 STS 4405 value Z90
	MESSWERTNACHBILDUNG_AUS_GEEICHTEN_WERTEN

	// MESSWERTNACHBILDUNG_AUS_NICHT_GEEICHTEN_WERTEN represents MSCONS SG10 STS 4405 value Z91
	MESSWERTNACHBILDUNG_AUS_NICHT_GEEICHTEN_WERTEN

	// INTERPOLATION represents MSCONS SG10 STS 4405 value Z92
	INTERPOLATION

	// HALTEWERT represents MSCONS SG10 STS 4405 value Z93
	HALTEWERT

	// BILANZIERUNG_NETZABSCHNITT represents MSCONS SG10 STS 4405 value Z94
	BILANZIERUNG_NETZABSCHNITT

	// HISTORISCHE_MESSWERTE represents MSCONS SG10 STS 4405 value Z95
	HISTORISCHE_MESSWERTE

	// BERUECKSICHTIGUNG_STOERMENGENZAEHLWERK represents MSCONS SG10 STS 4405 value Z98
	BERUECKSICHTIGUNG_STOERMENGENZAEHLWERK

	// MENGENUMWERTUNG_VOLLSTAENDIG represents MSCONS SG10 STS 4405 value Z99
	MENGENUMWERTUNG_VOLLSTAENDIG

	// UHRZEIT_GESTELLT_SYNCHRONISATION represents MSCONS SG10 STS 4405 value ZA0
	UHRZEIT_GESTELLT_SYNCHRONISATION

	// MESSWERT_UNPLAUSIBEL represents MSCONS SG10 STS 4405 value ZA1
	MESSWERT_UNPLAUSIBEL

	// FALSCHER_WANDLERFAKTOR represents MSCONS SG10 STS 4405 value ZA3
	FALSCHER_WANDLERFAKTOR

	// FEHLERHAFTE_ABLESUNG represents MSCONS SG10 STS 4405 value ZA4
	FEHLERHAFTE_ABLESUNG

	// AENDERUNG_DER_BERECHNUNG represents MSCONS SG10 STS 4405 value ZA5
	AENDERUNG_DER_BERECHNUNG

	// UMBAU_DER_MESSLOKATION represents MSCONS SG10 STS 4405 value ZA6
	UMBAU_DER_MESSLOKATION

	// DATENBEARBEITUNGSFEHLER represents MSCONS SG10 STS 4405 value ZA7
	DATENBEARBEITUNGSFEHLER

	// BRENNWERTKORREKTUR represents MSCONS SG10 STS 4405 value ZA8
	BRENNWERTKORREKTUR

	// Z_ZAHL_KORREKTUR represents MSCONS SG10 STS 4405 value ZA9
	Z_ZAHL_KORREKTUR

	// STOERUNG_DEFEKT_MESSEINRICHTUNG represents MSCONS SG10 STS 4405 value ZB0
	STOERUNG_DEFEKT_MESSEINRICHTUNG

	// AENDERUNG_TARIFSCHALTZEITEN represents MSCONS SG10 STS 4405 value ZB9
	AENDERUNG_TARIFSCHALTZEITEN

	// TARIFSCHALTGERAET_DEFEKT represents MSCONS SG10 STS 4405 value ZC2
	TARIFSCHALTGERAET_DEFEKT

	// AUSTAUSCH_DES_ERSATZWERTES represents MSCONS SG10 STS 4405 value ZC3
	AUSTAUSCH_DES_ERSATZWERTES

	// IMPULSWERTIGKEIT_NICHT_AUSREICHEND represents MSCONS SG10 STS 4405 value ZC4
	IMPULSWERTIGKEIT_NICHT_AUSREICHEND

	// UMSTELLUNG_GASQUALITAET represents MSCONS SG10 STS 4405 value ZG3
	UMSTELLUNG_GASQUALITAET

	// STATISTISCHE_METHODE represents MSCONS SG10 STS 4405 value ZJ2
	STATISTISCHE_METHODE

	// ENERGIEMENGE_IN_UNGEMESSENEM_ZEITINTERVALL represents MSCONS SG10 STS 4405 value ZJ8
	ENERGIEMENGE_IN_UNGEMESSENEM_ZEITINTERVALL

	// ENERGIEMENGE_AUS_DEM_UNGEPAIRTEN_ZEITINTERVALL represents MSCONS SG10 STS 4405 value ZJ9
	ENERGIEMENGE_AUS_DEM_UNGEPAIRTEN_ZEITINTERVALL

	// AUFTEILUNG represents MSCONS SG10 STS 4405 value ZQ8
	AUFTEILUNG

	// VERWENDUNG_VON_WERTEN_DES_STOERMENGENZAEHLWERKS represents MSCONS SG10 STS 4405 value ZQ9
	VERWENDUNG_VON_WERTEN_DES_STOERMENGENZAEHLWERKS

	// UMGANGS_UND_KORREKTURMENGEN represents MSCONS SG10 STS 4405 value ZR0
	UMGANGS_UND_KORREKTURMENGEN

	// WARTUNGSARBEITEN_AN_GEEICHTEM_MESSGERAET represents MSCONS SG10 STS 4405 value ZR1
	WARTUNGSARBEITEN_AN_GEEICHTEM_MESSGERAET

	// GESTOERTE_WERTE represents MSCONS SG10 STS 4405 value ZR2
	GESTOERTE_WERTE

	// WARTUNGSARBEITEN_AN_EICHRECHTSKONFORMEN_MESSGERAETEN represents MSCONS SG10 STS 4405 value ZR3
	WARTUNGSARBEITEN_AN_EICHRECHTSKONFORMEN_MESSGERAETEN

	// KONSISTENZ_UND_SYNCHRONPRUEFUNG represents MSCONS SG10 STS 4405 value ZR4
	KONSISTENZ_UND_SYNCHRONPRUEFUNG

	// RECHENWERT represents MSCONS SG10 STS 4405 value ZR5
	RECHENWERT
	ANGABEN_MESSLOKATION // ZS0
	BASIS_MME            // ZS2

	GRUND_ANGABEN_MESSLOKATION                                                       // ZS9
	ANFORDERUNG_IN_DIE_VERGANGENHEIT_ZUM_ANGEFORDERTEN_ZEITPUNKT_LIEGT_KEIN_WERT_VOR // ZT8
)
