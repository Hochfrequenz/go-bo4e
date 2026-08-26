package typemapper_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

var enum = map[string]int{
	"foo": 1,
	"bar": 2,
	"baz": 3,
}

func TestTypeFromValue(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		value any

		// check checks whether s and err match expectations.
		check checkFunc
	}{
		"raw nil": {
			value: nil,
			check: allOf(valueIs(0), errIsNil),
		},
		"string pointer nil": {
			value: (*string)(nil),
			check: allOf(valueIs(0), errIsNil),
		},
		"valid string": {
			value: "foo",
			check: allOf(valueIs(1), errIsNil),
		},
		"valid string pointer": {
			value: new("bar"),
			check: allOf(valueIs(2), errIsNil),
		},
		"invalid string": {
			value: "unknown",
			check: errIsNotNil,
		},
	}

	for name, testcase := range testcases {
		t.Run(
			name,
			func(t *testing.T) {
				if err := testcase.check(typemapper.TypeFromValue(testcase.value, enum)); err != nil {
					t.Error(err)
				}
			},
		)
	}
}

type checkFunc func(n int, err error) error

func allOf(checks ...checkFunc) checkFunc {
	return func(n int, err error) error {
		errs := make([]error, len(checks))

		for i := range checks {
			errs[i] = checks[i](n, err)
		}

		return errors.Join(errs...)
	}
}

func valueIs(expected int) checkFunc {
	return func(actual int, _ error) error {
		if actual != expected {
			return fmt.Errorf("expected %d, got %d", expected, actual)
		}

		return nil
	}
}

func errIsNil(_ int, err error) error {
	if err != nil {
		// Intentionally _not_ %w.
		return fmt.Errorf("expected no error, got: %v", err)
	}

	return nil
}

func errIsNotNil(_ int, err error) error {
	if err == nil {
		return errors.New("expected an error")
	}

	return nil
}
