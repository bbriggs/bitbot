package bitbot

import (
	"math/rand"
	"testing"
)

func TestRoll(t *testing.T) {
	rand.Seed(1)
	var r string

	checks := make(map[string]string)
	// Tests here
	checks["failure"] = "Usage: [num dice]d[sides](+/-num) (opt: if fudging)"
	checks["1d4"] = "4"
	checks["2d20"] = "10"

	for inp, out := range checks {
		r = roll(inp)
		if r != out {
			t.Errorf("roll('%s') outputed '%s', instead of wanted '%s'",
				inp, r, out)
		}
	}
}
