package testcase

import "testing"

// Map contains testcases. It maps names of subtests to testcases.
type Map[TC Testcase] map[string]TC

// Run runs all the tests contained in m as subtests.
func (m Map[TC]) Run(t *testing.T) {
	for name := range m {
		t.Run(name, m[name].Run)
	}
}

// Testcase represents any testcase.
type Testcase interface {
	// Run execute the testcase.
	Run(t *testing.T)
}
