package app

import (
	"database/sql"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/marktsoy/gomonolith_sample/internal/app/apiserver"
	"github.com/marktsoy/gomonolith_sample/internal/app/models"
	"github.com/marktsoy/gomonolith_sample/internal/app/sender"
	"github.com/marktsoy/gomonolith_sample/internal/app/store/sqlstore"
	"github.com/marktsoy/gomonolith_sample/internal/app/subscriber"
)

type Application struct {
	config     *Config
	bus        *Bus
	bot        *tgbotapi.BotAPI
	subscriber *subscriber.Subscriber
	sender     *sender.Sender
	server     *apiserver.Server
}

func New(config *Config) *Application {
	app := &Application{
		config: config,
	}
	app.configureBot()
	app.configureBus()
	app.configureSender()
	app.configureDispacher()
	app.configureServer()
	return app
}

func (a *Application) configureBus() {
	a.bus = NewBus()
}

func (a *Application) configureBot() {
	bot, err := tgbotapi.NewBotAPI(a.config.TgKey)
	if err != nil {
		panic(err.Error())
	}
	a.bot = bot
}

func (a *Application) configureDispacher() {
	a.subscriber = subscriber.New(a.bot)
}

func (a *Application) configureSender() {
	a.sender = sender.New(a.bot)
}

func (a *Application) configureServer() {
	db, err := sql.Open("postgres", a.config.Store)
	if err != nil {
		panic(err)
	}
	a.server = apiserver.New(a.bus.Messages, sqlstore.New(db))
}

func (a *Application) listenToChats() {
	for {
		select {
		case act := <-a.bus.Chats:
			switch act.Name {
			default:
				continue
			case "subscribe":
				a.sender.AddChat(act.Payload.(int64))
			case "unsubscribe":
				a.sender.RemoveChat(act.Payload.(int64))

			}
		}
	}
}

func (a *Application) listenToMessages() {
	for {
		select {
		case act := <-a.bus.Messages:
			switch act.Name {
			default:
				continue
			case "messageCreated":
				a.sender.AddMessage(act.Payload.(*models.Message))
			}
		}
	}
}

func (a *Application) Run() {
	go func() { a.subscriber.Run(a.bus.Chats) }()
	go func() { a.sender.Run() }()
	go a.listenToChats()
	go a.listenToMessages()
	http.ListenAndServe(a.config.BindAddr, a.server)
}
