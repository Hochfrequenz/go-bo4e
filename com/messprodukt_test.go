package com_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/verwendungszweck"
	"github.com/hochfrequenz/go-bo4e/internal"
)

func Test_Messprodukt_Deserialization(t *testing.T) {
	messprodukt := com.Messprodukt{
		MessproduktID: internal.Ptr("9991000000044"),
		Verwendungszwecke: []com.Verwendungszweck{
			{
				Marktrolle: marktrolle.NB,
				Zweck: []verwendungszweck.Verwendungszweck{
					verwendungszweck.NETZNUTZUNGSABRECHNUNG,
					verwendungszweck.MEHRMINDERMENGENABRECHNUNG,
				},
			},
		},
	}
	serializedMessprodukt, err := json.Marshal(messprodukt)
	jsonString := string(serializedMessprodukt)
	then.AssertThat(t, strings.Contains(jsonString, "messproduktId"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedMessprodukt, is.Not(is.NilArray[byte]()))

	var deserializedMessprodukt com.Messprodukt
	err = json.Unmarshal(serializedMessprodukt, &deserializedMessprodukt)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedMessprodukt, is.EqualTo[com.Messprodukt](messprodukt))
}
