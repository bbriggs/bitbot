package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"testing"
)

func TestBeefDetection(t *testing.T) {
	m := makeMockMessage("bot", "beefy")
	b := &hbot.Bot{}
	ok := BeefyTrigger.Condition(b, m)
	if !ok {
		t.Errorf("Trigger did not activate. Expected true when given m.Trailing of %s", m.Trailing())
	}
}

func TestBigBeefyLetters(t *testing.T) {
	m := makeMockMessage("bot", "BEEFY")
	b := &hbot.Bot{}
	ok := BeefyTrigger.Condition(b, m)
	if !ok {
		t.Errorf("Trigger did not activate. Expected true when given m.Trailing of %s", m.Trailing())
	}
}

func TestForNoBeef(t *testing.T) {
	m := makeMockMessage("bot", "pls work")
	b := &hbot.Bot{}
	ok := BeefyTrigger.Condition(b, m)
	if ok {
		t.Errorf("Trigger activated. Expected false when given m.Trailing of %s", m.Trailing())
	}
}
