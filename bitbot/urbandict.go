package bitbot

// partly stolen from https://github.com/dpatrie/urbandictionary
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var UrbanDictionaryTrigger = NamedTrigger{
	ID:   "urbandict",
	Help: "Get an urban dictionary issued definition. Usage: !urbd [term]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!ud")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := urbanDefinition(m.Content)
		irc.Reply(m, resp)
		return true
	},
}

func urbanDefinition(message string) string {
	term := strings.SplitAfterN(message, " ", 2)[1] // Strip trigger word
	res, err := urbanDictQuery(term)
	if err != nil {
		log.Println(err)
		return "The search failed"
	}

	if len(res.Results) > 0 {
		return fmt.Sprintf("%s: %s", term, cleanDef(res.Results[0].Definition))
	}
	return "No definition for that word"
}

func cleanDef(def string) string {
	def = strings.ReplaceAll(def, "[", "")
	def = strings.ReplaceAll(def, "]", "")

	return def
}

type searchResult struct {
	Type    string `json:"result_type"`
	Tags    []string
	Results []result `json:"list"`
	Sounds  []string
}

type result struct {
	Author     string
	Word       string
	Definition string
	Example    string
	Permalink  string
	Upvote     int `json:"thumbs_up"`
	Downvote   int `json:"thumbs_down"`
}

func urbanDictQuery(searchTerm string) (*searchResult, error) {
	const baseURL = "http://api.urbandictionary.com/v0/define?term="
	resp, err := http.Get(baseURL + url.QueryEscape(searchTerm))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Response was not a 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &searchResult{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
