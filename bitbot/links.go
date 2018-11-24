package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"golang.org/x/net/html"
	"io"
	"mvdan.cc/xurls"
	"net/http"
)

var URLReaderTrigger = NamedTrigger{
	ID: "urls",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isURL(m.Content)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
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
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	fmt.Println("Unable to lookup page")
	if title, ok := GetHtmlTitle(resp.Body); ok {
		return (title)
	} else {
		fmt.Println("Unable to lookup page")
		return ("")
	}
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		if len(n.FirstChild.Data) > 120 {
			return n.FirstChild.Data[:120], true
		}
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
	doc, err := html.Parse(&io.LimitedReader{R: r, N: 4096})
	if err != nil {
		return "", false
	}

	return traverse(doc)
}
