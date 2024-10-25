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

// TestErreichbarkeitDeserialization deserializes a Erreichbarkeit json
func Test_Erreichbarkeit_Deserialization(t *testing.T) {

	var erreichbarkeit = com.Erreichbarkeit{
		MontagErreichbarkeit: &com.Zeitfenster{

			Startzeit: internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 8, 0, 0, 0, time.UTC))),
			Endzeit:   internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 14, 30, 0, 0, time.UTC))),
		},
		FreitagErreichbarkeit: &com.Zeitfenster{

			Startzeit: internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 8, 0, 0, 0, time.UTC))),
			Endzeit:   internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 14, 30, 0, 0, time.UTC))),
		},
		Mittagspause: &com.Zeitfenster{

			Startzeit: internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 8, 0, 0, 0, time.UTC))),
			Endzeit:   internal.Ptr(internal.TimeOnly(time.Date(0, 1, 1, 14, 30, 0, 0, time.UTC))),
		},
	}
	serializedErreichbarkeit, err := json.Marshal(erreichbarkeit)
	jsonString := string(serializedErreichbarkeit)
	then.AssertThat(t, strings.Contains(jsonString, "14:30:00"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedErreichbarkeit, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Erreichbarkeit
	err = json.Unmarshal(serializedErreichbarkeit, &deserializedZeitreihenwert)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZeitreihenwert, is.EqualTo(erreichbarkeit))
}
