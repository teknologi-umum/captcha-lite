package rollbar

import (
	"fmt"

	"github.com/rs/zerolog"
	tb "gopkg.in/telebot.v3"
)

type sender struct {
	Id       int64
	Name     string
	Username string
}

func (s sender) MarshalZerologObject(e *zerolog.Event) {
	e.Int64("id", s.Id).
		Str("name", s.Name).
		Str("username", s.Username)
}

type message struct {
	Id       int
	Text     string
	Unixtime int64
}

func (msg message) MarshalZerologObject(e *zerolog.Event) {
	e.Int("id", msg.Id).
		Str("text", msg.Text).
		Int64("unixtime", msg.Unixtime)
}

type Config struct {
	Log zerolog.Logger
}

func New(log zerolog.Logger) *Config {
	return &Config{
		Log: log,
	}
}

// HandleError handles common errors
func (c *Config) HandleError(e error) {
	c.Log.Error().Err(e).Stack().Msg("")
	return
}

// HandleBotError is the handler for an error which a function has a
// bot and a message instance.
//
// For other errors that don't have one of those struct instance, use
// HandleError instead.
func (c *Config) HandleBotError(e error, bot *tb.Bot, m *tb.Message) {
	_, err := bot.Send(
		m.Chat,
		"Oh no, something went wrong with me! Can you guys help me to ping my masters?",
		&tb.SendOptions{ParseMode: tb.ModeHTML},
	)

	if err != nil {
		c.Log.Error().Err(err).Stack().Msg("")
	}

	s := sender{
		Id:       m.Sender.ID,
		Name:     fmt.Sprintf("%s %s", m.Sender.FirstName, m.Sender.LastName),
		Username: m.Sender.Username,
	}

	msg := message{
		Id:       m.ID,
		Text:     m.Text,
		Unixtime: m.Unixtime,
	}

	c.Log.
		Error().
		Err(e).
		Stack().
		Object("sender", s).
		Object("message", msg).
		Msg("")
}
