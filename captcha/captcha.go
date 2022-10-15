package captcha

import (
	"teknologi-umum-bot/locale"
	"teknologi-umum-bot/logger"

	"github.com/allegro/bigcache/v3"
	tb "gopkg.in/telebot.v3"
)

// Dependencies contains the dependency injection struct for
// methods in the captcha package.
type Dependencies struct {
	Memory *bigcache.BigCache
	Bot    *tb.Bot
	Log    logger.Logger
	Locale map[locale.Message]string
}
