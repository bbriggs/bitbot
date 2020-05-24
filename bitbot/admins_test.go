package bitbot

import (
	"testing"
)

func TestIsAdmin(t *testing.T) {
	mockconf := Config{
		Admins: ACL{
			Permitted: []string{"admin1@admin.net"}}}

	mockbot := &Bot{
		Bot:    makeMockBot("bitbot"),
		Config: mockconf}

	// Admins
	testMessage := makeMockMessage("admin1", "this should pass")
	testMessage.Host = "admin.net"

	if !mockbot.isAdmin(testMessage) {
		t.Errorf("Admins are incorrectly detected: %s@%s wasn't detected as an admin, despite appearing in %v",
			testMessage.From,
			testMessage.Host,
			mockbot.Config.Admins.Permitted)
	}

	//Non admins
	testMessage = makeMockMessage("pleb1", "this shouldn't pass")
	testMessage.Host = "pleb.net"

	if mockbot.isAdmin(testMessage) {
		t.Errorf("Admins are incorrectly detected: %s@%s was detected as an admin, despite not appearing in %v",
			testMessage.From,
			testMessage.Host,
			mockbot.Config.Admins.Permitted)
	}

	// Faking nick
	testMessage = makeMockMessage("admin1", "this shouldn't pass")
	testMessage.Host = "pleb.net"
	if mockbot.isAdmin(testMessage) {
		t.Errorf("Nick Stealing : Admins are incorrectly detected: %s@%s was detected as an admin, despite not appearing in %v",
			testMessage.From,
			testMessage.Host,
			mockbot.Config.Admins.Permitted)
	}

	// Faking host // Is that even possible without owning the admin ?
	testMessage = makeMockMessage("pleb1", "this shouldn't pass")
	testMessage.Host = "admin.net"
	if mockbot.isAdmin(testMessage) {
		t.Errorf("Host faking : Admins are incorrectly detected: %s@%s was detected as an admin, despite not appearing in %v",
			testMessage.From,
			testMessage.Host,
			mockbot.Config.Admins.Permitted)
	}

}
