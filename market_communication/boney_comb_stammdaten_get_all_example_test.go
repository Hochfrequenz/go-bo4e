package market_communication_test

import (
	"fmt"

	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

func ExampleGetAll() {
	boneyComb := &market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{
			bo.NewBusinessObject(botyp.MARKTTEILNEHMER),
			bo.NewBusinessObject(botyp.MESSLOKATION),
			bo.NewBusinessObject(botyp.ZAEHLER),
			bo.NewBusinessObject(botyp.RECHNUNG),
			bo.NewBusinessObject(botyp.MESSLOKATION),
			bo.NewBusinessObject(botyp.LASTGANG),

			// Business objects created via bo.NewBusinessObject are pointers, but the concrete types implement
			// bo.BusinessObject via value receivers, so a slice of these objects may contain non-pointer values
			// as well. We include one of them here to show that these work as expected, too.
			bo.Marktteilnehmer{},
		},
	}

	var preisblattValues []*bo.Preisblatt = market_communication.GetAll[bo.Preisblatt](boneyComb)
	fmt.Printf("objects of type Preisblatt     : %d\n", len(preisblattValues))

	var marktteilnehmerValues []*bo.Marktteilnehmer = market_communication.GetAll[bo.Marktteilnehmer](boneyComb)
	fmt.Printf("objects of type Marktteilnehmer: %d\n", len(marktteilnehmerValues))

	var lastgangValues []*bo.Lastgang = market_communication.GetAll[bo.Lastgang](boneyComb)
	fmt.Printf("objects of type Lastgang       : %d\n", len(lastgangValues))

	// Output:
	// objects of type Preisblatt     : 0
	// objects of type Marktteilnehmer: 2
	// objects of type Lastgang       : 1
}
