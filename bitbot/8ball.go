package bitbot

import (
	"github.com/whyrusleeping/hellabot"
	"math/rand"
	"strings"
)

var Magic8BallTrigger = NamedTrigger{ //nolint:gochecknoglobals
	ID:   "8ball",
	Help: "Beseech the magic 8ball. Usage: !8ball [question]",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!8ball")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		resp := magic8responses[rand.Intn(len(magic8responses)-1)]
		irc.Reply(m, resp)
		return true
	},
}

var magic8responses = []string{
	"Signs point to yes.",
	"Yes.",
	"Reply hazy, try again.",
	"Without a doubt.",
	"My sources say no.",
	"As I see it, yes.",
	"You may rely on it.",
	"Concentrate and ask again.",
	"Outlook not so good.",
	"It is decidedly so.",
	"Better not tell you now.",
	"Very doubtful.",
	"Yes - definitely.",
	"It is certain.",
	"Cannot predict now.",
	"Most likely.",
	"Ask again later.",
	"My reply is no.",
	"Outlook good.",
	"Don't count on it.",
	"Yes, in due time.",
	"My sources say no.",
	"Definitely not.",
	"Yes.",
	"You will have to wait.",
	"I have my doubts.",
	"Outlook so so.",
	"Looks good to me!",
	"Who knows?",
	"Looking good!",
	"Probably.",
	"Are you kidding?",
	"Go for it!",
	"Don't bet on it.",
	"Forget about it.",
	"Wtf were you thinking?!",
	"What do i look like, your own personal fortune teller?!",
	"Why o why didn't I take the blue pill.",
	"I do not understand your question, could you please repeat it, this time with a little more intelligence?",
	"I saw that episode... it didn't end well.",
	"You didn't read the fine print before signing that contract did you.",
	"Not happening. Not now, Not ever.",
	"You do realize what you just asked right?",
	"Wait, you must think I have the answer to your question.",
	"Yeah, not in this lifetime.",
	"Fantastic.",
	"You betcha!",
	"But of Course.",
	"I am not sure I know how to parse that.",
	"42",
	"Nein!",
	"Ja.",
	"I know nothing!",
	"I don't always answer your question, but when I do I-oh look a kitty."}
