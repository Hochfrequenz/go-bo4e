package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/internal"
	"strings"
	"testing"
	"time"
)

// TestZeitfensterDeserialization deserializes a Zeitfenster json
func Test_Zeitfenster_Deserialization(t *testing.T) {
	var zeitfenster = com.Zeitfenster{

		Startzeit: internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 8, 0, 0, 0, time.UTC))),
		Endzeit:   internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 14, 30, 0, 0, time.UTC))),
	}
	serializedZeitraum, err := json.Marshal(zeitfenster)
	jsonString := string(serializedZeitraum)
	then.AssertThat(t, strings.Contains(jsonString, "14:30:00"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedZeitraum, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Zeitfenster
	err = json.Unmarshal(serializedZeitraum, &deserializedZeitreihenwert)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZeitreihenwert, is.EqualTo(zeitfenster))
}
