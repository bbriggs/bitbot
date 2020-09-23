package bitbot

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/whyrusleeping/hellabot"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/net/html"
	"mvdan.cc/xurls/v2"
)

var URLReaderTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "urls",
	Help: "Looks up URLs in chat and returns the page title as a message.",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && isURL(m.Content)
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		title := lookupPageTitle(m.Content)
		if title != "" {
			if len(m.Content) > 70 {
				short := shortenURL(m.Content)
				short = strings.TrimRight(short, "\n") //triming
				title = fmt.Sprintf("%s %s", short, title)
			}
			title = cleanTitle(title)
			irc.Reply(m, title)
		}
		return true
	},
	Init: func() error {
		var err error
		return b.EmbDB.Update(func(tx *bolt.Tx) error {
			_, err = tx.CreateBucketIfNotExists([]byte("urlsCache"))
			if err != nil {
				b.Config.Logger.Error("couldn't create url caching bucket: %s", err)
			} else {
				b.Config.Logger.Info("Created/Read urls caching bucket")
			}
			return nil
		})
	},
}

func cleanTitle(title string) string {
	maxLength := 70

	re := regexp.MustCompile(`[ \t\r\n]+`)

	title = strings.Trim(title, " \t\r\n")

	title = re.ReplaceAllString(title, " ")

	if len(title) > maxLength {
		title = fmt.Sprintf("%s...", title[0:67])
	}
	return title
}

func shortenURL(uri string) string {
	// extract url
	uri = xurls.Strict().FindString(uri)

	/* We are using 0x0.st */
	resp, err := http.PostForm("https://0x0.st", url.Values{"shorten": {uri}})
	if err != nil {
		b.Config.Logger.Warn("Coudln't shorten url", "error", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		b.Config.Logger.Warn("Coudln't shorten url", "error", err)
	}

	short := string(body)
	return short
}

func isURL(message string) bool {
	return xurls.Strict().MatchString(message)
}

func lookupPageTitle(message string) string {
	var cached []byte

	url := xurls.Strict().FindString(message)

	err := b.EmbDB.View(func(tx *bolt.Tx) error {
		urlBucket := tx.Bucket([]byte("urlsCache"))
		cached = urlBucket.Get([]byte(url))

		return nil
	})
	if err != nil {
		b.Config.Logger.Warn("Couldn't access Embedded DB")
	}

	if cached != nil { // We already saw that url
		t := strings.SplitAfterN(string(cached), "|", 3)

		cachedTime, cachedTitle := strings.Trim(t[0], "|"), t[1]

		if lessThanAWeek(cachedTime) {
			b.Config.Logger.Info("Got a cached title")
			return fmt.Sprintf("REEEEEEEEpost (%s): %s", cachedTime, cachedTitle)
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close() //nolint:errcheck,gosec
	if title, ok := GetHtmlTitle(resp.Body); ok {
		err := b.EmbDB.Update(func(tx *bolt.Tx) error {
			urlBucket := tx.Bucket([]byte("urlsCache"))
			err2 := urlBucket.Put([]byte(url), []byte(fmt.Sprintf("%s|%s",
				time.Now().Format(time.UnixDate),
				title)))

			if err2 != nil {
				b.Config.Logger.Warn("Couldn't access Embedded DB")
			}

			return nil
		})
		if err != nil {
			b.Config.Logger.Warn("Couldn't access Embedded DB")
		}

		return (title)
	} else {
		b.Config.Logger.Warn("Unable to lookup page", "error", ok)
		return ("")
	}
}

func lessThanAWeek(t string) bool {
	tt, err := time.Parse(time.UnixDate, t)
	if err != nil {
		b.Config.Logger.Warn("Unable to parse time in embedded DB, it might be corrupted")
		return false
	}

	hoursInAWeek := 168.0
	lessThanAWeekAgo := time.Since(tt).Hours() < hoursInAWeek

	return lessThanAWeekAgo
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

func isTwitterURL(url string) bool {
	match, _ := regexp.MatchString("^https://twitter.com/.+/status/[0-9]+", url)
	return match
}
