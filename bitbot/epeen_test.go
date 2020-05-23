package bitbot

import (
	"regexp"
	"testing"
)

func TestMakeEpeenAnswer(t *testing.T) {
	matched, err := regexp.Match("8=+D", []byte(makeEpeenAnswer("testname")))
	if err != nil {
		t.Error("Couldn't test regex against makeEpeenAnswer")
	}

	if !matched {
		t.Error("makeEpeenAnswer outputed a malformed epeen")
	}
}
