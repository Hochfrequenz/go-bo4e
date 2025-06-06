package com_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/internal"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
)

// Test_Messlokationszuordnung_Deserialization deserializes an MesslokationszuordnungMesslokationszuordnung json
func Test_Messlokationszuordnung_Deserialization(t *testing.T) {
	var messlokationszuordnung = com.Messlokationszuordnung{
		MesslokationsId: internal.Ptr("DE0123456789012345678901234567890"),
		Arithmetik:      internal.Ptr(arithmetischeoperation.ADDITION),
		GueltigSeit:     internal.Ptr(time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)),
		GueltigBis:      internal.Ptr(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
	}
	serializedMesslokationszuordnung, err := json.Marshal(messlokationszuordnung)
	then.AssertThat(t, serializedMesslokationszuordnung, is.Not(is.NilArray[byte]()))

	jsonString := string(serializedMesslokationszuordnung)
	then.AssertThat(t, strings.Contains(jsonString, "ADDITION"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())

	var deserializedMesslokationszuordnung com.Messlokationszuordnung
	err = json.Unmarshal(serializedMesslokationszuordnung, &deserializedMesslokationszuordnung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedMesslokationszuordnung, is.EqualTo(messlokationszuordnung))
}

// Test_Failed_Validation asserts that the validation fails for invalid Messlokationszuordnung
func Test_Failed_MesslokationszuordnungValidation(t *testing.T) {
	validate := validator.New()
	invalidMesslokationszuordnung := map[string][]interface{}{
		"len": {
			com.Messlokationszuordnung{
				MesslokationsId: internal.Ptr("not33"),
			},
		},
		// removed validation at least until https://github.com/go-playground/validator/issues/972 is resolved
		/*
			"ltefield": {
				com.Messlokationszuordnung{
					GueltigSeit: internal.Ptr(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
					GueltigBis:  internal.Ptr(time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC))},
			},
			"gtefield": {
				com.Messlokationszuordnung{
					GueltigSeit: internal.Ptr(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
					GueltigBis:  internal.Ptr(time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC))},
			},*/
	}
	VerifyFailedValidations(t, validate, invalidMesslokationszuordnung)
}

// Test_Successful_Validation asserts that the validation does not fail for a valid Messlokationszuordnung
func Test_Successful_Messlokationszuordnung_Validation(t *testing.T) {
	validate := validator.New()
	validMesslokationszuordnung := []interface{}{
		com.Messlokationszuordnung{
			MesslokationsId: internal.Ptr("DE0123456789012345678901234567890"),
			Arithmetik:      internal.Ptr(arithmetischeoperation.ADDITION),
		},
		com.Messlokationszuordnung{
			MesslokationsId: internal.Ptr("DE0123456789012345678901234567890"),
			Arithmetik:      internal.Ptr(arithmetischeoperation.ADDITION),
			GueltigSeit:     internal.Ptr(time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)),
			GueltigBis:      internal.Ptr(time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)),
		},
	}
	VerifySuccessfulValidations(t, validate, validMesslokationszuordnung)
}

func Test_Serialized_Empty_Messlokationszuordnung_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Messlokationszuordnung{})
}
