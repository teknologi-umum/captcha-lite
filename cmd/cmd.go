package cmd

import (
	"captcha-lite/captcha"
	"captcha-lite/locale"
	"captcha-lite/logger"

	"github.com/allegro/bigcache/v3"
	tb "gopkg.in/telebot.v3"
)

// Dependency contains the dependency injection struct
// that is required for the main command to use.
//
// It will spread and use the correct dependencies for
// each packages on the captcha project.
type Dependency struct {
	Memory   *bigcache.BigCache
	Bot      *tb.Bot
	Logger   logger.Logger
	Language string
	captcha  *captcha.Dependencies
}

// New returns a pointer struct of Dependency
// which map the incoming dependencies provided
// into what's needed by each domain.
func New(deps Dependency) *Dependency {
	var localeLanguage map[locale.Message]string
	switch deps.Language {
	case "id":
		localeLanguage = locale.ID
	default:
		localeLanguage = locale.EN
	}

	return &Dependency{
		captcha: &captcha.Dependencies{
			Memory: deps.Memory,
			Bot:    deps.Bot,
			Locale: localeLanguage,
			Log:    deps.Logger,
		},
	}
}

// OnTextHandler handle any incoming text from the group
func (d *Dependency) OnTextHandler(c tb.Context) error {
	d.captcha.WaitForAnswer(c.Message())
	return nil
}

// OnUserJoinHandler handle any incoming user join,
// whether they were invited by someone (meaning they are
// added by someone else into the group), or they join
// the group all by themselves.
func (d *Dependency) OnUserJoinHandler(c tb.Context) error {
	d.captcha.CaptchaUserJoin(c.Message())
	return nil
}

// OnNonTextHandler meant to handle anything else
// than an incoming text message.
func (d *Dependency) OnNonTextHandler(c tb.Context) error {
	d.captcha.NonTextListener(c.Message())
	return nil
}

// OnUserLeftHandler handles during an event in which
// a user left the group.
func (d *Dependency) OnUserLeftHandler(c tb.Context) error {
	d.captcha.CaptchaUserLeave(c.Message())
	return nil
}
