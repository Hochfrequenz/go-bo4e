package testcase

import (
	"encoding/json"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

// JSONDeserializationSucceeds is a simple testcase that just tests that a given string can be JSON-deserialized into an entity type.
type JSONDeserializationSucceeds[Entity any] string

// Run tests whether deserializing the raw string represented by this testcase into a pointer to a value of type Entity succeeds.
func (tc JSONDeserializationSucceeds[Entity]) Run(t *testing.T) {
	t.Parallel()
	var entity Entity
	then.AssertThat(t, json.Unmarshal([]byte(tc), &entity), is.Nil())
}
