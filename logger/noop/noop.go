package noop

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

type Config struct{}

func New() *Config {
	return &Config{}
}

// HandleError handles common errors.
func (c *Config) HandleError(e error) {
	return
}

// HandleBotError is the handler for an error which a function has a
// bot and a message instance.
//
// For other errors that don't have one of those struct instance, use
// HandleError instead.
func (c *Config) HandleBotError(e error, bot *tb.Bot, m *tb.Message) {
	return
}
