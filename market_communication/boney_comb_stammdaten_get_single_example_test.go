package market_communication_test

import (
	"fmt"

	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

func ExampleGetSingle() {
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

	preisblatt, err := market_communication.GetSingle[bo.Preisblatt](boneyComb)
	if err != nil {
		fmt.Println("Preisblatt count is not 1")
	}
	if preisblatt != nil {
		fmt.Println("found a Preisblatt")
	}

	marktteilnehmer, err := market_communication.GetSingle[bo.Marktteilnehmer](boneyComb)
	if err != nil {
		fmt.Println("Marktteilnehmer count is not 1")
	}
	if marktteilnehmer != nil {
		fmt.Println("found a Marktteilnehmer")
	}

	lastgang, err := market_communication.GetSingle[bo.Lastgang](boneyComb)
	if err != nil {
		fmt.Println("Lastgang count is not 1")
	}
	if lastgang != nil {
		fmt.Println("found a Lastgang")
	}

	// Output:
	// Preisblatt count is not 1
	// Marktteilnehmer count is not 1
	// found a Lastgang
}
