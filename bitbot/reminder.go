package bitbot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/whyrusleeping/hellabot"
)

var (
	location   *time.Location
	timeFormat string
)

// ReminderEvent : The Gorm struct that represents an event in the DB.
type ReminderEvent struct {
	ID          int `gorm:"unique;AUTO_INCREMENT;PRIMARY_KEY"`
	Channel     string
	Author      string
	Description string
	People      string
	Time        time.Time
}

// This error is used when a badly formatted call of the trigger is made.
type wrongFormatError struct {
	arg string
}

func (e *wrongFormatError) Error() string {
	return fmt.Sprintf("%s : is not of the awaited format", e.arg)
}

func noAccessDBMessage() string {
	return "Something went wrong, no access to database"
}

var ReminderTrigger = NamedTrigger{ //nolint:gochecknoglobals,golint
	ID:   "reminder",
	Help: "Set up events and remind them to concerned people. Usage: !remind list|time|add|remove|join|part",
	Init: func() error {
		var err error

		location, err = time.LoadLocation("UTC")
		if err != nil {
			b.Config.Logger.Error("Reminder : Couldn't load UTC timezone", err.Error())
			return err
		}

		return b.DB.AutoMigrate(&ReminderEvent{}).Error
	},
	Condition: func(irc *hbot.Bot, m *hbot.Message) bool {
		return m.Command == "PRIVMSG" && strings.HasPrefix(m.Trailing, "!remind")
	},
	Action: func(irc *hbot.Bot, m *hbot.Message) bool {
		timeFormat = "2006-01-02 15:04"

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
			irc.Reply(m, removeEvent(m))
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

// Get the time in UTC, formatted in the expected way.
func getTime() string {
	return fmt.Sprintf("%s UTC", time.Now().In(location).Format(timeFormat))
}

// Parses an event adding message and adds the event
func addEvent(message *hbot.Message, bot *hbot.Bot) string {
	// Parsing the message
	channel := message.To
	author := message.From

	description, datetime, err := parseAddCommandMessage(message.Content)
	if err != nil {
		return fmt.Sprintf(
			"Wrong syntax, use !remind add Jitsi Meeting %s UTC",
			timeFormat)
	}
	// We take the two last parts of the message (with space as the separator)
	// and parse them as a time.
	timeOfEvent, err := time.Parse(timeFormat, datetime)
	if err != nil {
		return fmt.Sprintf(
			"Couldn't parse request format is \"!remind add Jitsi Meeting %s\"",
			timeFormat)
	}

	// Adding it to the DB
	event := ReminderEvent{
		Channel:     channel,
		Author:      author,
		Description: description,
		Time:        timeOfEvent,
		People:      fmt.Sprintf("%s ", author)}
	b.DB.NewRecord(event)

	if err := b.DB.Create(&event); err != nil {
		return "Something went wrong"
	}

	// Launch a background routine that will HL interested people and clean the DB.
	// The magic number 2 is indeed completely arbitrary, but we need it anyway.
	timeToEvent := time.Until(event.Time) - (2 * time.Second)
	eventTimer := time.NewTimer(2 * time.Second) //nolint:gomnd // Arbitrary

	go func() {
		time.Sleep(timeToEvent)
		<-eventTimer.C

		var timerEvent ReminderEvent

		b.DB.Where("Author = ? AND Description = ?",
			event.Author, event.Description).Find(&timerEvent)

		bot.Reply(message,
			fmt.Sprintf("%s : %s",
				timerEvent.Description,
				timerEvent.People))

		b.DB.Where("ID = ?", timerEvent.ID).Delete(ReminderEvent{})
	}()

	// Feedback
	return fmt.Sprintf("Adding event \"%s\" by %s, at %s UTC in %s",
		description,
		author,
		timeOfEvent.Format(timeFormat),
		channel)
}

// Remove an event given by his ID
func removeEvent(message *hbot.Message) string {
	var event ReminderEvent

	id, err := getMessageIDFromString(message.Content)
	if err != nil {
		return "Wrong command. format is : !remind remove [ID]"
	}

	if err := b.DB.Where("ID = ? AND Author = ?", id, message.Name).Take(&event); err != nil {
		return noAccessDBMessage()
	}

	// Feedback Message construction
	var feedbackMessage string
	if event.ID == id {
		feedbackMessage = fmt.Sprintf("Deleted event %d : %s",
			event.ID,
			event.Description)

		// Delete
		if err := b.DB.Delete(&event); err != nil {
			return noAccessDBMessage()
		}
	} else {
		feedbackMessage = "No event you own with that ID"
	}

	return feedbackMessage
}

// Lists all the awaiting events in PM
func listEvents(message *hbot.Message, bot *hbot.Bot) string {
	// Get all the db rows, iterate through them, format them and send them to pm
	rows, err := b.DB.Model(&ReminderEvent{}).Rows()
	if err != nil {
		b.Config.Logger.Warn("Reminder: Couldn't get DB rows", err)
	}

	var (
		event                   ReminderEvent
		eventDescriptionMessage string
	)

	for rows.Next() {
		err := b.DB.ScanRows(rows, &event)
		if err != nil {
			b.Config.Logger.Warn("Reminder: Couldn't get a db row", err)
		}

		eventDescriptionMessage = fmt.Sprintf(
			"%d : [ %s ] at %s UTC. Event author : %s, in channel %s, with %s",
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

// Signal yourself as interested in an event (Facebook™)
func joinEvent(message *hbot.Message) string {
	var event ReminderEvent

	id, err := getMessageIDFromString(message.Content)
	if err != nil {
		return "Wrong command. format is : !remind join [ID]"
	}

	if err := b.DB.Where("ID = ?", id).Take(&event); err != nil {
		return noAccessDBMessage()
	}

	if strings.Contains(event.People, message.Name) {
		if err := b.DB.Save(&event); err != nil {
			return noAccessDBMessage()
		}
		return "You already subscribed to this event"
	}
	event.People = fmt.Sprintf("%s%s ", event.People, message.Name)
	if err := b.DB.Save(&event); err != nil {
		return noAccessDBMessage()
	}

	feedback := fmt.Sprintf("Added %s to \"%s\"",
		message.Name,
		event.Description)

	return feedback
}

func partEvent(message *hbot.Message) string {
	var event ReminderEvent

	id, err := getMessageIDFromString(message.Content)
	if err != nil {
		return "Wrong command. format is : !remind part [ID]"
	}

	if err := b.DB.Where("ID = ?", id).Take(&event); err != nil {
		return noAccessDBMessage()
	}
	defer b.DB.Save(&event)

	if event.Author == message.Name {
		return "Author of an event can't leave it."
	}

	event.People = strings.Replace(event.People, message.Name+" ", "", -1)

	feedback := fmt.Sprintf("Removed %s from \"%s\"",
		message.Name,
		event.Description)

	return feedback
}

func parseAddCommandMessage(body string) (string, string, error) {
	var timeOfEvent string

	messageSplit := strings.SplitAfterN(body, " ", 3)
	if len(messageSplit) <= 2 {
		return "", "", &wrongFormatError{body}
	}
	body = messageSplit[2]

	timeOfEventSliced := strings.Split(body, " ")
	if len(timeOfEventSliced) > 2 {
		timeOfEvent = strings.Join(
			timeOfEventSliced[len(timeOfEventSliced)-2:],
			" ")
	} else {
		return "", "", &wrongFormatError{body}
	}

	description := strings.Replace(body, timeOfEvent, "", -1)

	return description, timeOfEvent, nil
}

func getMessageIDFromString(body string) (int, error) {
	// Parse message
	msg := strings.Split(body, " ")
	isAnID, err := regexp.MatchString("[0-9]+", msg[2])
	if err != nil {
		b.Config.Logger.Warn("Not and ID :", err)
	}

	if len(msg) != 3 || !isAnID {
		return -1, &wrongFormatError{body}
	}
	return strconv.Atoi(msg[2])
}
