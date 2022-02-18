package korrekturgrund

// Korrekturgrund gives hints on why something needs to be corrected
//go:generate stringer --type Korrekturgrund
//go:generate jsonenums --type Korrekturgrund
type Korrekturgrund int

const (
	// KEIN_ZUGANG corresponds to edi@energy qualifier Z74 "kein Zugang"
	KEIN_ZUGANG Korrekturgrund = iota + 1

	// KOMMUNIKATIONSSTOERUNG corresponds to edi@energy qualifier Z75 "Kommunikationsstörung"
	KOMMUNIKATIONSSTOERUNG

	// NETZAUSFALL corresponds to edi@energy qualifier Z76 "Netzausfall"
	NETZAUSFALL

	// SPANNUNGSAUSFALL corresponds to edi@energy qualifier Z77 "Spannungsausfall"
	SPANNUNGSAUSFALL

	// GERAETEWECHSEL corresponds to edi@energy qualifier Z78 "Gerätewechsel"
	GERAETEWECHSEL

	// KALIBRIERUNG corresponds to edi@energy qualifier Z79 "Kalibrierung"
	KALIBRIERUNG

	// GERAET_ARBEITET_AUSSERHALB_DER_BETRIEBSBEDINGUNG corresponds to edi@energy qualifier
	// Z80 "Gerät arbeitet außerhalb der Betriebsbedingungen"
	GERAET_ARBEITET_AUSSERHALB_DER_BETRIEBSBEDINGUNG

	// MESSEINRICHTUNG_GESTOERT_DEFEKT corresponds to edi@energy qualifier Z81 "Messeinrichtung gestört/defekt"
	MESSEINRICHTUNG_GESTOERT_DEFEKT

	// UNSICHERHEIT_MESSUNG corresponds to edi@energy qualifier Z82 "Unsicherheit Messung"
	UNSICHERHEIT_MESSUNG

	// BERUECKSICHTIGUNG_STOERMENGENZAEHLWERK corresponds to edi@energy qualifier
	// Z98 "Berücksichtigung Störmengenzählwerk"
	BERUECKSICHTIGUNG_STOERMENGENZAEHLWERK

	// MENGENUMWERTUNG_UNVOLLSTAENDIG corresponds to edi@energy qualifier Z99 "Mengenumwertung unvollständig"
	MENGENUMWERTUNG_UNVOLLSTAENDIG

	// UHRZEITGESTELLT_SYNCHRONISATION corresponds to edi@energy qualifier ZA0 "Uhrzeit gestellt / Synchronisation"
	UHRZEITGESTELLT_SYNCHRONISATION

	// MESSWERT_UNPLAUSIBEL corresponds to edi@energy qualifier ZA1 "Messwert unplausibel"
	MESSWERT_UNPLAUSIBEL

	// FALSCHER_WANDLERFAKTOR corresponds to edi@energy qualifier ZA3 "Falscher Wandlerfaktor"
	FALSCHER_WANDLERFAKTOR

	// FEHLERHAFTE_ABLESUNG corresponds to edi@energy qualifier ZA4 "Fehlerhafte Ablesung"
	FEHLERHAFTE_ABLESUNG

	// AENDERUNG_DER_BERECHNUNG corresponds to edi@energy qualifier ZA5 "Änderung der Berechnung"
	AENDERUNG_DER_BERECHNUNG

	// UMBAU_DER_MESSLOKATION corresponds to edi@energy qualifier ZA6 "Umbau der Messlokation"
	UMBAU_DER_MESSLOKATION

	// DATENBEARBEITUNGSFEHLER corresponds to edi@energy qualifier ZA7 "Datenbearbeitungsfehler"
	DATENBEARBEITUNGSFEHLER

	// BRENNWERTKORREKTUR corresponds to edi@energy qualifier ZA8 "Brennwertkorrektur"
	BRENNWERTKORREKTUR

	// Z_ZAHL_KORREKTUR corresponds to edi@energy qualifier ZA9 "Z-Zahl-Korrektur"
	Z_ZAHL_KORREKTUR

	// STOERUNG_DEFEKT_MESSEINRICHTUNG corresponds to edi@energy qualifier ZB0 "Störung / Defekt Messeinrichtung"
	STOERUNG_DEFEKT_MESSEINRICHTUNG

	// AENDERUNG_TARIFSCHALTZEITEN corresponds to edi@energy qualifier ZB9 "Änderung Tarifschaltzeiten"
	AENDERUNG_TARIFSCHALTZEITEN

	// TARIFSCHALTGERAET_DEFEKT corresponds to edi@energy qualifier ZC2 "Tarifschaltgerät defekt"
	TARIFSCHALTGERAET_DEFEKT

	// IMPULSWERTIGKEIT_NICHT_AUSREICHEND corresponds to edi@energy qualifier ZC4 "Impulswertigkeit nicht ausreichend"
	IMPULSWERTIGKEIT_NICHT_AUSREICHEND

	// ENERGIEMENGE_IN_UNGEMESSENEM_ZEITINTERVALL corresponds to edi@energy qualifier
	// ZJ8 "Energiemenge in ungemessenem Zeitintervall"
	ENERGIEMENGE_IN_UNGEMESSENEM_ZEITINTERVALL

	// ENERGIE_AUS_DEM_UNGEPAIRTEN_ZEITINTERVALL corresponds to edi@energy qualifier
	// ZJ9 "Energie aus dem ungepairten Zeitintervall"
	ENERGIE_AUS_DEM_UNGEPAIRTEN_ZEITINTERVALL

	// WARTUNGSARBEITEN_AN_GEEICHTEM_MESSGERAET corresponds to edi@energy qualifier
	// ZR1 "Wartungsarbeiten an geeichtem Messgerät"
	WARTUNGSARBEITEN_AN_GEEICHTEM_MESSGERAET

	// GESTOERTE_WERTE corresponds to edi@energy qualifier ZR2 "Gestörte Werte"
	GESTOERTE_WERTE

	// WARTUNGSARBEITEN_AN_EICHRECHTSKONFORMEN_MESSGERAETEN corresponds to edi@energy qualifier
	// ZR3 "Wartungsarbeiten an eichrechtskonformen Messgeräten"
	WARTUNGSARBEITEN_AN_EICHRECHTSKONFORMEN_MESSGERAETEN

	// KONSISTENZ_UND_SYNCHRONPRUEFUNG corresponds to edi@energy qualifier ZR4 "Konsistenz- und Synchronprüfung"
	KONSISTENZ_UND_SYNCHRONPRUEFUNG
)
