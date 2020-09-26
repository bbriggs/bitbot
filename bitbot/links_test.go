package bitbot

import (
	"io"
	"strings"
	"testing"
)

func TestCleanTitle(t *testing.T) {
	answered := ""
	// tests["title"]"awaited answer"
	tests := make(map[string]string)
	tests["aa"] = "aa"
	tests[" aa "] = "aa"
	tests["  aa  "] = "aa"
	tests["a  a"] = "a a"
	tests["	a	a"] = "a a"
	tests["aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"] =
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa..."

	for original, awaited := range tests {
		answered = cleanTitle(original)
		if answered != awaited {
			t.Errorf(
				"Failed at cleaning title cleanTitle('%s'): expected '%s', got '%s'.",
				original,
				awaited,
				answered)
		}
	}
}

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

func TestUrlShortening(t *testing.T) {
	go func() {
		title := shortenURL("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
		if !strings.Contains(title, "0x0.st") {
			t.Log("Didn't properly shorten URL")
			t.Fail()
		}
	}()
}
