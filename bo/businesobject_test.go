package bo_test

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"

	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

func TestNewBusinessObject(t *testing.T) {
	// botyp.Anfrage ist der größte ENUM, neue Enums dahinter würden nicht getestet werden.
	// Dynamisch alle zu ermitteln ist aber nicht so easy.
	for i := 1; i <= int(botyp.ANFRAGE); i++ {
		t.Run(botyp.BOTyp(i).String(), func(t *testing.T) {
			object := bo.NewBusinessObject(botyp.BOTyp(i))
			switch botyp.BOTyp(i) { // manche Typen werden nicht gemapped, weil die Objekte nicht implementiert sind
			// TODO: diese Typen könnten noch in BO4E dotnet und hier implementiert oder im Standard deprecated werden
			case botyp.AUSSCHREIBUNG,
				botyp.KOSTEN,
				botyp.PREISBLATTDIENSTLEISTUNG,
				botyp.TARIFINFO,
				botyp.TARIFPREISBLATT,
				botyp.ZEITREIHE:
				t.SkipNow()
			}
			then.AssertThat(t, object, is.Not(is.EqualTo(bo.BusinessObject(nil))))
			if object != nil {
				then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
				then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.BOTyp(i)))
			}
		})
	}
}
