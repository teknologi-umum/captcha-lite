package rollbar

import (
	"log"
	"os"

	"github.com/pkg/errors"
	rb "github.com/rollbar/rollbar-go"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Config struct {
	Client *rb.Client
}

func New(client *rb.Client) *Config {
	if client == nil {
		panic("rollbar.New: client is nil")
	}

	return &Config{
		Client: client,
	}
}

// HandleError handles common errors
func (c *Config) HandleError(e error) {
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Println(e)
	}

	c.Client.ErrorWithLevel(rb.ERR, errors.WithStack(e))
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
	if err != nil {
		c.Client.ErrorWithLevel(rb.ERR, errors.WithStack(err))
	}

	c.Client.ErrorWithExtras(rb.ERR, errors.WithStack(e), map[string]interface{}{
		"sender:id":       m.Sender.ID,
		"sender:name":     m.Sender.FirstName + " " + m.Sender.LastName,
		"sender:username": m.Sender.Username,
		"message:id":      m.ID,
		"message:text":    m.Text,
		"message:unix":    m.Unixtime,
	})
}
