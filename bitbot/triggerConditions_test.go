package bitbot

import (
	"fmt"
	"testing"
)

// These should test basic passing cases
func TestBasicNamedTriggers(t *testing.T) {
	triggerTests := map[string]NamedTrigger{ //nolint:gochecknoglobals,golint
		"!shrug":                      ShrugTrigger,
		"!skip":                       SkipTrigger,
		"!info":                       InfoTrigger,
		"!roll":                       RollTrigger,
		"bitbot choose foo, bar":      DecisionsTrigger,
		"Nickname is already in use.": NickTakenTrigger,
		"!tableflip":                  TableFlipTrigger,
		"!unflip":                     TableUnflipTrigger,
		"!help":                       HelpTrigger,
		"!8ball":                      Magic8BallTrigger,
		"!babble":                     MarkovResponseTrigger,
		"!markov reset":               MarkovInitTrigger,
		"!markov init":                MarkovInitTrigger,
		"!epeen":                      EpeenTrigger,
		"!ipinfo":                     IPinfoTrigger,
		"owo":                         WeebTrigger,
		"uwu":                         WeebTrigger,
		"oWo":                         WeebTrigger,
		"UWU":                         WeebTrigger,
		"gUwUtinne":                   WeebTrigger,
		"guWutinne":                   WeebTrigger,
		"wow, it's gUwUtinne time":    WeebTrigger,
	}
	b := makeMockBot("bitbot")

	// Batch test all the easy triggers
	for content, trigger := range triggerTests {
		testname := fmt.Sprintf("Trigger %s activation test: %s", trigger.ID, content)
		t.Run(testname, func(t *testing.T) {
			m := makeMockMessage("foo", content)
			ok := trigger.Condition(b, m)
			if !ok {
				t.Errorf("Trigger %s did not activate. Expected true when given m.Content of %s", trigger.ID, m.Content)
			}
		})
	}
}

// False positive testing
func TestBasicNamedTriggersFalsePositives(t *testing.T) {
	triggerTests := map[string]NamedTrigger{
		"away":     WeebTrigger,
		"coworker": WeebTrigger,
		"wow, my guwutinne really needs some polishing": WeebTrigger,
	}
	b := makeMockBot("bitbot")

	// Batch test all the easy triggers
	for content, trigger := range triggerTests {
		testname := fmt.Sprintf("Trigger %s false positive test: %s", trigger.ID, content)
		t.Run(testname, func(t *testing.T) {
			m := makeMockMessage("foo", content)
			ok := trigger.Condition(b, m)
			if ok {
				t.Errorf("Trigger %s activate when it shouldn't have. (m.Content of %s)", trigger.ID, m.Content)
			}
		})
	}
}
