package bitbot

import (
	"fmt"
	"github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
	"regexp"
	"strconv"
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
	Help: "Set up events and remind them to concerned people. Usage: !remind list|time|add|remove|join|part",
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!remind")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		timeFormat = "2006-01-02 15:04"

		var err error
		location, err = time.LoadLocation("UTC")
		if err != nil {
			irc.Reply(m, "Something went wrong: Couldn't load timezone")
			log.Error("Reminder : Couldn't load UTC timezone", err.Error())
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
		case "remove":
			irc.Reply(m, removeEvent(m, irc))
		case "list":
			irc.Reply(m, listEvents(m, irc))
		case "join":
			irc.Reply(m, joinEvent(m))
		case "part":
			irc.Reply(m, partEvent(m))
		default:
			irc.Reply(m, "Wrong argument")
		}
		return true
	},
}

type wrongFormatError struct {
	arg string
}

func (e *wrongFormatError) Error() string {
	return fmt.Sprintf("%s : is not of the awaited !remind [command] [ID] format", e.arg)
}

func getMessageIDFromString(body string) (int, error) {
	// Parse message
	msg := strings.Split(body, " ")
	isAnID, err := regexp.MatchString("[0-9]+", msg[2])
	if err != nil {
		log.Info("Not and ID :",err)
	}

	if len(msg) != 3 || !isAnID {
		return -1, &wrongFormatError{body}
	}
	id, _ := strconv.Atoi(msg[2])
	return id, nil
}

// Signal yourself as interested in an event (Facebook™)
func joinEvent(message *hbot.Message) string {
	id, err := getMessageIDFromString(message.Content)
	if err != nil {
		return "Wrong command. format is : !remind join [ID]"
	}

	var event Event
	b.DB.Where("ID = ?", id).Take(&event)

	if strings.Contains(event.People, message.Name) {
		b.DB.Save(&event)
		return "You already subscribed to this event"
	}
	event.People = fmt.Sprintf("%s%s ", event.People, message.Name)
	b.DB.Save(&event)

	feedback := fmt.Sprintf("Added %s to \"%s\"",
		message.Name,
		event.Description)

	return feedback
}

func partEvent(message *hbot.Message) string {
	id, err := getMessageIDFromString(message.Content)
	if err != nil {
		return "Wrong command. format is : !remind part [ID]"
	}

	var event Event
	b.DB.Where("ID = ?", id).Take(&event)

	event.People = strings.Replace(event.People, message.Name+" ", "", -1)
	b.DB.Save(&event)

	feedback := fmt.Sprintf("Removed %s from \"%s\"",
		message.Name,
		event.Description)

	return feedback
}

// Remove an event given by his ID
func removeEvent(message *hbot.Message, bot *hbot.Bot) string {
	id, err := getMessageIDFromString(message.Content)
	if err != nil {
		return "Wrong command. format is : !remind remove [ID]"
	}

	var event Event
	b.DB.Where("ID = ? AND Author = ?", id, message.Name).Take(&event)

	// Feedback Message construction
	var feedbackMessage string
	if event.ID == id {
		feedbackMessage = fmt.Sprintf("Deleted event %d : %s",
			event.ID,
			event.Description)

		// Delete
		b.DB.Delete(&event)
	} else {
		feedbackMessage = "No event you own with that ID"
	}

	return feedbackMessage
}

// Lists all the awaiting events in PM
func listEvents(message *hbot.Message, bot *hbot.Bot) string {
	// Get all the db rows, iterate through them, format them and send them to pm
	rows, err := b.DB.Model(&Event{}).Rows()
	if err != nil {
		log.Error("Reminder: Couldn't get DB rows", err)
	}

	var (
		event                   Event
		eventDescriptionMessage string
	)
	for rows.Next() {
		b.DB.ScanRows(rows, &event)
		eventDescriptionMessage = fmt.Sprintf(
			"%d : [ %s ] at %s. Event author : %s, in channel %s, with %s",
			event.ID,
			event.Description,
			event.Time.Format(timeFormat),
			event.Author,
			event.Channel,
			event.People)
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
	// Everything in the message content that isn't a command or a time 
	// specification is considered description
	description := strings.Join(msg[2:len(msg)-2], " ")
	
	// We take the two last parts of the message (with space as the separator)
	// and parse them as a time.
	timeOfEvent, err := time.Parse(timeFormat, strings.Join(msg[len(msg)-2:], " "))
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

	// Launch a background routine that will HL interested people and clean the DB.
	eventTimer := time.NewTimer(event.Time.Sub(time.Now()))
	go func() {
		<-eventTimer.C
		var timerEvent Event
		b.DB.Where("Author = ? AND Description = ?",
			event.Author, event.Description).Find(&timerEvent)

		bot.Reply(message,
			fmt.Sprintf("%s : %s",
				timerEvent.Description,
				timerEvent.People))

		b.DB.Where("ID = ?", timerEvent.ID).Delete(Event{})
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
