package bitbot

import (
	"math/rand"
	"testing"
)

func TestAbyssCondition(t *testing.T) {
	rand.Seed(1)
	m := makeMockMessage("foo", "bar")
	b := makeMockBot("bitbot")

	for i := 1; i <= 1000; i++ {
		if AbyssTrigger.Condition(b, m) {
			return
		}
	}
	t.Error("Trigger did not activate after 1000 attempts with a random seed of 1")
}
