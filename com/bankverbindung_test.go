package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"strings"
	"testing"
)

func Test_Bankverbindung_Deserialization(t *testing.T) {
	var bankverbindung = com.Bankverbindung{
		IBAN:         "0000000000000000000000",
		Bankkennung:  "13089204",
		Kontoinhaber: "Max Mustermann",
		Bankname:     "Eine Bank",
	}

	serializedBankverbindung, err := json.Marshal(bankverbindung)
	jsonString := string(serializedBankverbindung)
	then.AssertThat(t, strings.Contains(jsonString, "0000000000000000000000"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedBankverbindung, is.Not(is.NilArray[byte]()))
	var deserializedBankverbindung com.Bankverbindung
	err = json.Unmarshal(serializedBankverbindung, &deserializedBankverbindung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedBankverbindung, is.EqualTo(bankverbindung))
}
