package bitbot

import (
	"io"
	"net/http"
	"net/http/httptest"
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

func TestLookupPageTitle(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><head><title>thetitle</title></head></html>")
	}))

	title := lookupPageTitle("take a look at this " + testServer.URL)
	if title != "thetitle" {
		t.Log("Title not extracted from response")
		t.Fail()
	}
}

func TestLookupPageTitleRedirect(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/redirect/" {
			t.Log("Redirecting...")
			http.Redirect(w, r, "/good/", 302)
		} else {
			if !strings.HasSuffix(r.Header.Get("Referer"), "/redirect/") {
				t.Log("Redirect was bypassed")
				t.Fail()
			}
			io.WriteString(w, "<html><head><title>the_redirect_title</title></head></html>")
		}
	}))

	title := lookupPageTitle("take a look at this " + testServer.URL + "/redirect/")
	if title != "the_redirect_title" {
		t.Log("Title not extracted from response")
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
