package typemapper_test

import (
	"testing"

	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

func BenchmarkTypeFromValue(b *testing.B) {
	values := make([]any, 900)
	for i := range 300 {
		values[i*3+0] = (*string)(nil)
		values[i*3+1] = "foo"
		values[i*3+2] = new("bar")
	}

	results := make([]int, 900)

	for b.Loop() {
		for i := range values {
			results[i], _ = typemapper.TypeFromValue(values[i], enum)
		}
	}
}
