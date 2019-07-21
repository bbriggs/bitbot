package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	bbolt "go.etcd.io/bbolt"
	"golang.org/x/net/html"
	"io"
	"mvdan.cc/xurls/v2"
	"net/http"
	"time"
	"strings"
	"reflect"
)

var URLReaderTrigger = NamedTrigger{
	ID: "urls",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isURL(m.Content)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := lookupPageTitle(m.Content)
		if resp != "" {
			irc.Reply(m, resp)
		}
		return true
	},
}

func isURL(message string) bool {
	return xurls.Strict().MatchString(message)
}

func lookupPageTitle(message string) string {
	url := xurls.Strict().FindString(message)
	if cached(url) {
		return cacheGetTitle(url)
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return ""
		}
		defer resp.Body.Close()
		if title, ok := GetHtmlTitle(resp.Body); ok {
			cacheURL(url, title)
			return (title)
		} else {
			fmt.Println("Unable to lookup page")
			return ("")
		}
	}
}

func cached(url string) bool {
	/* Return true if the url is cached
	for less than 2 minutes, else false.
	If it has been for more than 2 mins,
	it deletes the row.*/

	//create the bucket if needed
	b.DB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	urlInCache := false

	b.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		value := string(b.Get([]byte(url)))
		if reflect.TypeOf(value) != nil {
			// If the url was already passed, we check for how long
			lastLookup, _ := time.Parse(time.RFC850,
				strings.Split(value, ":")[1])
				if time.Now().Sub(lastLookup).Minutes() < 2 {
				urlInCache = true
			} else {
				b.Delete([]byte(url))
			}
		}
		return nil
	})

	return urlInCache
}

func cacheURL(url string, title string) {
	/* Puts the url, title and time in a pair
	like url => title:time */
	data := title + ":" + time.Now().Format(time.RFC850)

	b.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		err := b.Put([]byte(url), []byte(data))
		return err
	})
}

func cacheGetTitle(url string) string {
	/* Gets the cached title and updates the lookup time */
	return ""
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
