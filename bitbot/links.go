package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"mvdan.cc/xurls/v2"
	"net/http"
	"net/url"
)

var URLReaderTrigger = NamedTrigger{
	ID:   "urls",
	Help: "Looks up URLs in chat and returns the page title as a message.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isURL(m.Content)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := lookupPageTitle(m.Content)
		if resp != "" {
			irc.Reply(m, lookupPageTitle(m.Content))
			if len(m.Content) > 70 {
				irc.Reply(m, shortenURL(m.Content))
			}
		}
		return true
	},
}

func shortenURL(uri string) string {
	/* We are using 0x0.st */
	resp, err := http.PostForm("https://0x0.st", url.Values{"shorten": {uri}})
	if err != nil {
		log.Println("Coudln't shorten url : ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Coudln't shorten url : ", err)
	}
	return string(body)
}

func isURL(message string) bool {
	return xurls.Strict().MatchString(message)
}

func lookupPageTitle(message string) string {
	url := xurls.Strict().FindString(message)
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	log.Println("Unable to lookup page")
	if title, ok := GetHtmlTitle(resp.Body); ok {
		return (title)
	} else {
		log.Println("Unable to lookup page")
		return ("")
	}
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		if n.FirstChild != nil {
			if len(n.FirstChild.Data) > 350 {
				return (n.FirstChild.Data[:350] + "..."), true
			}
			return n.FirstChild.Data, true
		} else {
			return "", false
		}
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
	doc, err := html.Parse(&io.LimitedReader{R: r, N: 65535})
	if err != nil {
		return "", false
	}
	title, ok := traverse(doc)
	if !ok {
		return "", false
	}
	if len(title) == 0 {
		return " ", false
	}
	return title, ok
}
