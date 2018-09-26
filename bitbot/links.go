package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"mvdan.cc/xurls"
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"io"
)

var URLReaderTrigger = hbot.Trigger{
	func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isURL(m.Content)
	},
	func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := lookupPageTitle(m.Content)
		if resp != "" {
			irc.Reply(m, lookupPageTitle(m.Content))
		}
		return true
	},
}

func isURL(message string) bool {
	return xurls.Strict.MatchString(message)
}

func lookupPageTitle(message string) string {
	url := xurls.Strict.FindString(message)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return ""
	}
	fmt.Println("Unable to lookup page")
	if title, ok := GetHtmlTitle(resp.Body); ok {
		return(title)
	} else {
		fmt.Println("Unagle to lookup page")
		return("")
	}
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func GetHtmlTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}
