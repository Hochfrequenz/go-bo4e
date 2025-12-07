package enum_compatibility_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hochfrequenz/go-bo4e/enum/angebotsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/artvolumenerfassung"
	"github.com/hochfrequenz/go-bo4e/enum/aufabschlagsziel"
	"github.com/hochfrequenz/go-bo4e/enum/auftragsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/auftragsstornogrund"
	"github.com/hochfrequenz/go-bo4e/enum/ausschreibungsportal"
	"github.com/hochfrequenz/go-bo4e/enum/ausschreibungsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/ausschreibungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/bearbeitungsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/befestigungsart"
	"github.com/hochfrequenz/go-bo4e/enum/berechnungsformelablehngrund"
	"github.com/hochfrequenz/go-bo4e/enum/berechnungsformelnotwendigkeit"
	"github.com/hochfrequenz/go-bo4e/enum/betriebszustand"
	"github.com/hochfrequenz/go-bo4e/enum/bezeichnungsummenzeitreihe"
	"github.com/hochfrequenz/go-bo4e/enum/blindarbeitszahler"
	"github.com/hochfrequenz/go-bo4e/enum/einordnungtechnischeressource"
	"github.com/hochfrequenz/go-bo4e/enum/energieflussrichtung"
	"github.com/hochfrequenz/go-bo4e/enum/geraeteart"
	"github.com/hochfrequenz/go-bo4e/enum/grundlagezurverringerungderumlagennachenfg"
	"github.com/hochfrequenz/go-bo4e/enum/gueltigkeitstyp"
	"github.com/hochfrequenz/go-bo4e/enum/inbetriebsetzungtechnischeressource"
	"github.com/hochfrequenz/go-bo4e/enum/informationweiteretechnischeressource"
	"github.com/hochfrequenz/go-bo4e/enum/kategorietechnischeressource"
	"github.com/hochfrequenz/go-bo4e/enum/kundengruppe"
	"github.com/hochfrequenz/go-bo4e/enum/kundentyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengenoperator"
	"github.com/hochfrequenz/go-bo4e/enum/messpreistyp"
	"github.com/hochfrequenz/go-bo4e/enum/normierungsfaktor"
	"github.com/hochfrequenz/go-bo4e/enum/oekolabel"
	"github.com/hochfrequenz/go-bo4e/enum/oekozertifikat"
	"github.com/hochfrequenz/go-bo4e/enum/preisgarantietyp"
	"github.com/hochfrequenz/go-bo4e/enum/preismodell"
	"github.com/hochfrequenz/go-bo4e/enum/preistyp"
	"github.com/hochfrequenz/go-bo4e/enum/prioritaet"
	"github.com/hochfrequenz/go-bo4e/enum/profilklasse"
	"github.com/hochfrequenz/go-bo4e/enum/qualitaet"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungslegung"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungspositionsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/regionskriteriumtyp"
	"github.com/hochfrequenz/go-bo4e/enum/schalthandlung"
	"github.com/hochfrequenz/go-bo4e/enum/schwachlastfaehig"
	"github.com/hochfrequenz/go-bo4e/enum/servicetyp"
	"github.com/hochfrequenz/go-bo4e/enum/sperrauftragsablehngrund"
	"github.com/hochfrequenz/go-bo4e/enum/sperrauftragsart"
	"github.com/hochfrequenz/go-bo4e/enum/sperrauftragsverhinderungsgrund"
	"github.com/hochfrequenz/go-bo4e/enum/tarifkalkulationsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/tarifmerkmal"
	"github.com/hochfrequenz/go-bo4e/enum/tarifregionskriterium"
	"github.com/hochfrequenz/go-bo4e/enum/tariftyp"
	"github.com/hochfrequenz/go-bo4e/enum/verbrauchsmengetyp"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsform"
	"github.com/hochfrequenz/go-bo4e/enum/vertragstatus"
	"github.com/hochfrequenz/go-bo4e/enum/voraussetzungen"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertypspezifikation"
)

// AllEnumsTest mirrors the C# AllEnumsTest class for direct JSON deserialization.
type AllEnumsTest struct {
	Angebotsstatus_KONZEPTION                                                       angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_KONZEPTION"`
	Angebotsstatus_UNVERBINDLICH                                                    angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_UNVERBINDLICH"`
	Angebotsstatus_VERBINDLICH                                                      angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_VERBINDLICH"`
	Angebotsstatus_BEAUFTRAGT                                                       angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_BEAUFTRAGT"`
	Angebotsstatus_UNGUELTIG                                                        angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_UNGUELTIG"`
	Angebotsstatus_ABGELEHNT                                                        angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_ABGELEHNT"`
	Angebotsstatus_NACHGEFASST                                                      angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_NACHGEFASST"`
	Angebotsstatus_AUSSTEHEND                                                       angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_AUSSTEHEND"`
	Angebotsstatus_ERLEDIGT                                                         angebotsstatus.Angebotsstatus                                                       `json:"Angebotsstatus_ERLEDIGT"`
	ArtVolumenerfassung_KENNLINIENKORREKTUR                                         artvolumenerfassung.ArtVolumenerfassung                                             `json:"ArtVolumenerfassung_KENNLINIENKORREKTUR"`
	ArtVolumenerfassung_SCHLEICHMENGENUNTERDRUECKUNG                                artvolumenerfassung.ArtVolumenerfassung                                             `json:"ArtVolumenerfassung_SCHLEICHMENGENUNTERDRÜCKUNG"`
	AufAbschlagsziel_ARBEITSPREIS_HT                                                aufabschlagsziel.AufAbschlagsziel                                                   `json:"AufAbschlagsziel_ARBEITSPREIS_HT"`
	AufAbschlagsziel_ARBEITSPREIS_NT                                                aufabschlagsziel.AufAbschlagsziel                                                   `json:"AufAbschlagsziel_ARBEITSPREIS_NT"`
	AufAbschlagsziel_ARBEITSPREIS_HT_NT                                             aufabschlagsziel.AufAbschlagsziel                                                   `json:"AufAbschlagsziel_ARBEITSPREIS_HT_NT"`
	AufAbschlagsziel_GRUNDPREIS                                                     aufabschlagsziel.AufAbschlagsziel                                                   `json:"AufAbschlagsziel_GRUNDPREIS"`
	AufAbschlagsziel_GESAMTPREIS                                                    aufabschlagsziel.AufAbschlagsziel                                                   `json:"AufAbschlagsziel_GESAMTPREIS"`
	Auftragsstatus_GESCHEITERT                                                      auftragsstatus.Auftragsstatus                                                       `json:"Auftragsstatus_GESCHEITERT"`
	Auftragsstatus_ERFOLGREICH                                                      auftragsstatus.Auftragsstatus                                                       `json:"Auftragsstatus_ERFOLGREICH"`
	Auftragsstatus_GEPLANT                                                          auftragsstatus.Auftragsstatus                                                       `json:"Auftragsstatus_GEPLANT"`
	Auftragsstatus_ZUGESTIMMT                                                       auftragsstatus.Auftragsstatus                                                       `json:"Auftragsstatus_ZUGESTIMMT"`
	Auftragsstatus_WIDERSPROCHEN                                                    auftragsstatus.Auftragsstatus                                                       `json:"Auftragsstatus_WIDERSPROCHEN"`
	Auftragsstatus_ABGELEHNT                                                        auftragsstatus.Auftragsstatus                                                       `json:"Auftragsstatus_ABGELEHNT"`
	Auftragsstornogrund_STORNIERUNG_DER_ENTSPERRUNG_ERFOLGT                         auftragsstornogrund.Auftragsstornogrund                                             `json:"Auftragsstornogrund_STORNIERUNG_DER_ENTSPERRUNG_ERFOLGT"`
	Auftragsstornogrund_STORNIERUNG_DER_SPERRUNG_ERFOLGT                            auftragsstornogrund.Auftragsstornogrund                                             `json:"Auftragsstornogrund_STORNIERUNG_DER_SPERRUNG_ERFOLGT"`
	Ausschreibungsportal_ENPORTAL                                                   ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_ENPORTAL"`
	Ausschreibungsportal_ENERGIE_AGENTUR                                            ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_ENERGIE_AGENTUR"`
	Ausschreibungsportal_BMWI                                                       ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_BMWI"`
	Ausschreibungsportal_ENERGIE_HANDELSPLATZ                                       ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_ENERGIE_HANDELSPLATZ"`
	Ausschreibungsportal_BUND                                                       ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_BUND"`
	Ausschreibungsportal_VERA_ONLINE                                                ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_VERA_ONLINE"`
	Ausschreibungsportal_ISPEX                                                      ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_ISPEX"`
	Ausschreibungsportal_ENERGIEMARKTPLATZ                                          ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_ENERGIEMARKTPLATZ"`
	Ausschreibungsportal_EVERGABE                                                   ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_EVERGABE"`
	Ausschreibungsportal_DTAD                                                       ausschreibungsportal.Ausschreibungsportal                                           `json:"Ausschreibungsportal_DTAD"`
	Ausschreibungsstatus_PHASE1                                                     ausschreibungsstatus.Ausschreibungsstatus                                           `json:"Ausschreibungsstatus_PHASE1"`
	Ausschreibungsstatus_PHASE2                                                     ausschreibungsstatus.Ausschreibungsstatus                                           `json:"Ausschreibungsstatus_PHASE2"`
	Ausschreibungsstatus_PHASE3                                                     ausschreibungsstatus.Ausschreibungsstatus                                           `json:"Ausschreibungsstatus_PHASE3"`
	Ausschreibungsstatus_PHASE4                                                     ausschreibungsstatus.Ausschreibungsstatus                                           `json:"Ausschreibungsstatus_PHASE4"`
	Ausschreibungstyp_OEFFENTLICHRECHTLICH                                          ausschreibungstyp.Ausschreibungstyp                                                 `json:"Ausschreibungstyp_OEFFENTLICHRECHTLICH"`
	Ausschreibungstyp_EUROPAWEIT                                                    ausschreibungstyp.Ausschreibungstyp                                                 `json:"Ausschreibungstyp_EUROPAWEIT"`
	Bearbeitungsstatus_OFFEN                                                        bearbeitungsstatus.Bearbeitungsstatus                                               `json:"Bearbeitungsstatus_OFFEN"`
	Bearbeitungsstatus_IN_BEARBEITUNG                                               bearbeitungsstatus.Bearbeitungsstatus                                               `json:"Bearbeitungsstatus_IN_BEARBEITUNG"`
	Bearbeitungsstatus_ABGESCHLOSSEN                                                bearbeitungsstatus.Bearbeitungsstatus                                               `json:"Bearbeitungsstatus_ABGESCHLOSSEN"`
	Bearbeitungsstatus_STORNIERT                                                    bearbeitungsstatus.Bearbeitungsstatus                                               `json:"Bearbeitungsstatus_STORNIERT"`
	Bearbeitungsstatus_QUITTIERT                                                    bearbeitungsstatus.Bearbeitungsstatus                                               `json:"Bearbeitungsstatus_QUITTIERT"`
	Bearbeitungsstatus_IGNORIERT                                                    bearbeitungsstatus.Bearbeitungsstatus                                               `json:"Bearbeitungsstatus_IGNORIERT"`
	Befestigungsart_STECKTECHNIK                                                    befestigungsart.Befestigungsart                                                     `json:"Befestigungsart_STECKTECHNIK"`
	Befestigungsart_DREIPUNKT                                                       befestigungsart.Befestigungsart                                                     `json:"Befestigungsart_DREIPUNKT"`
	Befestigungsart_HUTSCHIENE                                                      befestigungsart.Befestigungsart                                                     `json:"Befestigungsart_HUTSCHIENE"`
	Befestigungsart_EINSTUTZEN                                                      befestigungsart.Befestigungsart                                                     `json:"Befestigungsart_EINSTUTZEN"`
	Befestigungsart_ZWEISTUTZEN                                                     befestigungsart.Befestigungsart                                                     `json:"Befestigungsart_ZWEISTUTZEN"`
	BerechnungsformelAblehngrund_START_DATUM_UNPLAUSIBEL                            berechnungsformelablehngrund.BerechnungsformelAblehngrund                           `json:"BerechnungsformelAblehngrund_START_DATUM_UNPLAUSIBEL"`
	BerechnungsformelAblehngrund_ZU_VIELE_MESSLOKATIONEN                            berechnungsformelablehngrund.BerechnungsformelAblehngrund                           `json:"BerechnungsformelAblehngrund_ZU_VIELE_MESSLOKATIONEN"`
	BerechnungsformelAblehngrund_FEHLENDE_MESSLOKATIONEN                            berechnungsformelablehngrund.BerechnungsformelAblehngrund                           `json:"BerechnungsformelAblehngrund_FEHLENDE_MESSLOKATIONEN"`
	BerechnungsformelAblehngrund_FALSCHE_MESSLOKATION                               berechnungsformelablehngrund.BerechnungsformelAblehngrund                           `json:"BerechnungsformelAblehngrund_FALSCHE_MESSLOKATION"`
	BerechnungsformelNotwendigkeit_BERECHNUNGSFORMEL_NOTWENDIG                      berechnungsformelnotwendigkeit.BerechnungsformelNotwendigkeit                       `json:"BerechnungsformelNotwendigkeit_BERECHNUNGSFORMEL_NOTWENDIG"`
	BerechnungsformelNotwendigkeit_BERECHNUNGSFORMEL_TRIVIAL                        berechnungsformelnotwendigkeit.BerechnungsformelNotwendigkeit                       `json:"BerechnungsformelNotwendigkeit_BERECHNUNGSFORMEL_TRIVIAL"`
	BerechnungsformelNotwendigkeit_BERECHNUNGSFORMEL_NICHT_NOTWENDIG                berechnungsformelnotwendigkeit.BerechnungsformelNotwendigkeit                       `json:"BerechnungsformelNotwendigkeit_BERECHNUNGSFORMEL_NICHT_NOTWENDIG"`
	Betriebszustand_GESPERRT_NICHT_ENTSPERREN                                       betriebszustand.Betriebszustand                                                     `json:"Betriebszustand_GESPERRT_NICHT_ENTSPERREN"`
	Betriebszustand_GESPERRT                                                        betriebszustand.Betriebszustand                                                     `json:"Betriebszustand_GESPERRT"`
	Betriebszustand_REGELBETRIEB                                                    betriebszustand.Betriebszustand                                                     `json:"Betriebszustand_REGELBETRIEB"`
	BezeichnungSummenzeitreihe_BG_SZR_B                                              bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_BG_SZR_B"`
	BezeichnungSummenzeitreihe_BG_SZR_C                                              bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_BG_SZR_C"`
	BezeichnungSummenzeitreihe_BK_SZR_A                                              bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_BK_SZR_A"`
	BezeichnungSummenzeitreihe_BK_SZR_B_RZ                                           bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_BK_SZR_B_RZ"`
	BezeichnungSummenzeitreihe_BK_SZR_B_BG                                           bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_BK_SZR_B_BG"`
	BezeichnungSummenzeitreihe_BK_SZR_C                                              bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_BK_SZR_C"`
	BezeichnungSummenzeitreihe_LF_SZR_A                                              bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_LF_SZR_A"`
	BezeichnungSummenzeitreihe_LF_SZR_B_RZ                                           bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_LF_SZR_B_RZ"`
	BezeichnungSummenzeitreihe_LF_SZR_B_BG                                           bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_LF_SZR_B_BG"`
	BezeichnungSummenzeitreihe_DZUE                                                  bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_DZUE"`
	BezeichnungSummenzeitreihe_NZR                                                   bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_NZR"`
	BezeichnungSummenzeitreihe_ASZR                                                  bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_ASZR"`
	BezeichnungSummenzeitreihe_NGZ                                                   bezeichnungsummenzeitreihe.BezeichnungSummenzeitreihe                               `json:"BezeichnungSummenzeitreihe_NGZ"`
	Blindarbeitszahler_ANSCHLUSSNEHMER                                              blindarbeitszahler.Blindarbeitszahler                                               `json:"Blindarbeitszahler_ANSCHLUSSNEHMER"`
	Blindarbeitszahler_LIEFERANT                                                    blindarbeitszahler.Blindarbeitszahler                                               `json:"Blindarbeitszahler_LIEFERANT"`
	EinordnungTechnischeRessource_WECHSEL_IN_14A_EINMALIG_MOEGLICH                  einordnungtechnischeressource.EinordnungTechnischeRessource                         `json:"EinordnungTechnischeRessource_WECHSEL_IN_14A_EINMALIG_MOEGLICH"`
	EinordnungTechnischeRessource_WECHSEL_IN_14A_NICHT_MOEGLICH                     einordnungtechnischeressource.EinordnungTechnischeRessource                         `json:"EinordnungTechnischeRessource_WECHSEL_IN_14A_NICHT_MOEGLICH"`
	EinordnungTechnischeRessource_BEFRISTET_ALTES_14A                               einordnungtechnischeressource.EinordnungTechnischeRessource                         `json:"EinordnungTechnischeRessource_BEFRISTET_ALTES_14A"`
	EinordnungTechnischeRessource_WECHSEL_DURCHGEFUEHRT                             einordnungtechnischeressource.EinordnungTechnischeRessource                         `json:"EinordnungTechnischeRessource_WECHSEL_DURCHGEFUEHRT"`
	Energieflussrichtung_VERBRAUCH                                                  energieflussrichtung.Energieflussrichtung                                           `json:"Energieflussrichtung_VERBRAUCH"`
	Energieflussrichtung_ERZEUGUNG                                                  energieflussrichtung.Energieflussrichtung                                           `json:"Energieflussrichtung_ERZEUGUNG"`
	Geraeteart_WANDLER                                                              geraeteart.Geraeteart                                                               `json:"Geraeteart_WANDLER"`
	Geraeteart_KOMMUNIKATIONSEINRICHTUNG                                            geraeteart.Geraeteart                                                               `json:"Geraeteart_KOMMUNIKATIONSEINRICHTUNG"`
	Geraeteart_TECHNISCHE_STEUEREINRICHTUNG                                         geraeteart.Geraeteart                                                               `json:"Geraeteart_TECHNISCHE_STEUEREINRICHTUNG"`
	Geraeteart_MENGENUMWERTER                                                       geraeteart.Geraeteart                                                               `json:"Geraeteart_MENGENUMWERTER"`
	Geraeteart_SMARTMETER_GATEWAY                                                   geraeteart.Geraeteart                                                               `json:"Geraeteart_SMARTMETER_GATEWAY"`
	Geraeteart_STEUERBOX                                                            geraeteart.Geraeteart                                                               `json:"Geraeteart_STEUERBOX"`
	Geraeteart_ZAEHLEINRICHTUNG                                                     geraeteart.Geraeteart                                                               `json:"Geraeteart_ZAEHLEINRICHTUNG"`
	GrundlageZurVerringerungDerUmlagenNachEnfg_KUNDE_ERFUELLT_VORAUSSETZUNG         grundlagezurverringerungderumlagennachenfg.GrundlageZurVerringerungDerUmlagenNachEnfg `json:"GrundlageZurVerringerungDerUmlagenNachEnfg_KUNDE_ERFUELLT_VORAUSSETZUNG"`
	GrundlageZurVerringerungDerUmlagenNachEnfg_KUNDE_ERFUELLT_VORAUSSETZUNG_NICHT   grundlagezurverringerungderumlagennachenfg.GrundlageZurVerringerungDerUmlagenNachEnfg `json:"GrundlageZurVerringerungDerUmlagenNachEnfg_KUNDE_ERFUELLT_VORAUSSETZUNG_NICHT"`
	GrundlageZurVerringerungDerUmlagenNachEnfg_KEINE_ANGABE                         grundlagezurverringerungderumlagennachenfg.GrundlageZurVerringerungDerUmlagenNachEnfg `json:"GrundlageZurVerringerungDerUmlagenNachEnfg_KEINE_ANGABE"`
	Gueltigkeitstyp_NICHT_IN                                                        gueltigkeitstyp.Gueltigkeitstyp                                                     `json:"Gueltigkeitstyp_NICHT_IN"`
	InbetriebsetzungTechnischeRessource_NACH_2023                                   inbetriebsetzungtechnischeressource.InbetriebsetzungTechnischeRessource             `json:"InbetriebsetzungTechnischeRessource_NACH_2023"`
	InbetriebsetzungTechnischeRessource_VOR_2024                                    inbetriebsetzungtechnischeressource.InbetriebsetzungTechnischeRessource             `json:"InbetriebsetzungTechnischeRessource_VOR_2024"`
	InformationWeitereTechnischeRessource_WEITERE_EINRICHTUNG_VORHANDEN             informationweiteretechnischeressource.InformationWeitereTechnischeRessource         `json:"InformationWeitereTechnischeRessource_WEITERE_EINRICHTUNG_VORHANDEN"`
	InformationWeitereTechnischeRessource_KEINE_WEITERE_EINRICHTUNG_VORHANDEN       informationweiteretechnischeressource.InformationWeitereTechnischeRessource         `json:"InformationWeitereTechnischeRessource_KEINE_WEITERE_EINRICHTUNG_VORHANDEN"`
	KategorieTechnischeRessource_FAELLT_UNTER_14A                                   kategorietechnischeressource.KategorieTechnischeRessource                           `json:"KategorieTechnischeRessource_FAELLT_UNTER_14A"`
	KategorieTechnischeRessource_FAELLT_NICHT_UNTER_14A                             kategorietechnischeressource.KategorieTechnischeRessource                           `json:"KategorieTechnischeRessource_FAELLT_NICHT_UNTER_14A"`
	Kundengruppe_RLM                                                                kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_RLM"`
	Kundengruppe_SLP_S_G0                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G0"`
	Kundengruppe_SLP_S_G1                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G1"`
	Kundengruppe_SLP_S_G2                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G2"`
	Kundengruppe_SLP_S_G3                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G3"`
	Kundengruppe_SLP_S_G4                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G4"`
	Kundengruppe_SLP_S_G5                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G5"`
	Kundengruppe_SLP_S_G6                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G6"`
	Kundengruppe_SLP_S_G7                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_G7"`
	Kundengruppe_SLP_S_L0                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_L0"`
	Kundengruppe_SLP_S_L1                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_L1"`
	Kundengruppe_SLP_S_L2                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_L2"`
	Kundengruppe_SLP_S_H0                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_H0"`
	Kundengruppe_SLP_S_SB                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_SB"`
	Kundengruppe_SLP_S_HZ                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_HZ"`
	Kundengruppe_SLP_S_WP                                                           kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_S_WP"`
	Kundengruppe_SLP_G_GKO                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GKO"`
	Kundengruppe_SLP_G_GHA                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GHA"`
	Kundengruppe_SLP_G_GMK                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GMK"`
	Kundengruppe_SLP_G_GBD                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GBD"`
	Kundengruppe_SLP_G_GGA                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GGA"`
	Kundengruppe_SLP_G_GBH                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GBH"`
	Kundengruppe_SLP_G_GBA                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GBA"`
	Kundengruppe_SLP_G_GWA                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GWA"`
	Kundengruppe_SLP_G_GGB                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GGB"`
	Kundengruppe_SLP_G_GPD                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GPD"`
	Kundengruppe_SLP_G_GMF                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_GMF"`
	Kundengruppe_SLP_G_HEF                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_HEF"`
	Kundengruppe_SLP_G_HMF                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_HMF"`
	Kundengruppe_SLP_G_HKO                                                          kundengruppe.Kundengruppe                                                           `json:"Kundengruppe_SLP_G_HKO"`
	Kundentyp_PRIVAT                                                                kundentyp.Kundentyp                                                                 `json:"Kundentyp_PRIVAT"`
	Kundentyp_LANDWIRT                                                              kundentyp.Kundentyp                                                                 `json:"Kundentyp_LANDWIRT"`
	Kundentyp_SONSTIGE                                                              kundentyp.Kundentyp                                                                 `json:"Kundentyp_SONSTIGE"`
	Kundentyp_HAUSHALT                                                              kundentyp.Kundentyp                                                                 `json:"Kundentyp_HAUSHALT"`
	Kundentyp_DIREKTHEIZUNG                                                         kundentyp.Kundentyp                                                                 `json:"Kundentyp_DIREKTHEIZUNG"`
	Kundentyp_GEMEINSCHAFT_MFH                                                      kundentyp.Kundentyp                                                                 `json:"Kundentyp_GEMEINSCHAFT_MFH"`
	Kundentyp_KIRCHE                                                                kundentyp.Kundentyp                                                                 `json:"Kundentyp_KIRCHE"`
	Kundentyp_KWK                                                                   kundentyp.Kundentyp                                                                 `json:"Kundentyp_KWK"`
	Kundentyp_LADESAEULE                                                            kundentyp.Kundentyp                                                                 `json:"Kundentyp_LADESAEULE"`
	Kundentyp_BELEUCHTUNG_OEFFENTLICH                                               kundentyp.Kundentyp                                                                 `json:"Kundentyp_BELEUCHTUNG_OEFFENTLICH"`
	Kundentyp_BELEUCHTUNG_STRASSE                                                   kundentyp.Kundentyp                                                                 `json:"Kundentyp_BELEUCHTUNG_STRASSE"`
	Kundentyp_SPEICHERHEIZUNG                                                       kundentyp.Kundentyp                                                                 `json:"Kundentyp_SPEICHERHEIZUNG"`
	Kundentyp_UNTERBR_EINRICHTUNG                                                   kundentyp.Kundentyp                                                                 `json:"Kundentyp_UNTERBR_EINRICHTUNG"`
	Kundentyp_WAERMEPUMPE                                                           kundentyp.Kundentyp                                                                 `json:"Kundentyp_WAERMEPUMPE"`
	Mengenoperator_KLEINER_ALS                                                      mengenoperator.Mengenoperator                                                       `json:"Mengenoperator_KLEINER_ALS"`
	Mengenoperator_GROESSER_ALS                                                     mengenoperator.Mengenoperator                                                       `json:"Mengenoperator_GROESSER_ALS"`
	Mengenoperator_GLEICH                                                           mengenoperator.Mengenoperator                                                       `json:"Mengenoperator_GLEICH"`
	Messpreistyp_MESSPREIS_G4                                                       messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_MESSPREIS_G4"`
	Messpreistyp_MESSPREIS_G6                                                       messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_MESSPREIS_G6"`
	Messpreistyp_MESSPREIS_G10                                                      messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_MESSPREIS_G10"`
	Messpreistyp_MESSPREIS_G16                                                      messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_MESSPREIS_G16"`
	Messpreistyp_MESSPREIS_G25                                                      messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_MESSPREIS_G25"`
	Messpreistyp_MESSPREIS_G40                                                      messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_MESSPREIS_G40"`
	Messpreistyp_ELEKTRONISCHER_AUFSATZ                                             messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_ELEKTRONISCHER_AUFSATZ"`
	Messpreistyp_SMART_METER_MESSPREIS_G2_5                                         messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_SMART_METER_MESSPREIS_G2_5"`
	Messpreistyp_SMART_METER_MESSPREIS_G4                                           messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_SMART_METER_MESSPREIS_G4"`
	Messpreistyp_SMART_METER_MESSPREIS_G6                                           messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_SMART_METER_MESSPREIS_G6"`
	Messpreistyp_SMART_METER_MESSPREIS_G10                                          messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_SMART_METER_MESSPREIS_G10"`
	Messpreistyp_SMART_METER_MESSPREIS_G16                                          messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_SMART_METER_MESSPREIS_G16"`
	Messpreistyp_SMART_METER_MESSPREIS_G25                                          messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_SMART_METER_MESSPREIS_G25"`
	Messpreistyp_SMART_METER_MESSPREIS_G40                                          messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_SMART_METER_MESSPREIS_G40"`
	Messpreistyp_VERRECHNUNGSPREIS_ET_WECHSEL                                       messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_VERRECHNUNGSPREIS_ET_WECHSEL"`
	Messpreistyp_VERRECHNUNGSPREIS_ET_DREH                                          messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_VERRECHNUNGSPREIS_ET_DREH"`
	Messpreistyp_VERRECHNUNGSPREIS_ZT_WECHSEL                                       messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_VERRECHNUNGSPREIS_ZT_WECHSEL"`
	Messpreistyp_VERRECHNUNGSPREIS_ZT_DREH                                          messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_VERRECHNUNGSPREIS_ZT_DREH"`
	Messpreistyp_VERRECHNUNGSPREIS_L_ET                                             messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_VERRECHNUNGSPREIS_L_ET"`
	Messpreistyp_VERRECHNUNGSPREIS_L_ZT                                             messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_VERRECHNUNGSPREIS_L_ZT"`
	Messpreistyp_VERRECHNUNGSPREIS_SM                                               messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_VERRECHNUNGSPREIS_SM"`
	Messpreistyp_AUFSCHLAG_WANDLER                                                  messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_AUFSCHLAG_WANDLER"`
	Messpreistyp_AUFSCHLAG_TARIFSCHALTUNG                                           messpreistyp.Messpreistyp                                                           `json:"Messpreistyp_AUFSCHLAG_TARIFSCHALTUNG"`
	Normierungsfaktor_NORMIERUNGSFAKTOR_1_000_000_KWH_A                              normierungsfaktor.Normierungsfaktor                                                 `json:"Normierungsfaktor_NORMIERUNGSFAKTOR_1_000_000_KWH_A"`
	Normierungsfaktor_NORMIERUNGSFAKTOR_300_KWH_K                                    normierungsfaktor.Normierungsfaktor                                                 `json:"Normierungsfaktor_NORMIERUNGSFAKTOR_300_KWH_K"`
	Normierungsfaktor_NORMIERUNGSFAKTOR_1_000_000_KW                                 normierungsfaktor.Normierungsfaktor                                                 `json:"Normierungsfaktor_NORMIERUNGSFAKTOR_1_000_000_KW"`
	Oekolabel_GASGREEN_GRUENER_STROM                                                oekolabel.Oekolabel                                                                 `json:"Oekolabel_GASGREEN_GRUENER_STROM"`
	Oekolabel_GASGREEN                                                              oekolabel.Oekolabel                                                                 `json:"Oekolabel_GASGREEN"`
	Oekolabel_GRUENER_STROM_GOLD                                                    oekolabel.Oekolabel                                                                 `json:"Oekolabel_GRUENER_STROM_GOLD"`
	Oekolabel_GRUENER_STROM_SILBER                                                  oekolabel.Oekolabel                                                                 `json:"Oekolabel_GRUENER_STROM_SILBER"`
	Oekolabel_GRUENER_STROM                                                         oekolabel.Oekolabel                                                                 `json:"Oekolabel_GRUENER_STROM"`
	Oekolabel_GRUENES_GAS                                                           oekolabel.Oekolabel                                                                 `json:"Oekolabel_GRUENES_GAS"`
	Oekolabel_NATURWATT_STROM                                                       oekolabel.Oekolabel                                                                 `json:"Oekolabel_NATURWATT_STROM"`
	Oekolabel_OK_POWER                                                              oekolabel.Oekolabel                                                                 `json:"Oekolabel_OK_POWER"`
	Oekolabel_RENEWABLE_PLUS                                                        oekolabel.Oekolabel                                                                 `json:"Oekolabel_RENEWABLE_PLUS"`
	Oekolabel_WATERGREEN                                                            oekolabel.Oekolabel                                                                 `json:"Oekolabel_WATERGREEN"`
	Oekolabel_WATERGREEN_PLUS                                                       oekolabel.Oekolabel                                                                 `json:"Oekolabel_WATERGREEN_PLUS"`
	Oekozertifikat_CMS_EE02                                                         oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_CMS_EE02"`
	Oekozertifikat_EECS                                                             oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_EECS"`
	Oekozertifikat_FRAUNHOFER                                                       oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_FRAUNHOFER"`
	Oekozertifikat_BET                                                              oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_BET"`
	Oekozertifikat_KLIMA_INVEST                                                     oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_KLIMA_INVEST"`
	Oekozertifikat_LGA                                                              oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_LGA"`
	Oekozertifikat_FREIBERG                                                         oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_FREIBERG"`
	Oekozertifikat_RECS                                                             oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_RECS"`
	Oekozertifikat_REGS_EGL                                                         oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_REGS_EGL"`
	Oekozertifikat_TUEV                                                             oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_TUEV"`
	Oekozertifikat_TUEV_HESSEN                                                      oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_TUEV_HESSEN"`
	Oekozertifikat_TUEV_NORD                                                        oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_TUEV_NORD"`
	Oekozertifikat_TUEV_RHEINLAND                                                   oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_TUEV_RHEINLAND"`
	Oekozertifikat_TUEV_SUED                                                        oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_TUEV_SUED"`
	Oekozertifikat_TUEV_SUED_EE01                                                   oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_TUEV_SUED_EE01"`
	Oekozertifikat_TUEV_SUED_EE02                                                   oekozertifikat.Oekozertifikat                                                       `json:"Oekozertifikat_TUEV_SUED_EE02"`
	Preisgarantietyp_ALLE_PREISBESTANDTEILE_NETTO                                   preisgarantietyp.Preisgarantietyp                                                   `json:"Preisgarantietyp_ALLE_PREISBESTANDTEILE_NETTO"`
	Preisgarantietyp_PREISBESTANDTEILE_OHNE_ABGABEN                                 preisgarantietyp.Preisgarantietyp                                                   `json:"Preisgarantietyp_PREISBESTANDTEILE_OHNE_ABGABEN"`
	Preisgarantietyp_NUR_ENERGIEPREIS                                               preisgarantietyp.Preisgarantietyp                                                   `json:"Preisgarantietyp_NUR_ENERGIEPREIS"`
	Preismodell_TRANCHE                                                             preismodell.Preismodell                                                             `json:"Preismodell_TRANCHE"`
	Preistyp_ARBEITSPREIS_EINTARIF                                                  preistyp.Preistyp                                                                   `json:"Preistyp_ARBEITSPREIS_EINTARIF"`
	Preistyp_ARBEITSPREIS_HT                                                        preistyp.Preistyp                                                                   `json:"Preistyp_ARBEITSPREIS_HT"`
	Preistyp_ARBEITSPREIS_NT                                                        preistyp.Preistyp                                                                   `json:"Preistyp_ARBEITSPREIS_NT"`
	Preistyp_LEISTUNGSPREIS                                                         preistyp.Preistyp                                                                   `json:"Preistyp_LEISTUNGSPREIS"`
	Preistyp_MESSPREIS                                                              preistyp.Preistyp                                                                   `json:"Preistyp_MESSPREIS"`
	Preistyp_ENTGELT_ABLESUNG                                                       preistyp.Preistyp                                                                   `json:"Preistyp_ENTGELT_ABLESUNG"`
	Preistyp_ENTGELT_ABRECHNUNG                                                     preistyp.Preistyp                                                                   `json:"Preistyp_ENTGELT_ABRECHNUNG"`
	Preistyp_ENTGELT_MSB                                                            preistyp.Preistyp                                                                   `json:"Preistyp_ENTGELT_MSB"`
	Preistyp_PROVISION                                                              preistyp.Preistyp                                                                   `json:"Preistyp_PROVISION"`
	Prioritaet_SEHR_NIEDRIG                                                         prioritaet.Prioritaet                                                               `json:"Prioritaet_SEHR_NIEDRIG"`
	Prioritaet_NIEDRIG                                                              prioritaet.Prioritaet                                                               `json:"Prioritaet_NIEDRIG"`
	Prioritaet_NORMAL                                                               prioritaet.Prioritaet                                                               `json:"Prioritaet_NORMAL"`
	Prioritaet_HOCH                                                                 prioritaet.Prioritaet                                                               `json:"Prioritaet_HOCH"`
	Prioritaet_SEHR_HOCH                                                            prioritaet.Prioritaet                                                               `json:"Prioritaet_SEHR_HOCH"`
	Profilklasse_GEWERBE                                                            profilklasse.Profilklasse                                                           `json:"Profilklasse_GEWERBE"`
	Profilklasse_GEWERBE_WERKTAG_9_BIS_18                                           profilklasse.Profilklasse                                                           `json:"Profilklasse_GEWERBE_WERKTAG_9_BIS_18"`
	Profilklasse_GEWERBE_ABEND                                                      profilklasse.Profilklasse                                                           `json:"Profilklasse_GEWERBE_ABEND"`
	Profilklasse_GEWERBE_DURCHLAUFEND                                               profilklasse.Profilklasse                                                           `json:"Profilklasse_GEWERBE_DURCHLAUFEND"`
	Profilklasse_GEWERBE_LADEN_FRISEUR                                              profilklasse.Profilklasse                                                           `json:"Profilklasse_GEWERBE_LADEN_FRISEUR"`
	Profilklasse_GEWERBE_BAECKEREI                                                  profilklasse.Profilklasse                                                           `json:"Profilklasse_GEWERBE_BAECKEREI"`
	Profilklasse_GEWERBE_WOCHENENDE                                                 profilklasse.Profilklasse                                                           `json:"Profilklasse_GEWERBE_WOCHENENDE"`
	Profilklasse_LANDWIRTSCHAFT                                                     profilklasse.Profilklasse                                                           `json:"Profilklasse_LANDWIRTSCHAFT"`
	Profilklasse_LANDWIRTSCHAFT_MIT_MILCH                                           profilklasse.Profilklasse                                                           `json:"Profilklasse_LANDWIRTSCHAFT_MIT_MILCH"`
	Profilklasse_LANDWIRTSCHAFT_OHNE_MILCH                                          profilklasse.Profilklasse                                                           `json:"Profilklasse_LANDWIRTSCHAFT_OHNE_MILCH"`
	Profilklasse_HAUSHALT                                                           profilklasse.Profilklasse                                                           `json:"Profilklasse_HAUSHALT"`
	Profilklasse_BANDLAST                                                           profilklasse.Profilklasse                                                           `json:"Profilklasse_BANDLAST"`
	Profilklasse_HEIZWAERMESPEICHER                                                 profilklasse.Profilklasse                                                           `json:"Profilklasse_HEIZWAERMESPEICHER"`
	Profilklasse_STRASSENBELEUCHTUNG                                                profilklasse.Profilklasse                                                           `json:"Profilklasse_STRASSENBELEUCHTUNG"`
	Profilklasse_PHOTOVOLTAIK                                                       profilklasse.Profilklasse                                                           `json:"Profilklasse_PHOTOVOLTAIK"`
	Profilklasse_BLOCKHEIZKRAFTWERK                                                 profilklasse.Profilklasse                                                           `json:"Profilklasse_BLOCKHEIZKRAFTWERK"`
	Profilklasse_SONSTIGE_ERZEUGENDE_MARKTLOKATION                                  profilklasse.Profilklasse                                                           `json:"Profilklasse_SONSTIGE_ERZEUGENDE_MARKTLOKATION"`
	Profilklasse_EMOB_OEFFENTLICH                                                   profilklasse.Profilklasse                                                           `json:"Profilklasse_EMOB_OEFFENTLICH"`
	Profilklasse_EMOB_HAUSHALT                                                      profilklasse.Profilklasse                                                           `json:"Profilklasse_EMOB_HAUSHALT"`
	Profilklasse_EMOB_GEWERBE                                                       profilklasse.Profilklasse                                                           `json:"Profilklasse_EMOB_GEWERBE"`
	Qualitaet_VOLLSTAENDIG                                                          qualitaet.Qualitaet                                                                 `json:"Qualitaet_VOLLSTAENDIG"`
	Qualitaet_INFORMATIV                                                            qualitaet.Qualitaet                                                                 `json:"Qualitaet_INFORMATIV"`
	Qualitaet_IM_SYSTEM_VORHANDEN                                                   qualitaet.Qualitaet                                                                 `json:"Qualitaet_IM_SYSTEM_VORHANDEN"`
	Qualitaet_ERWARTET                                                              qualitaet.Qualitaet                                                                 `json:"Qualitaet_ERWARTET"`
	Qualitaet_VORLAEUFIG                                                            qualitaet.Qualitaet                                                                 `json:"Qualitaet_VORLAEUFIG"`
	Qualitaet_UNVOLLSTAENDIG                                                        qualitaet.Qualitaet                                                                 `json:"Qualitaet_UNVOLLSTAENDIG"`
	Rechnungslegung_MONATSRECHN                                                     rechnungslegung.Rechnungslegung                                                     `json:"Rechnungslegung_MONATSRECHN"`
	Rechnungslegung_ABSCHL_MONATSRECHN                                              rechnungslegung.Rechnungslegung                                                     `json:"Rechnungslegung_ABSCHL_MONATSRECHN"`
	Rechnungslegung_ABSCHL_JAHRESRECHN                                              rechnungslegung.Rechnungslegung                                                     `json:"Rechnungslegung_ABSCHL_JAHRESRECHN"`
	Rechnungslegung_MONATSRECHN_JAHRESRECHN                                         rechnungslegung.Rechnungslegung                                                     `json:"Rechnungslegung_MONATSRECHN_JAHRESRECHN"`
	Rechnungslegung_VORKASSE                                                        rechnungslegung.Rechnungslegung                                                     `json:"Rechnungslegung_VORKASSE"`
	RechnungspositionsStatus_ROH                                                    rechnungspositionsstatus.RechnungspositionsStatus                                   `json:"RechnungspositionsStatus_ROH"`
	RechnungspositionsStatus_ROH_AUSGENOMMEN                                        rechnungspositionsstatus.RechnungspositionsStatus                                   `json:"RechnungspositionsStatus_ROH_AUSGENOMMEN"`
	RechnungspositionsStatus_ABRECHENBAR                                            rechnungspositionsstatus.RechnungspositionsStatus                                   `json:"RechnungspositionsStatus_ABRECHENBAR"`
	RechnungspositionsStatus_ABRECHENBAR_AUSGENOMMEN                                rechnungspositionsstatus.RechnungspositionsStatus                                   `json:"RechnungspositionsStatus_ABRECHENBAR_AUSGENOMMEN"`
	RechnungspositionsStatus_ABGERECHNET                                            rechnungspositionsstatus.RechnungspositionsStatus                                   `json:"RechnungspositionsStatus_ABGERECHNET"`
	Regionskriteriumtyp_BUNDESLANDKENNZIFFER                                        regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_BUNDESLANDKENNZIFFER"`
	Regionskriteriumtyp_BUNDESLAND_NAME                                             regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_BUNDESLAND_NAME"`
	Regionskriteriumtyp_MARKTGEBIET_NUMMER                                          regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_MARKTGEBIET_NUMMER"`
	Regionskriteriumtyp_MARKTGEBIET_NAME                                            regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_MARKTGEBIET_NAME"`
	Regionskriteriumtyp_REGELGEBIET_NUMMER                                          regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_REGELGEBIET_NUMMER"`
	Regionskriteriumtyp_REGELGEBIET_NAME                                            regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_REGELGEBIET_NAME"`
	Regionskriteriumtyp_NETZBETREIBER_NUMMER                                        regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_NETZBETREIBER_NUMMER"`
	Regionskriteriumtyp_NETZBETREIBER_NAME                                          regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_NETZBETREIBER_NAME"`
	Regionskriteriumtyp_BILANZIERUNGS_GEBIET_NUMMER                                 regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_BILANZIERUNGS_GEBIET_NUMMER"`
	Regionskriteriumtyp_MSB_NUMMER                                                  regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_MSB_NUMMER"`
	Regionskriteriumtyp_MSB_NAME                                                    regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_MSB_NAME"`
	Regionskriteriumtyp_VERSORGER_NUMMER                                            regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_VERSORGER_NUMMER"`
	Regionskriteriumtyp_VERSORGER_NAME                                              regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_VERSORGER_NAME"`
	Regionskriteriumtyp_GRUNDVERSORGER_NUMMER                                       regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_GRUNDVERSORGER_NUMMER"`
	Regionskriteriumtyp_GRUNDVERSORGER_NAME                                         regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_GRUNDVERSORGER_NAME"`
	Regionskriteriumtyp_KREIS_NAME                                                  regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_KREIS_NAME"`
	Regionskriteriumtyp_KREISKENNZIFFER                                             regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_KREISKENNZIFFER"`
	Regionskriteriumtyp_GEMEINDE_NAME                                               regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_GEMEINDE_NAME"`
	Regionskriteriumtyp_GEMEINDEKENNZIFFER                                          regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_GEMEINDEKENNZIFFER"`
	Regionskriteriumtyp_POSTLEITZAHL                                                regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_POSTLEITZAHL"`
	Regionskriteriumtyp_ORT                                                         regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_ORT"`
	Regionskriteriumtyp_EINWOHNERZAHL_GEMEINDE                                      regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_EINWOHNERZAHL_GEMEINDE"`
	Regionskriteriumtyp_EINWOHNERZAHL_ORT                                           regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_EINWOHNERZAHL_ORT"`
	Regionskriteriumtyp_KM_UMKREIS                                                  regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_KM_UMKREIS"`
	Regionskriteriumtyp_BUNDESWEIT                                                  regionskriteriumtyp.Regionskriteriumtyp                                             `json:"Regionskriteriumtyp_BUNDESWEIT"`
	Schalthandlung_LEISTUNG_LOKATION_AN                                             schalthandlung.Schalthandlung                                                       `json:"Schalthandlung_LEISTUNG_LOKATION_AN"`
	Schalthandlung_LEISTUNG_LOKATION_AUS                                            schalthandlung.Schalthandlung                                                       `json:"Schalthandlung_LEISTUNG_LOKATION_AUS"`
	Schwachlastfaehig_NICHT_SCHWACHLASTFAEHIG                                       schwachlastfaehig.Schwachlastfaehig                                                 `json:"Schwachlastfaehig_NICHT_SCHWACHLASTFAEHIG"`
	Schwachlastfaehig_SCHWACHLASTFAEHIG                                             schwachlastfaehig.Schwachlastfaehig                                                 `json:"Schwachlastfaehig_SCHWACHLASTFAEHIG"`
	Servicetyp_STROM_NB                                                             servicetyp.Servicetyp                                                               `json:"Servicetyp_STROM_NB"`
	Servicetyp_STROM_MSB                                                            servicetyp.Servicetyp                                                               `json:"Servicetyp_STROM_MSB"`
	Servicetyp_STROM_LIEF                                                           servicetyp.Servicetyp                                                               `json:"Servicetyp_STROM_LIEF"`
	Servicetyp_GAS_NB                                                               servicetyp.Servicetyp                                                               `json:"Servicetyp_GAS_NB"`
	Servicetyp_GAS_MSB                                                              servicetyp.Servicetyp                                                               `json:"Servicetyp_GAS_MSB"`
	Servicetyp_GAS_LIEF                                                             servicetyp.Servicetyp                                                               `json:"Servicetyp_GAS_LIEF"`
	Sperrauftragsablehngrund_DUPLIKAT                                               sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_DUPLIKAT"`
	Sperrauftragsablehngrund_FALSCHER_MSB                                           sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_FALSCHER_MSB"`
	Sperrauftragsablehngrund_FALSCHE_SPANNUNGSEBENE                                 sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_FALSCHE_SPANNUNGSEBENE"`
	Sperrauftragsablehngrund_WEITERE_MALO_BETROFFEN                                 sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_WEITERE_MALO_BETROFFEN"`
	Sperrauftragsablehngrund_ANDERER_ABLEHNGRUND                                    sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_ANDERER_ABLEHNGRUND"`
	Sperrauftragsablehngrund_FRISTVERLETZUNG_TERMINGEBUNDEN                         sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_FRISTVERLETZUNG_TERMINGEBUNDEN"`
	Sperrauftragsablehngrund_ANDERER_FEHLER                                         sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_ANDERER_FEHLER"`
	Sperrauftragsablehngrund_LIEGT_BEREITS_VOR                                      sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_LIEGT_BEREITS_VOR"`
	Sperrauftragsablehngrund_ANDERER_ZUKUENFTIGER_LIEFERANT                         sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_ANDERER_ZUKUENFTIGER_LIEFERANT"`
	Sperrauftragsablehngrund_BESTAETIGTER_LIEFERBEGINN                              sperrauftragsablehngrund.Sperrauftragsablehngrund                                   `json:"Sperrauftragsablehngrund_BESTAETIGTER_LIEFERBEGINN"`
	Sperrauftragsart_SPERREN                                                        sperrauftragsart.Sperrauftragsart                                                   `json:"Sperrauftragsart_SPERREN"`
	Sperrauftragsart_ENTSPERREN                                                     sperrauftragsart.Sperrauftragsart                                                   `json:"Sperrauftragsart_ENTSPERREN"`
	Sperrauftragsverhinderungsgrund_RECHTLICHER_GRUND_FEHLT                         sperrauftragsverhinderungsgrund.Sperrauftragsverhinderungsgrund                     `json:"Sperrauftragsverhinderungsgrund_RECHTLICHER_GRUND_FEHLT"`
	Sperrauftragsverhinderungsgrund_AKTIVE_ZUTRITTSVERWEIGERUNG                     sperrauftragsverhinderungsgrund.Sperrauftragsverhinderungsgrund                     `json:"Sperrauftragsverhinderungsgrund_AKTIVE_ZUTRITTSVERWEIGERUNG"`
	Sperrauftragsverhinderungsgrund_PASSIVE_ZUTRITTSVERWEIGERUNG                    sperrauftragsverhinderungsgrund.Sperrauftragsverhinderungsgrund                     `json:"Sperrauftragsverhinderungsgrund_PASSIVE_ZUTRITTSVERWEIGERUNG"`
	Sperrauftragsverhinderungsgrund_ANDERER_VERHINDERUNGSGRUND                      sperrauftragsverhinderungsgrund.Sperrauftragsverhinderungsgrund                     `json:"Sperrauftragsverhinderungsgrund_ANDERER_VERHINDERUNGSGRUND"`
	Sperrauftragsverhinderungsgrund_TATSAECHLICHER_VERHINDERUNGSGRUND               sperrauftragsverhinderungsgrund.Sperrauftragsverhinderungsgrund                     `json:"Sperrauftragsverhinderungsgrund_TATSAECHLICHER_VERHINDERUNGSGRUND"`
	Sperrauftragsverhinderungsgrund_TECHNISCHER_VERHINDERUNGSGRUND                  sperrauftragsverhinderungsgrund.Sperrauftragsverhinderungsgrund                     `json:"Sperrauftragsverhinderungsgrund_TECHNISCHER_VERHINDERUNGSGRUND"`
	Tarifkalkulationsmethode_STAFFELN                                               tarifkalkulationsmethode.Tarifkalkulationsmethode                                   `json:"Tarifkalkulationsmethode_STAFFELN"`
	Tarifkalkulationsmethode_ZONEN                                                  tarifkalkulationsmethode.Tarifkalkulationsmethode                                   `json:"Tarifkalkulationsmethode_ZONEN"`
	Tarifkalkulationsmethode_BESTABRECHNUNG_STAFFEL                                 tarifkalkulationsmethode.Tarifkalkulationsmethode                                   `json:"Tarifkalkulationsmethode_BESTABRECHNUNG_STAFFEL"`
	Tarifkalkulationsmethode_PAKETPREIS                                             tarifkalkulationsmethode.Tarifkalkulationsmethode                                   `json:"Tarifkalkulationsmethode_PAKETPREIS"`
	Tarifmerkmal_VORKASSE                                                           tarifmerkmal.Tarifmerkmal                                                           `json:"Tarifmerkmal_VORKASSE"`
	Tarifmerkmal_PAKET                                                              tarifmerkmal.Tarifmerkmal                                                           `json:"Tarifmerkmal_PAKET"`
	Tarifmerkmal_KOMBI                                                              tarifmerkmal.Tarifmerkmal                                                           `json:"Tarifmerkmal_KOMBI"`
	Tarifmerkmal_FESTPREIS                                                          tarifmerkmal.Tarifmerkmal                                                           `json:"Tarifmerkmal_FESTPREIS"`
	Tarifmerkmal_BAUSTROM                                                           tarifmerkmal.Tarifmerkmal                                                           `json:"Tarifmerkmal_BAUSTROM"`
	Tarifmerkmal_HAUSLICHT                                                          tarifmerkmal.Tarifmerkmal                                                           `json:"Tarifmerkmal_HAUSLICHT"`
	Tarifmerkmal_HEIZSTROM                                                          tarifmerkmal.Tarifmerkmal                                                           `json:"Tarifmerkmal_HEIZSTROM"`
	Tarifregionskriterium_NETZ_NUMMER                                               tarifregionskriterium.Tarifregionskriterium                                         `json:"Tarifregionskriterium_NETZ_NUMMER"`
	Tarifregionskriterium_POSTLEITZAHL                                              tarifregionskriterium.Tarifregionskriterium                                         `json:"Tarifregionskriterium_POSTLEITZAHL"`
	Tarifregionskriterium_ORT                                                       tarifregionskriterium.Tarifregionskriterium                                         `json:"Tarifregionskriterium_ORT"`
	Tarifregionskriterium_GRUNDVERSORGER_NUMMER                                     tarifregionskriterium.Tarifregionskriterium                                         `json:"Tarifregionskriterium_GRUNDVERSORGER_NUMMER"`
	Tariftyp_GRUND_ERSATZVERSORGUNG                                                 tariftyp.Tariftyp                                                                   `json:"Tariftyp_GRUND_ERSATZVERSORGUNG"`
	Tariftyp_GRUNDVERSORGUNG                                                        tariftyp.Tariftyp                                                                   `json:"Tariftyp_GRUNDVERSORGUNG"`
	Tariftyp_ERSATZVERSORGUNG                                                       tariftyp.Tariftyp                                                                   `json:"Tariftyp_ERSATZVERSORGUNG"`
	Tariftyp_SONDERTARIF                                                            tariftyp.Tariftyp                                                                   `json:"Tariftyp_SONDERTARIF"`
	Verbrauchsmengetyp_ARBEITLEISTUNGTAGESPARAMETERABHMALO                          verbrauchsmengetyp.Verbrauchsmengetyp                                               `json:"Verbrauchsmengetyp_ARBEITLEISTUNGTAGESPARAMETERABHMALO"`
	Verbrauchsmengetyp_VERANSCHLAGTEJAHRESMENGE                                     verbrauchsmengetyp.Verbrauchsmengetyp                                               `json:"Verbrauchsmengetyp_VERANSCHLAGTEJAHRESMENGE"`
	Verbrauchsmengetyp_TUMKUNDENWERT                                                verbrauchsmengetyp.Verbrauchsmengetyp                                               `json:"Verbrauchsmengetyp_TUMKUNDENWERT"`
	Vertragsform_ONLINE                                                             vertragsform.Vertragsform                                                           `json:"Vertragsform_ONLINE"`
	Vertragsform_DIREKT                                                             vertragsform.Vertragsform                                                           `json:"Vertragsform_DIREKT"`
	Vertragsform_FAX                                                                vertragsform.Vertragsform                                                           `json:"Vertragsform_FAX"`
	Vertragstatus_IN_ARBEIT                                                         vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_IN_ARBEIT"`
	Vertragstatus_UEBERMITTELT                                                      vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_UEBERMITTELT"`
	Vertragstatus_ANGENOMMEN                                                        vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_ANGENOMMEN"`
	Vertragstatus_AKTIV                                                             vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_AKTIV"`
	Vertragstatus_ABGELEHNT                                                         vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_ABGELEHNT"`
	Vertragstatus_WIDERRUFEN                                                        vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_WIDERRUFEN"`
	Vertragstatus_STORNIERT                                                         vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_STORNIERT"`
	Vertragstatus_GEKUENDIGT                                                        vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_GEKUENDIGT"`
	Vertragstatus_BEENDET                                                           vertragstatus.Vertragstatus                                                         `json:"Vertragstatus_BEENDET"`
	Voraussetzungen_EINZUGSERMAECHTIGUNG                                            voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_EINZUGSERMAECHTIGUNG"`
	Voraussetzungen_ZEITPUNKT                                                       voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ZEITPUNKT"`
	Voraussetzungen_LIEFERANBINDUNG_EINE                                            voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_LIEFERANBINDUNG_EINE"`
	Voraussetzungen_LIEFERANBINDUNG_ALLE                                            voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_LIEFERANBINDUNG_ALLE"`
	Voraussetzungen_GEWERBE                                                         voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_GEWERBE"`
	Voraussetzungen_LASTPROFIL                                                      voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_LASTPROFIL"`
	Voraussetzungen_ZAEHLERTYP_GROESSE                                              voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ZAEHLERTYP_GROESSE"`
	Voraussetzungen_AUSSCHLUSS_GROSSVERBRAUCHER                                     voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_AUSSCHLUSS_GROSSVERBRAUCHER"`
	Voraussetzungen_NEUKUNDE                                                        voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_NEUKUNDE"`
	Voraussetzungen_BESTIMMTE_VERTRAGSFORMALITAETEN                                 voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_BESTIMMTE_VERTRAGSFORMALITAETEN"`
	Voraussetzungen_SELBSTABLESUNG                                                  voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_SELBSTABLESUNG"`
	Voraussetzungen_ONLINEVORAUSSETZUNG                                             voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ONLINEVORAUSSETZUNG"`
	Voraussetzungen_MINDESTUMSATZ                                                   voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_MINDESTUMSATZ"`
	Voraussetzungen_ZUSATZPRODUKT                                                   voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ZUSATZPRODUKT"`
	Voraussetzungen_NEUKUNDE_MIT_VORAUSSETZUNGEN                                    voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_NEUKUNDE_MIT_VORAUSSETZUNGEN"`
	Voraussetzungen_DIREKTVERTRIEB                                                  voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_DIREKTVERTRIEB"`
	Voraussetzungen_ANSCHLUSSART                                                    voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ANSCHLUSSART"`
	Voraussetzungen_ANSCHLUSSWERT                                                   voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ANSCHLUSSWERT"`
	Voraussetzungen_ALTER_KUNDENANLAGE                                              voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ALTER_KUNDENANLAGE"`
	Voraussetzungen_ANLAGEBESCHAFFENHEIT                                            voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ANLAGEBESCHAFFENHEIT"`
	Voraussetzungen_BETRIEBSSTUNDENBEGRENZUNG                                       voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_BETRIEBSSTUNDENBEGRENZUNG"`
	Voraussetzungen_FREIGABEZEITEN                                                  voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_FREIGABEZEITEN"`
	Voraussetzungen_FAMILIENSTRUKTUR                                                voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_FAMILIENSTRUKTUR"`
	Voraussetzungen_MITGLIEDSCHAFT                                                  voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_MITGLIEDSCHAFT"`
	Voraussetzungen_STAATLICHE_FOERDERUNG                                           voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_STAATLICHE_FOERDERUNG"`
	Voraussetzungen_BESONDERE_VERBRAUCHSSTELLE                                      voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_BESONDERE_VERBRAUCHSSTELLE"`
	Voraussetzungen_NIEDRIGENERGIE                                                  voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_NIEDRIGENERGIE"`
	Voraussetzungen_ORTSTEILE_LIEFERGEBIET                                          voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ORTSTEILE_LIEFERGEBIET"`
	Voraussetzungen_WAERMEBEDARF_ERDGAS                                             voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_WAERMEBEDARF_ERDGAS"`
	Voraussetzungen_MAX_ZAEHLER_LIEFERSTELLEN                                       voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_MAX_ZAEHLER_LIEFERSTELLEN"`
	Voraussetzungen_LIEFERUNGSBESCHRAENKUNG_GASART                                  voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_LIEFERUNGSBESCHRAENKUNG_GASART"`
	Voraussetzungen_KOMBI_BONI                                                      voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_KOMBI_BONI"`
	Voraussetzungen_ALTVERTRAG                                                      voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ALTVERTRAG"`
	Voraussetzungen_VORGESCHRIEBENE_ZUSATZANLAGE                                    voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_VORGESCHRIEBENE_ZUSATZANLAGE"`
	Voraussetzungen_MEHRERE_ZAEHLER_ABNAHMESTELLEN                                  voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_MEHRERE_ZAEHLER_ABNAHMESTELLEN"`
	Voraussetzungen_BESTIMMTER_ABNAHMEFALL                                          voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_BESTIMMTER_ABNAHMEFALL"`
	Voraussetzungen_ZUSATZMODALITAET                                                voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_ZUSATZMODALITAET"`
	Voraussetzungen_NACHWEIS_ZAHLUNGSFAEHIGKEIT                                     voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_NACHWEIS_ZAHLUNGSFAEHIGKEIT"`
	Voraussetzungen_UMSTELLUNG_ENERGIEART                                           voraussetzungen.Voraussetzungen                                                     `json:"Voraussetzungen_UMSTELLUNG_ENERGIEART"`
	ZaehlertypSpezifikation_EDL40                                                   zaehlertypspezifikation.ZaehlertypSpezifikation                                     `json:"ZaehlertypSpezifikation_EDL40"`
	ZaehlertypSpezifikation_EDL21                                                   zaehlertypspezifikation.ZaehlertypSpezifikation                                     `json:"ZaehlertypSpezifikation_EDL21"`
	ZaehlertypSpezifikation_SONSTIGER_EHZ                                           zaehlertypspezifikation.ZaehlertypSpezifikation                                     `json:"ZaehlertypSpezifikation_SONSTIGER_EHZ"`
	ZaehlertypSpezifikation_MME_STANDARD                                            zaehlertypspezifikation.ZaehlertypSpezifikation                                     `json:"ZaehlertypSpezifikation_MME_STANDARD"`
	ZaehlertypSpezifikation_MME_MEDA                                                zaehlertypspezifikation.ZaehlertypSpezifikation                                     `json:"ZaehlertypSpezifikation_MME_MEDA"`
}

// TestCSharpEnumCompatibility verifies that enum values serialized by the C# BO4E-dotnet
// library can be correctly deserialized by the Go go-bo4e library.
//
// This test requires running the C# script first to generate csharp_enum_output.json.
// The GitHub Actions workflow handles this automatically.
func TestCSharpEnumCompatibility(t *testing.T) {
	data, err := os.ReadFile("csharp_enum_output.json")
	if err != nil {
		if os.IsNotExist(err) {
			t.Skip("Skipping test: csharp_enum_output.json not found. Run 'dotnet script serialize_enums.csx > csharp_enum_output.json' first.")
		}
		t.Fatalf("Error reading csharp_enum_output.json: %v", err)
	}

	var result AllEnumsTest
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Error deserializing JSON into AllEnumsTest struct: %v", err)
	}

	t.Log("SUCCESS: All enum values deserialized correctly!")
}
