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
	if len(title) > 120 || !strings.HasPrefix(title, "aaaa") {
		t.Log("unexpected title")
		t.Fail()
	}
}
