package fehlercode

// Gibt den FehlerCode eines Fehlers an
//
//go:generate stringer --type FehlerCode
//go:generate jsonenums --type FehlerCode
type FehlerCode int

const (
	// ID unbekannt
	ID_UNBEKANNT                       FehlerCode = iota + 1
	ABSENDER_NICHT_ZUGEORDNET                     // Absender zum Zeitpunkt nicht zugeordnet
	EMPFAENGER_NICHT_ZUGEORDNET                   // Empfaenger zum Zeitpunkt nicht zugeordnet
	GERAET_UNBEKANNT                              // Geraet in Messlokation nicht bekannt
	OBIS_UNBEKANNT                                // Obis-Kennzahl in Meldepunkt/Tranche nicht bekannt
	REFERENZIERUNG_FEHLERHAFT                     // Fehlerhafte Referenzierung
	TUPEL_UNBEKANNT                               // Zuodnungstupel unbekannt
	ABSENDER_TUPEL_NICHT_ZUGEORDNET               // Absender zum Zeitpunkt dem Tupel nicht zugeordnet
	EMPFAENGER_TUPEL_NICHT_ZUGEORDNET             // Empfaenger zum Zeitpunkt dem Tupel nicht zugeordnet
	VORKOMMA_ZU_VIELE_STELLEN                     // Vorkommastellen zu lang
	ZEITREIHE_UNVOLLSTAENDIG                      // Zeitreihe ist unvollständig
	REFERENZIERTES_TUPEL_UNBEKANNT                // Referenziertes Tupel ist unbekannt
	MARKTLOKATION_UNBEKANNT                       // Marktlokation nicht gefunden
	MESSLOKATION_UNBEKANNT                        // Messlokation nicht gefunden
	MELDEPUNKT_NICHT_MEHR_IM_NETZ                 // Meldepunkt nicht mehr im Netzgebiet
	ERFORDERLICHE_ANGABE_FEHLT                    // Pflichtfeld nicht gefüllt
	GESCHAEFTSVORFALL_ZURUECKGEWIESEN             // Geschaeftsvorfall wurde zurückgewiesen
	ZEITINTERVALL_NEGATIV                         // Zeitintervall ist negativ oder null
	FORMAT_NICHT_EINGEHALTEN                      // Formatvorgaben wurden nicht eingehalten
	GESCHAEFTSVORFALL_ABSENDER                    // Geschaeftsvorfall darf vom Absender nicht benutzt werden
	KONFIGURATIONSID_UNBEKANNT                    // Konfigurations-ID unbekannt
	SEGMENTWIEDERHOLUNG_UEBERSCHRITTEN            // Segmentwiederholung ueberschritten
	ANZAHLCODES_UEBERSCHRITTEN                    // Anzahl Codes ueberschritten
	ZEITANGABE_UNPLAUSIBEL                        // Zeitangabe unplausibel
	SYNTAXVERSION_NICHT_UNTERSTUETZT              // Syntaxversion nicht unterstuetzt
	FALSCHER_EMPFAENGER                           // Falscher Empfaenger
	WERT_UNGUELTIG                                // Wert ungueltig
	WERT_FEHLT                                    // Wert fehlt
	WERT_UEBERFLUESSIG                            // Wert ueberfluessig
	BEGRENZER_UNPLAUSIBEL                         // Begrenzer unplausibel
	ZEICHEN_UNPLAUSIBEL                           // Zeichen unplausibel
	ABSENDER_UNBEKANNT                            // Absender unbekannt
	TESTKENNZEICHEN_UNPLAUSIBEL                   // Testkennzeichen unplausibel
	DUPLIKAT                                      // Duplikat
	KONTROLLZAEHLER_UNPLAUSIBEL                   // Kontrollzaehler unplausibel
	WERT_ZU_LANG                                  // Wert zu lang
	WIEDERHOLUNG_UNPLAUSIBEL                      // Wiederholung unplausibel
)
