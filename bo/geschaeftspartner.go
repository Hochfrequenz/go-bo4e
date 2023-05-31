package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/kontaktart"
)

// Geschaeftspartner models business partners, both companies and private persons.
type Geschaeftspartner struct {
	Geschaeftsobjekt
	Anrede                  *anrede.Anrede                                  `json:"anrede,omitempty"`                                                               // Die Anrede für den Geschaeftspartner, Z.B. HERR
	Name1                   string                                          `json:"name1,omitempty" validate:"required" example:"Yellow Strom GmbH,Hagen"`          // Name1 ist der erste Teil des Namens. Hier kann der Firmenname oder bei Privatpersonen beispielsweise der Nachname dargestellt werden.
	Name2                   string                                          `json:"name2,omitempty" example:"'Bereich Süd','Nina'"`                                 // Name2 ist der zweite Teil des Namens. Hier kann der eine Erweiterung zum Firmennamen oder bei Privatpersonen beispielsweise der Vorname dargestellt werden.
	Name3                   string                                          `json:"name3,omitempty" example:"Afrika,Sängerin"`                                      // Name3 enthält weitere Ergänzungen zum Firmennamen oder bei Privatpersonen Zusätze zum Namen dargestellt werden
	Gewerbekennzeichnung    bool                                            `json:"gewerbekennzeichnung" validate:"omitempty,required"`                             // Gewerbekennzeichnung markiert, ob es sich um einen Gewerbe/Unternehmen (true) oder eine Privatperson handelt (false)
	HrNummer                string                                          `json:"hrnummer,omitempty"`                                                             // HrNummer ist die Handelsregisternummer des Geschäftspartners
	Amtsgericht             string                                          `json:"amtsgericht,omitempty"`                                                          // Amtsgericht bzw Handelsregistergericht, das die Handelsregisternummer herausgegeben hat
	Kontaktwege             []kontaktart.Kontaktart                         `json:"kontaktweg,omitempty"`                                                           // Kontaktwege sind die bevorzugten Kontaktwege des Geschäftspartners (todo: rename json tag, too)
	UmsatzsteuerId          string                                          `json:"umsatzsteuerId,omitempty" example:"DE 813281825"`                                // UmsatzsteuerId ist die Steuer-ID des Geschäftspartners
	GlaeubigerId            string                                          `json:"glaeubigerId,omitempty" example:"DE 47116789"`                                   // GlaeubigerId ist die Gläubiger-ID, welche im Zahlungsverkehr verwendet wird
	EMailAdresse            string                                          `json:"eMailAdresse,omitempty" validate:"omitempty,email" example:"info@mp-energie.de"` // EMailAdresse ist die E-Mail-Adresse des Ansprechpartners
	Website                 string                                          `json:"website,omitempty" validate:"omitempty,url" example:"https://www.mp-energie.de"` // Website ist die Internetseite des Marktpartners
	Geschaeftspartnerrollen []geschaeftspartnerrolle.Geschaeftspartnerrolle `json:"geschaeftspartnerrollen,omitempty" validate:"required,min=1"`                    // Geschaeftspartnerrollen sind die Rollen, die der Geschäftspartner hat
	Partneradresse          *com.Adresse                                    `json:"partneradresse,omitempty" validate:"required"`                                   // Partneradresse ist die Adresse des Geschäftspartners, an der sich der Hauptsitz befindet.
}

func (_ Geschaeftspartner) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
