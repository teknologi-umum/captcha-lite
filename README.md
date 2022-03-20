# Very lightweight Telegram captcha bot written in Go

Some captcha bot on Telegram requires you to pay to be used on groups with 1000+ members.
Some is just so bad, bots can bypass through them.

We made something that's can't be bypassed by bots and free of cost.

The [full-version one](https://github.com/teknologi-umum/captcha) has more dependencies,
database connections (CockroachDB, MongoDB, and Redis), and a constant worker that logs
the group. Yet, it only runs on a free VM on [fly.io](https://fly.io/docs/about/pricing/)
which only has 1 shared vCPU, 3GB volume storage, and 256 MB of RAM.

Well, the point is: **if the full-blown version is running on a very low-end VM and still
be able to handle lots and lots of traffic (yes, we had an encounter of spam bots joining
our server before), this one should use even less resources!**

## Deployment

Before you deploy the bot, you'll need to setup some environment variables to make
everything works. See [Environment Variables](#environment-variables) section.

The easiest way to deploy the bot is through Docker.

```
docker build -t teknologi-umum/captcha-lite .
docker run teknologi-umum/captcha-lite
```

You can use [Heroku](https://www.heroku.com/) or [Fly.io](https://fly.io/) to deploy the app.
We've provided the corresponding configuration file.

Or if you prefer to build everything from source.

```
# We assume you've installed Go 1.18 or higher (https://go.dev/dl)
go build .
./teknologi-umum-bot
```

Or run the `teknologi-umum-bot` binary file through systemd or supervisord.

## Environment Variables

- ENVIRONMENT: denotes the environment stage.
  Available options: "development" / "production"
- BOT_TOKEN: Your Telegram bot token that is acquired from BotFather
- LANGUAGE: The language of the bot.
  Available options: "ID" (for Indonesian) / "EN" (for English)
  Defaults to "EN"
- LOG_PROVIDER: Error log provider.
  Available options:
    - "noop" -- stands for no-operation. It literally do nothing.
    - "sentry" -- See https://sentry.io/
    - "rollbar" -- See https://rollbar.com/
- SENTRY_DSN: Sentry's DSN URL. Required if using "sentry" as the LOG_RPOVIDER
- ROLLBAR_TOKEN: Rollbar's token. Required if using "rollbar" as the LOG_PROVIDER
- ROLLBAR_SERVERHOST: Rollbar's server host. Required if using "rollbar" as the LOG_PROVIDER
- ROLLBAR_SERVERROOT: Rollbar's server root. Required if using "rollbar" as the LOG_PROVIDER

## License

```
MIT License

Copyright © 2022 Teknologi Umum <opensource@teknologiumum.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
