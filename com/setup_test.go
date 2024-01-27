package com_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

// VerifySuccessfulValidations asserts that the vali validator does not fail for all objects provided
func VerifySuccessfulValidations(t *testing.T, vali *validator.Validate, validObjects []interface{}) {
	// ToDo: use generics as soon as golangs allows to
	for _, validObject := range validObjects {
		err := vali.Struct(validObject)
		then.AssertThat(t, err, is.Nil())
	}
}

// VerifyFailedValidations asserts that the vali validator fails with the expected tag for every object
func VerifyFailedValidations(t *testing.T, vali *validator.Validate, tagInvalidObjectsMap map[string][]interface{}) {
	// ToDo: use generics as soon as golangs allows to
	for validationTag, invalidObjects := range tagInvalidObjectsMap {
		for _, invalidObject := range invalidObjects {
			err := vali.Struct(invalidObject)
			then.AssertThat(t, err, is.Not(is.Nil()))
			tagFound := false
			for _, validationError := range err.(validator.ValidationErrors) {
				then.AssertThat(t, validationError, is.Not(is.EqualTo[validator.FieldError](nil)))
				// sometimes there's more than one tag/validation error
				if validationError.Tag() == validationTag {
					tagFound = true
					break
				}
			}
			then.AssertThat(t, tagFound, is.True())
		}
	}
}

// newDecimalFromString returns a new decimal for a given string. This allows to inline create new decimals without caring about the error. Anyways this function will panic if there is an error.
// Background ist, that for decimal.NewFromString it is guaranteed, that the created decimal values are deep equatable (https://github.com/shopspring/decimal/blob/9ffd602a49ccb618af362e7a17f6a1e4bcb0afa8/decimal_test.go#L343 )
// This is _not_ the case for decimal.New() or decimal.NewFromFloat()
func newDecimalFromString(s string) decimal.Decimal {
	result, err := decimal.NewFromString(s)
	if err != nil {
		panic(fmt.Errorf("Error while converting '%s'", s))
	}
	return result
}

func assertDoesNotSerializeDefaultEnums(t *testing.T, com interface{}) {
	jsonBytes, err := json.Marshal(com)
	then.AssertThat(t, err, is.Nil())
	jsonString := string(jsonBytes)
	then.AssertThat(t, strings.Contains(jsonString, "(0)"), is.False())
}
