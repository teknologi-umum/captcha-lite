package logger

import (
	tb "gopkg.in/telebot.v3"
)

type Logger interface {
	// HandleError handles common errors.
	HandleError(e error)
	// HandleBotError is the handler for an error which a function has a
	// bot and a message instance.
	//
	// For other errors that don't have one of those struct instance, use
	// HandleError instead.
	HandleBotError(e error, bot *tb.Bot, m *tb.Message)
}
