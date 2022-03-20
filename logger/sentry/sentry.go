package sentry

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

type Config struct {
	Client *sentry.Client
}

func New(client *sentry.Client) *Config {
	if client == nil {
		panic("sentry.New: client is nil")
	}

	return &Config{
		Client: client,
	}
}

// HandleError handles common errors.
func (c *Config) HandleError(e error) {
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Println(e)
	}

	_ = c.Client.CaptureException(
		errors.WithStack(e),
		&sentry.EventHint{OriginalException: e},
		nil,
	)
}

// HandleBotError is the handler for an error which a function has a
// bot and a message instance.
//
// For other errors that don't have one of those struct instance, use
// HandleError instead.
func (c *Config) HandleBotError(e error, bot *tb.Bot, m *tb.Message) {
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Println(e)
	}

	_, err := bot.Send(
		m.Chat,
		"Oh no, something went wrong with me! Can you guys help me to ping my masters?",
		&tb.SendOptions{ParseMode: tb.ModeHTML},
	)

	scope := sentry.NewScope()
	scope.SetContext("tg:sender", map[string]interface{}{
		"id":       m.Sender.ID,
		"name":     m.Sender.FirstName + " " + m.Sender.LastName,
		"username": m.Sender.Username,
	})
	scope.SetContext("tg:message", map[string]interface{}{
		"id":   m.ID,
		"text": m.Text,
		"unix": m.Unixtime,
	})

	if err != nil {
		// Come on? Another error?
		_ = c.Client.CaptureException(
			errors.WithStack(err),
			&sentry.EventHint{OriginalException: err},
			scope,
		)
	}

	_ = c.Client.CaptureException(
		errors.WithStack(e),
		&sentry.EventHint{OriginalException: e},
		scope,
	)
}
