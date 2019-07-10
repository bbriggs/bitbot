package bitbot

import (
	"io"
	"strings"
	"testing"
)

func TestGetHTMLTitle(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		io.WriteString(w, "<html><head><title>")
		for i := 0; i < 1000000000; i++ {
			io.WriteString(w, "aaa")
		}
		io.WriteString(w, "</title></head></html>")
		w.Close()
	}()
	title, found := GetHtmlTitle(r)
	if found != true {
		t.Log("could not parse huge title")
		t.Fail()
	}
	if len(title) > 353 || !strings.HasPrefix(title, "aaaa") {
		t.Errorf("The title is too long. Expected 353 chars at most, got %d", len(title))
		t.Fail()
	}
}

func TestGetHTMLTitleWithSmallTitle(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		io.WriteString(w, "<html><head><title>")
		io.WriteString(w, "aaa")
		io.WriteString(w, "</title></head></html>")
		w.Close()
	}()
	title, found := GetHtmlTitle(r)
	if found != true {
		t.Log("could not parse huge title")
		t.Fail()
	}
	if !strings.HasPrefix(title, "aaa") {
		t.Log("unexpected title")
		t.Fail()
	}
}

func TestGetHTMLTitleWithEmptyTitle(t *testing.T) {

	r, w := io.Pipe()
	go func() {
		io.WriteString(w, "<html><head><title>")
		io.WriteString(w, "</title></head></html>")
		w.Close()
	}()
	_, found := GetHtmlTitle(r)
	if found {
		t.Log("Returned true on an empty title")
		t.Fail()
	}
}

func TestAntiUrlFlood(t *testing.T) {
	t.Log("Note : If the previous test failed, this one will hang indefinitely")

	message := "https://secops.space/"

	lookupPageTitle(message) // We make sure that the title musn't be looked for

	title := lookupPageTitle(message)
	if title != "" {
		t.Log(lastTimeTitleLookup)
		t.Log(title)
		t.Log("Doesn't avoid title flooding")
		t.Fail()
	}
}
