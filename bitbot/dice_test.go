package bitbot

import (
	"strconv"
	"testing"
)

func TestRoll(t *testing.T) {
	var r string

	checks := make(map[string]string)
	// Tests here
	checks["1d4"] = "4"

	if r = roll("failure"); r != DICE_USAGE {
		t.Errorf("roll('failure') outputed '%s', instead of wanted '%s'",
			r, DICE_USAGE)
	}

	if ri, _ := strconv.Atoi(roll("1d4")); ri > 20 {
		t.Errorf("roll('failure') outputed '%s', instead of wanted '%s'",
			r, DICE_USAGE)
	}
}
