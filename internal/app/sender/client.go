package sender

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Subscribe(bot *tgbotapi.BotAPI, c chan int64) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		c <- update.Message.Chat.ID
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func Write(bot *tgbotapi.BotAPI, c chan int64, fn func()) {
	for {
		select {
		case chatID := <-c:
			fmt.Println("received", chatID)
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Chat %v; Message:%v", chatID, "Bot set"))
			bot.Send(msg)
			fn()
		default:
			fmt.Println("*******************")
		}
		time.Sleep(time.Second * 1)
	}
}
