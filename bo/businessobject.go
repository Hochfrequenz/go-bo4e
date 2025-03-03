package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
	"reflect"
)

// BusinessObject is the interface that all Business Objects implement.
type BusinessObject interface {
	// GetBoTyp returns the Geschaeftsobjekt.BoTyp
	GetBoTyp() botyp.BOTyp
	// GetVersionStruktur returns the Geschaeftsobjekt.VersionStruktur
	GetVersionStruktur() string
	// GetGueltigkeitszeitraum returns the Geschaeftsobjekt.Gueltigkeitszeitraum
	GetGueltigkeitszeitraum() *com.Zeitraum
}

// Geschaeftsobjekt is the common base struct of all Business Objects
type Geschaeftsobjekt struct {
	unmappeddatamarshaller.ExtensionData
	BoTyp                botyp.BOTyp           `json:"boTyp" validate:"required"`           // BoTyp is the type of business object, may be used as discriminator
	VersionStruktur      string                `json:"versionStruktur" validate:"required"` // VersionStruktur is the version of BO4E used
	ExterneReferenzen    []com.ExterneReferenz `json:"externeReferenzen,omitempty"`         // ExterneReferenzen are external references of this object in various systems
	Gueltigkeitszeitraum *com.Zeitraum         `json:"gueltigkeitszeitraum,omitempty"`      // Defines the validity of a business object in terms of time (maybe multiple versions exist).
}

func (gob Geschaeftsobjekt) GetBoTyp() botyp.BOTyp {
	return gob.BoTyp
}

func (gob Geschaeftsobjekt) GetVersionStruktur() string {
	if gob.VersionStruktur == "" {
		return defaultVersionStruktur
	}
	return gob.VersionStruktur
}

func (gob Geschaeftsobjekt) GetGueltigkeitszeitraum() *com.Zeitraum {
	return gob.Gueltigkeitszeitraum
}

func newBoTypAndVersion[T BusinessObject](typ botyp.BOTyp) *T {
	a := new(T)
	gob := Geschaeftsobjekt{
		BoTyp:           typ,
		VersionStruktur: defaultVersionStruktur,
	}
	val := reflect.ValueOf(a)
	if val.CanInterface() {
		val = val.Elem()
	}
	field := val.FieldByName("Geschaeftsobjekt")
	if field.CanSet() {
		field.Set(reflect.ValueOf(gob))
	}
	return a
}
func setBoTypAndVersion[T BusinessObject](in T, typ botyp.BOTyp) T {
	a := new(T)
	gob := Geschaeftsobjekt{
		BoTyp:           typ,
		VersionStruktur: defaultVersionStruktur,
	}
	field := reflect.ValueOf(a).Elem().FieldByName("Geschaeftsobjekt")
	if field.CanSet() {
		field.Set(reflect.ValueOf(gob))
	}
	return *a
}

// defaultVersionStruktur is the Geschaeftsobjekt.VersionStruktur that a NewBusinessObject has by default
const defaultVersionStruktur = "1.1" // because Geschaeftsobjekt.ExterneReferenzen (plural) was introduced in v1.1. In v1.0 the name was "externeReferenz" (singular)

// NewBusinessObject creates an empty BusinessObject based on the provided type; Returns nil if the type is not implemented.
func NewBusinessObject(typ botyp.BOTyp) BusinessObject {
	var bo BusinessObject
	switch typ {
	case botyp.ANFRAGE:
		bo = newBoTypAndVersion[Anfrage](typ)
	case botyp.ANGEBOT:
		bo = newBoTypAndVersion[Angebot](typ)
	case botyp.ANSPRECHPARTNER:
		bo = newBoTypAndVersion[Ansprechpartner](typ)
	case botyp.BILANZIERUNG:
		bo = newBoTypAndVersion[Bilanzierung](typ)
	case botyp.EINSPEISUNG:
		bo = newBoTypAndVersion[Einspeisung](typ)
	case botyp.ENERGIEMENGE:
		bo = newBoTypAndVersion[Energiemenge](typ)
	case botyp.GESCHAEFTSPARTNER:
		bo = newBoTypAndVersion[Geschaeftspartner](typ)
	case botyp.LASTGANG:
		bo = newBoTypAndVersion[Lastgang](typ)
	case botyp.LOKATIONSZUORDNUNG:
		bo = newBoTypAndVersion[Lokationszuordnung](typ)
	case botyp.MARKTLOKATION:
		bo = newBoTypAndVersion[Marktlokation](typ)
	case botyp.MARKTTEILNEHMER:
		bo = newBoTypAndVersion[Marktteilnehmer](typ)
	case botyp.MESSLOKATION:
		bo = newBoTypAndVersion[Messlokation](typ)
	case botyp.NETZLOKATION:
		bo = newBoTypAndVersion[Netzlokation](typ)
	case botyp.NETZNUTZUNGSRECHNUNG:
		bo = newBoTypAndVersion[Netznutzungsrechnung](typ)
		bo.(*Netznutzungsrechnung).Rechnungsempfaenger = *newBoTypAndVersion[Geschaeftspartner](botyp.GESCHAEFTSPARTNER)
		bo.(*Netznutzungsrechnung).Rechnungsersteller = *newBoTypAndVersion[Geschaeftspartner](botyp.GESCHAEFTSPARTNER)
	case botyp.PREISBLATT:
		bo = newBoTypAndVersion[Preisblatt](typ)
		bo.(*Preisblatt).Herausgeber = *NewBusinessObject(botyp.MARKTTEILNEHMER).(*Marktteilnehmer)
	case botyp.PRODUKTPAKET:
		bo = newBoTypAndVersion[Produktpaket](typ)
	case botyp.RECHNUNG:
		bo = newBoTypAndVersion[Rechnung](typ)
		bo.(*Rechnung).Rechnungsempfaenger = *newBoTypAndVersion[Geschaeftspartner](botyp.GESCHAEFTSPARTNER)
		bo.(*Rechnung).Rechnungsersteller = *newBoTypAndVersion[Geschaeftspartner](botyp.GESCHAEFTSPARTNER)
	case botyp.REKLAMATION:
		bo = newBoTypAndVersion[Reklamation](typ)
	case botyp.VERTRAG:
		bo = newBoTypAndVersion[Vertrag](typ)
	case botyp.ZAEHLER:
		bo = newBoTypAndVersion[Zaehler](typ)
	case botyp.HANDELSUNSTIMMIGKEIT:
		bo = newBoTypAndVersion[Handelsunstimmigkeit](typ)
	case botyp.AVIS:
		bo = newBoTypAndVersion[Avis](typ)
	case botyp.STATUSBERICHT:
		bo = newBoTypAndVersion[Statusbericht](typ)
	case botyp.STEUERBARERESSOURCE:
		bo = newBoTypAndVersion[SteuerbareRessource](typ)
	case botyp.SUMMENZEITREIHE:
		bo = newBoTypAndVersion[Summenzeitreihe](typ)
	case botyp.TECHNISCHERESSOURCE:
		bo = newBoTypAndVersion[TechnischeRessource](typ)
	case botyp.TRANCHE:
		bo = newBoTypAndVersion[Tranche](typ)
	case botyp.ZAEHLZEITDEFINITION:
		bo = newBoTypAndVersion[Zaehlzeitdefinition](typ)
	case botyp.PREISBLATTUMLAGEN:
		bo = newBoTypAndVersion[PreisblattUmlagen](typ)
	case botyp.PREISBLATTMESSUNG:
		bo = newBoTypAndVersion[PreisblattMessung](typ)
	case botyp.PREISBLATTKONZESSIONSABGABE:
		bo = newBoTypAndVersion[PreisblattKonzessionsabgabe](typ)
	case botyp.AUSSCHREIBUNG, botyp.KOSTEN, botyp.PREISBLATTDIENSTLEISTUNG, botyp.TARIFINFO, botyp.TARIFPREISBLATT, botyp.ZEITREIHE:
		// TODO: diese Typen könnten noch in BO4E dotnet und hier implementiert oder im Standard deprecated werden
		return nil
	default:
		return nil
	}
	return bo
}
