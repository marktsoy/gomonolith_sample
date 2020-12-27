package subscriber

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/marktsoy/gomonolith_sample/internal/app/utils"
)

type Subscriber struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Subscriber {
	return &Subscriber{
		bot: bot,
	}
}

func (s *Subscriber) Run(c chan *utils.Action) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := s.bot.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		var a *utils.Action

		if update.Message.Text == "/subscribe" {
			a = &utils.Action{
				Name:    "subscribe",
				Payload: update.Message.Chat.ID,
			}
		}
		if update.Message.Text == "/unsubscribe" {
			a = &utils.Action{
				Name:    "unsubscribe",
				Payload: update.Message.Chat.ID,
			}
		}
		if a != nil {
			c <- a
		}
	}
}
