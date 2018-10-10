package bitbot

import (
	"mvdan.cc/xurls"
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"io"
	"time"
	"os"
	"encoding/binary"
	bolt "go.etcd.io/bbolt"
	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/whyrusleeping/hellabot"
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

func wasLookedUpInTheLastMintues(url string) bool {
	db, err := newDB()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	wasLooked := false

	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(url)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			t := time.Unix(int64(binary.BigEndian.Uint64(v)), 0).UTC()
			duration := time.Since(t)
			if (duration.Minutes() <= 2) {
				wasLooked = true
			}
		}
		return nil
	})

	if !wasLooked {
		cacheUrl(url, db)
	}

	db.Close()
	return wasLooked
}

func cacheUrl(url string, db *bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {

		if err := tx.Bucket([]byte(url)).Put([]byte(url), []byte(time.Now().Format(time.RFC3339))); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("caching: ", err)
	}
}

func lookupPageTitle(message string) string {
	url := xurls.Strict.FindString(message)

	if wasLookedUpInTheLastMintues(url) {
		return ""
	}

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	fmt.Println("Unable to lookup page")
	if title, ok := GetHtmlTitle(resp.Body); ok {
		return(title)
	} else {
		fmt.Println("Unable to lookup page")
		return("")
	}
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		if (len(n.FirstChild.Data) > 120) {
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
		panic("Fail to parse html")
	}

	return traverse(doc)
}
