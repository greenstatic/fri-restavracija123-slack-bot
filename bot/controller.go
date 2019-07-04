package bot

import (
	"github.com/greenstatic/fri-restavracija123-slack-bot/restavracija123"
	"github.com/greenstatic/fri-restavracija123-slack-bot/slack"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"time"
)

type Bot struct {
	config slack.Config

	// When we cause the bot to send the message with an up to 1 minute delay
	dailyTrigger time.Time
	lastRun      time.Time
}

func New(c slack.Config, t time.Time) Bot {
	b := Bot{}
	b.config = c
	b.dailyTrigger = t
	return b
}

func (b *Bot) Start() {
	const secondInterval = 1

	logrus.WithField("secondInterval", secondInterval).Info("Starting bot")

	ticker := time.NewTicker(secondInterval * time.Second)

	for _ = range ticker.C {
		// Infinite loop

		logrus.Debug("Checking if we can trigger Slack message")
		if b.canRun() {

			logrus.Info("Allowed to publish message")
			if err := b.run(); err != nil {
				logrus.WithField("error", err).Warning("Failed to trigger Slack message, trying again in a short while.")
				time.Sleep(3 * time.Second)
			} else {
				// Mutate state
				b.lastRun = time.Now()

				logrus.Info("Successfully triggered Slack message")
			}

		}

	}
}

func (b Bot) todayDailyTrigger() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), b.dailyTrigger.Hour(), b.dailyTrigger.Minute(), b.dailyTrigger.Second(), 0, n.Location())
}

func (b Bot) canRun() bool {
	trigger := b.todayDailyTrigger()
	durr := time.Now().Sub(trigger)
	if durr < 0 {
		logrus.WithField("duration", durr.String()).Info("Running in")
	} else {
		logrus.Debug("Done for today.")
	}

	// Last run enables us to record if we already have performed a successful operation
	// today. This enables us to retry if an error is presented. (We have a 1 minute period
	// where we will retry).
	return time.Now().Sub(b.lastRun).Hours() > 1 && durr > 0 && durr < time.Minute
}

func (b Bot) run() error {
	foods, err := restavracija123.DailyMenu(time.Now())
	if err != nil {
		return stacktrace.Propagate(err, "Cannot run bot due to data source API failure")
	}

	if len(foods) == 0 {
		logrus.Info("No foods on the menu today, not sending any message.")
		return nil
	}

	msg := MenuMarkdownContent(foods)

	if err := b.config.SendMessage(msg); err != nil {
		return stacktrace.Propagate(err, "Cannot run due to failure to send Slack message")
	}

	return nil
}
