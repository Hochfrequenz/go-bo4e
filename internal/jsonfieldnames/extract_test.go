package jsonfieldnames_test

import (
	"sort"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
)

func TestExtractFromInvalidTypes(t *testing.T) {
	t.Parallel()

	var stringValue = "foobar"
	var intValue = 42
	var boolValue = true

	inputs := map[string]interface{}{
		"nil":               nil,
		"string":            stringValue,
		"int":               intValue,
		"bool":              boolValue,
		"pointer to string": &stringValue,
		"pointer to int":    &intValue,
		"pointer to bool":   &boolValue,
	}

	for name := range inputs {
		input := inputs[name]

		t.Run(
			name,
			func(t *testing.T) {
				fields, err := jsonfieldnames.Extract(input)

				then.AssertThat(t, fields, is.NilArray[string]())
				then.AssertThat(t, err, is.Not(is.Nil()))
			},
		)
	}
}

func TestExtractFromValidTypes(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		// input is the value that JSON field names are extracted from.
		input interface{}

		// expectedFields are the fields that we expected to be extracted, in alphabetical order.
		expectedFields []string
	}{
		"struct with no fields": {
			input:          struct{}{},
			expectedFields: []string{},
		},
		"struct with fields w/o tags": {
			input: struct {
				Foo int
				Bar int
			}{},
			expectedFields: []string{"Bar", "Foo"},
		},
		"struct with fields with simple tags": {
			input: struct {
				Foo int `json:"foo"`
				Bar int `json:"bar"`
			}{},
			expectedFields: []string{"bar", "foo"},
		},
		"struct with fields that should be ignored by json package": {
			input: struct {
				Foo int `json:"-"`
				Bar int `json:"-"`
			}{},
			expectedFields: []string{},
		},
		"struct with fields that have additional options": {
			input: struct {
				Foo  int `json:"boo,omitempty"`
				Bar  int `json:"far,"`
				Baz  int `json:",omitempty"`
				Dash int `json:"-,"`
			}{},
			expectedFields: []string{"-", "Baz", "boo", "far"},
		},
		"pointer to struct": {
			input: &struct {
				XYZ int `json:"xyz"`
			}{},
			expectedFields: []string{"xyz"},
		},
		"struct with embedded structs w/o explit JSON name": {
			input: struct {
				embeddedStructHello
				embeddedStructWorld `json:",omitempty"`
			}{},
			expectedFields: []string{"hello", "world"},
		},
		"struct with pointers to embedded structs w/o explicit JSON name": {
			input: struct {
				*embeddedStructHello
				*embeddedStructWorld `json:",omitempty"`
			}{},
			expectedFields: []string{"hello", "world"},
		},
		"struct with embedded structs with tags": {
			input: struct {
				embeddedStructHello `json:"embed_1"`
				embeddedStructWorld `json:"embed_2,omitempty"`
			}{},
			expectedFields: []string{"embed_1", "embed_2"},
		},
		"struct with embedded non-structs": {
			input: struct {
				EmbeddedString
				EmbeddedInt
				EmbeddedBool
				EmbeddedInterface
			}{},
			expectedFields: []string{"EmbeddedBool", "EmbeddedInt", "EmbeddedInterface", "EmbeddedString"},
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				fields, err := jsonfieldnames.Extract(testcase.input)
				sort.Strings(fields)

				then.AssertThat(t, fields, is.EqualTo(testcase.expectedFields))
				then.AssertThat(t, err, is.Nil())
			},
		)
	}
}

// embeddedStructHello is used as embedded struct in TestExtractFromValidTypes.
type embeddedStructHello struct {
	Hello string `json:"hello"`
}

// embeddedStructWorld is used as embedded struct in TestExtractFromValidTypes.
type embeddedStructWorld struct {
	World string `json:"world"`
}

// EmbeddedString is used as embedded type in TestExtractFromValidTypes.
type EmbeddedString string

// EmbeddedInt is used as embedded type in TestExtractFromValidTypes.
type EmbeddedInt int

// EmbeddedBool is used as embedded type in TestExtractFromValidTypes.
type EmbeddedBool bool

// EmbeddedInterface is used as embedded type in TestExtractFromValidTypes.
type EmbeddedInterface interface{}
