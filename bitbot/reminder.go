package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	"log"
	"strings"
	"time"
)

var (
	location   *time.Location
	timeFormat string
)

type Event struct {
	ID          int `gorm:"unique;AUTO_INCREMENT;PRIMARY_KEY"`
	Channel     string
	Author      string
	Description string
	People      string
	Time        time.Time
}

var ReminderTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "reminder",
	Help: "Set up events and remind them to concerned people. Usage: !remind list|time|add|remove|delete|join|part",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!remind")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		timeFormat = "2006-01-02 15:04"

		var err error
		location, err = time.LoadLocation("UTC")
		if err != nil {
			irc.Reply(m, "Something went wrong: Couldn't load timezone")
			log.Println("Reminder : Couldn't load UTC timezone", err)
		}

		b.DB.AutoMigrate(&Event{})

		splitMSG := strings.Split(m.Content, " ")
		if len(splitMSG) < 2 {
			irc.Reply(m, "Not enough arguments provided")
			return true
		}

		switch splitMSG[1] {
		case "time":
			irc.Reply(m, getTime())
		case "add":
			irc.Reply(m, addEvent(m, irc))
		case "list":
			irc.Reply(m, listEvents(m, irc))
		default:
			irc.Reply(m, "Wrong argument")
		}
		return true
	},
}

// Lists all the awaiting events in PM
func listEvents(message *hbot.Message, bot *hbot.Bot) string {
	// Get all the db rows, iterate through them, format them and send them to pm
	rows, err := b.DB.Model(&Event{}).Rows()
	if err != nil {
		log.Println("Reminder: Couldn't get DB rows", err)
	}

	var (
		event                   Event
		eventDescriptionMessage string
	)
	for rows.Next() {
		b.DB.ScanRows(rows, &event)
		eventDescriptionMessage = fmt.Sprintf(
			"%d : [ %s ] at %s. Event author : %s, in channel %s",
			event.ID,
			event.Description,
			event.Time.Format(timeFormat),
			event.Author,
			event.Channel)
		bot.Msg(message.Name, eventDescriptionMessage)
	}

	return "I've PM'd you the list of awaiting events"
}

// Parses an event adding message and adds the event
func addEvent(message *hbot.Message, bot *hbot.Bot) string {
	// Parsing the message
	channel := message.To
	author := message.From
	msg := strings.Split(message.Content, " ")
	description := strings.Join(msg[2:len(msg)-2], " ") // FIXME: ugly
	timeOfEvent, err := time.Parse(timeFormat, strings.Join(msg[len(msg)-2:], " ")) // FIXME: ugly
	if err != nil {
		return fmt.Sprintf(
			"Couldn't parse request format is \"!remind add Jitsi Meeting %s\"",
			timeFormat)
	}

	// Adding it to the DB
	event := Event{
		Channel:     channel,
		Author:      author,
		Description: description,
		Time:        timeOfEvent,
		People:      fmt.Sprintf("%s ", author)}
	b.DB.NewRecord(event)
	b.DB.Create(&event)

	// Launch a background routine that will HL interested people and remove the event from the DB.
	eventTimer := time.NewTimer(event.Time.Sub(time.Now()))
	go func (){
		<-eventTimer.C
		var timerEvent Event
		b.DB.Where("Author = ? AND Description = ?",
			event.Author, event.Description).Find(&timerEvent)

		bot.Reply(message,
			fmt.Sprintf("%s : %s",
				timerEvent.Description,
				timerEvent.People))

		deleteEventByID(timerEvent.ID)
	}()

	// Feedback
	return fmt.Sprintf("Adding event \"%s\" by %s, at %s in %s",
		description,
		author,
		timeOfEvent.Format(timeFormat),
		channel)
}

func getTime() string {
	now := time.Now().In(location)
	return now.Format(timeFormat)
}

func deleteEventByID(id int) {
	b.DB.Where("ID = ?", id).Delete(Event{})
}
