package bo

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
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
func VerfiySuccessfulValidations(s *Suite, vali *validator.Validate, validObjects []interface{}) {
	// ToDo: use generics as soon as golangs allows to
	for _, validObject := range validObjects {
		err := vali.Struct(validObject)
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
				then.AssertThat(s.T(), validationError, is.Not(is.Nil()))
				// sometimes theres more than one tag/validation error
				if validationError.Tag() == validationTag {
					tagFound = true
					break
				}
			}
			then.AssertThat(s.T(), tagFound, is.True())
		}
	}
}
