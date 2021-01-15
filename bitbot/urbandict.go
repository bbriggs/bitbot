package bitbot

// partly stolen from https://github.com/dpatrie/urbandictionary
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/whyrusleeping/hellabot"
)

var UrbanDictionaryTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
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
		b.Config.Logger.Warn("Couldn't query urbandictionary", "error", err)
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
	def = strings.ReplaceAll(def, "\r\n", " ")

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
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", baseURL, url.QueryEscape(searchTerm)), nil) //nolint:noctx
	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Response was not a 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		b.Config.Logger.Warn("UD: Couldn't close body", "error", err)
	}

	res := &searchResult{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
