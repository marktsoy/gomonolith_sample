package sender_test

import (
	"log"
	"sync"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/marktsoy/gomonolith_sample/internal/app/sender"
)

func TestTG_SendTest(t *testing.T) {
	t.Helper()
	c := make(chan int64)

	bot, err := tgbotapi.NewBotAPI("1392295365:AAHHzzpaO8fCVyBq_oI8v6Yqi5OqjfAGONo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	var wg sync.WaitGroup
	go func() {
		sender.Subscribe(bot, c)
	}()
	wg.Add(10)
	go func() {
		sender.Write(bot, c, func() {
			wg.Done()
		})
	}()
	wg.Wait()
}
