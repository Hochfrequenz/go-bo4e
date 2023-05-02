package bo_test

import (
	"fmt"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

// SetupSuite sets up the tests
func (s *Suite) SetupSuite() {
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

// VerfiySuccessfulValidations asserts that the vali validator does not fail for all objects provided
func VerfiySuccessfulValidations(s *Suite, vali *validator.Validate, validObjects []bo.BusinessObject) {
	// ToDo: use generics as soon as golangs allows to
	for _, validObject := range validObjects {
		err := vali.Struct(validObject)
		then.AssertThat(s.T(), validObject.GetBoTyp(), is.Not(is.EqualTo[botyp.BOTyp](0))) // 0 would be iota/uninitialized botyp.BOTyp
		then.AssertThat(s.T(), err, is.Nil())
	}
}

// VerfiyFailedValidations asserts that the vali validator fails with the expected tag for every object
func VerfiyFailedValidations(s *Suite, vali *validator.Validate, tagInvalidObjectsMap map[string][]interface{}) {
	// ToDo: use generics as soon as golangs allows to
	for validationTag, invalidObjects := range tagInvalidObjectsMap {
		for _, invalidObject := range invalidObjects {
			err := vali.Struct(invalidObject)
			then.AssertThat(s.T(), err, is.Not(is.Nil()))
			tagFound := false
			for _, validationError := range err.(validator.ValidationErrors) {
				then.AssertThat(s.T(), validationError, is.Not(is.EqualTo[validator.FieldError](nil)))
				// sometimes there's more than one tag/validation error
				if validationError.Tag() == validationTag {
					tagFound = true
					break
				}
			}
			if !tagFound {
				then.AssertThat(s.T(), validationTag, is.Nil())
			}
			then.AssertThat(s.T(), tagFound, is.True())
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
