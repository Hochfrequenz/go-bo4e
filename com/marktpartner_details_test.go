package com_test

import (
	"encoding/json"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/internal"
)

func TestMarshalMarktpartnerDetails(t *testing.T) {
	marktpartnerDetails := com.MarktpartnerDetails{
		Rollencodenummer:   internal.Ptr("RC1"),
		Code:               internal.Ptr("C1"),
		Marktrolle:         internal.Ptr(marktrolle.MSB),
		Weiterverpflichtet: internal.Ptr(true),
	}
	marshalledMarktpartnerDetails, err := json.Marshal(marktpartnerDetails)
	then.AssertThat(t, err, is.Nil())
	jsonString := string(marshalledMarktpartnerDetails)
	then.AssertThat(t, jsonString, is.StringContaining(`"rollencodenummer":"RC1"`))
	then.AssertThat(t, jsonString, is.StringContaining(`"code":"C1"`))
	then.AssertThat(t, jsonString, is.StringContaining(`"marktrolle":"MSB"`))
	then.AssertThat(t, jsonString, is.StringContaining(`"weiterverpflichtet":true`))
}

func TestUnmarshalMarktpartnerDetails(t *testing.T) {
	var marktpartnerDetails com.MarktpartnerDetails
	// empty MarktpartnerDetails
	err := json.Unmarshal([]byte("{}"), &marktpartnerDetails)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, marktpartnerDetails.Rollencodenummer, is.NilPtr[string]())
	then.AssertThat(t, marktpartnerDetails.Code, is.NilPtr[string]())
	then.AssertThat(t, marktpartnerDetails.Marktrolle, is.NilPtr[marktrolle.Marktrolle]())
	then.AssertThat(t, marktpartnerDetails.Weiterverpflichtet, is.NilPtr[bool]())
	// complete MarktpartnerDetails
	err = json.Unmarshal([]byte(`
	{
		"rollencodenummer": "RC2",
		"code": "C2",
		"marktrolle": "MSB",
		"weiterverpflichtet": true
	}
	`), &marktpartnerDetails)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, marktpartnerDetails.Rollencodenummer, is.EqualTo(internal.Ptr("RC2")))
	then.AssertThat(t, marktpartnerDetails.Code, is.EqualTo(internal.Ptr("C2")))
	then.AssertThat(t, marktpartnerDetails.Marktrolle, is.EqualTo(internal.Ptr(marktrolle.MSB)))
	then.AssertThat(t, marktpartnerDetails.Weiterverpflichtet, is.EqualTo(internal.Ptr(true)))
}
