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
	Anrede                  anrede.Anrede                                   `json:"anrede"`                                                                         // Die Anrede für den GePa, Z.B. HERR
	Name1                   string                                          `json:"name1" validate:"required" example:"Yellow Strom GmbH,Hagen"`                    // Erster Teil des Namens. Hier kann der Firmenname oder bei Privatpersonen beispielsweise der Nachname dargestellt werden.
	Name2                   string                                          `json:"name2,omitempty" example:"'Bereich Süd','Nina'"`                                 // Zweiter Teil des Namens. Hier kann der eine Erweiterung zum Firmennamen oder bei Privatpersonen beispielsweise der Vorname dargestellt werden.
	Name3                   string                                          `json:"name3,omitempty" example:"Afrika,Sängerin"`                                      // Hier können weitere Ergänzungen zum Firmennamen oder bei Privatpersonen Zusätze zum Namen dargestellt werden
	Gewerbekennzeichnung    bool                                            `json:"gewerbekennzeichnung" validate:"omitempty,required"`                             // Kennzeichnung, ob es sich um einen Gewerbe/Unternehmen (true) oder eine Privatperson handelt (false)
	HrNummer                string                                          `json:"hrnummer,omitempty"`                                                             // Handelsregisternummer des Geschäftspartners
	Amtsgericht             string                                          `json:"amtsgericht,omitempty"`                                                          // Amtsgericht bzw Handelsregistergericht, das die Handelsregisternummer herausgegeben hat
	Kontaktweg              kontaktart.Kontaktart                           `json:"kontaktweg,omitempty"`                                                           // Bevorzugter Kontaktweg des Geschäftspartners.
	UmsatzsteuerId          string                                          `json:"umsatzsteuerId,omitempty" example:"DE 813281825"`                                // Die Steuer-ID des Geschäftspartners
	GlaeubigerId            string                                          `json:"glaeubigerId,omitempty" example:"DE 47116789"`                                   // Die Gläubiger-ID welche im Zahlungsverkehr verwendet wird
	EMailAdresse            string                                          `json:"eMailAdresse,omitempty" validate:"omitempty,email" example:"info@mp-energie.de"` // E-Mail-Adresse des Ansprechpartners
	Website                 string                                          `json:"website,omitempty" validate:"omitempty,url" example:"https://www.mp-energie.de"` // Internetseite des Marktpartners
	Geschaeftspartnerrollen []geschaeftspartnerrolle.Geschaeftspartnerrolle `json:"geschaeftspartnerrollen" validate:"required,min=1"`                              // Rollen die der Geschäftspartner hat
	Partneradresse          com.Adresse                                     `json:"partneradresse" validate:"required"`                                             // Adresse des Geschäftspartners, an der sich der Hauptsitz befindet.
}
