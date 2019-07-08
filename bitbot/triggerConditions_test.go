package bitbot

import (
	"testing"
)

// These should test basic passing cases
func TestBasicNamedTriggers(t *testing.T) {
	triggerTests := map[string]NamedTrigger{
		"!shrug":                      ShrugTrigger,
		"!skip":                       SkipTrigger,
		"!info":                       InfoTrigger,
		"!roll":                       RollTrigger,
		"bitbot choose foo, bar":      DecisionsTrigger,
		"Nickname is already in use.": NickTakenTrigger,
	}
	b := makeMockBot("bitbot")

	// Batch test all the easy triggers
	for content, trigger := range triggerTests {
		m := makeMockMessage("foo", content)
		ok := trigger.Condition(b, m)
		if !ok {
			t.Errorf("Trigger %s did not activate. Expected true when given m.Content of %s", trigger.ID, m.Content)
		}
	}
}
