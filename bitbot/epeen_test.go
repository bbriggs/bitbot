package bitbot

import (
	"fmt"
	"testing"
)

func TestEpeenAction(t *testing.T) {
	fmt.Println("aa")
	m := makeMockMessage("foo", "!epeen")
	b := makeMockBot("bitbot")

	b.Reply(m, "ha")
	fmt.Println("aaaa")

	if EpeenTrigger.Action(b, m) {
		return
	}

	t.Error("EpeenTrigger condition isn't consuming messages")
}
