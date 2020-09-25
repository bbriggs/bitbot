package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v1"
	"testing"
	"time"
)

func TestBeefDetection(t *testing.T) {
	m := &hbot.Message{
		&irc.Message{
			irc.
				ParsePrefix("testServ | testNick [ '!' usr ] [ '@' zbeul ]"),
			"PRIVMSG",
			[]string{},
			":beefy",
			true,
		},
		":beefy",
		time.Now(),
		"",
		"",
	}
	b := &hbot.Bot{}
	ok := BeefyTrigger.Condition(b, m)
	if !ok {
		t.Errorf("Trigger did not activate. Expected true when given m.Trailing of %s", m.Trailing)
	}
}

func TestBigBeefyLetters(t *testing.T) {
	m := &hbot.Message{
		&irc.Message{
			&irc.Prefix{
				"",
				"",
				"",
			},
			"PRIVMSG",
			[]string{},
			":BEEFY",
			true,
		},
		"",
		time.Now(),
		"",
		"",
	}
	b := &hbot.Bot{}
	ok := BeefyTrigger.Condition(b, m)
	if !ok {
		t.Errorf("Trigger did not activate. Expected true when given m.Trailing of %s", m.Trailing)
	}
}

func TestForNoBeef(t *testing.T) {
	m := &hbot.Message{
		&irc.Message{
			&irc.Prefix{
				"",
				"",
				"",
			},
			"PRIVMSG",
			[]string{},
			":soy patties",
			true,
		},
		"",
		time.Now(),
		"",
		"",
	}
	b := &hbot.Bot{}
	ok := BeefyTrigger.Condition(b, m)
	if ok {
		t.Errorf("Trigger activated. Expected false when given m.Trailing of %s", m.Trailing)
	}
}
