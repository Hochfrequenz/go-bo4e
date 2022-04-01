package sparte

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestStringifySparteForDB checks if the Value method converts the Sparte to its String representaion and if the Scan reads it
func TestStringifySparteForDB(t *testing.T) {
	s := STROM

	v, err := s.Value()

	sb := v.(string)

	var rs Sparte
	rs.Scan(&sb)

	assert.NoError(t, err)
	assert.Equal(t, "STROM", sb)
	assert.Equal(t, STROM, rs)
}
