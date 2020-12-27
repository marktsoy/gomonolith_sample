package sender

import (
	"container/heap"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/marktsoy/gomonolith_sample/internal/app/models"
)

type Sender struct {
	bot      *tgbotapi.BotAPI
	mutex    sync.Mutex
	messages *PriorityQueue
	chats    []int64
}

func New(bot *tgbotapi.BotAPI) *Sender {
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	return &Sender{
		bot:      bot,
		messages: &pq,
		chats:    make([]int64, 0),
		mutex:    sync.Mutex{},
	}
}

func (sender *Sender) Run() {
	for {
		time.Sleep(time.Second * 1)
		if sender.messages.Len() > 0 {
			sender.mutex.Lock()
			item := heap.Pop(sender.messages).(*Item)
			if item != nil {
				sender.send(item.value)
			}
			sender.mutex.Unlock()
		}
	}
}

func (sender *Sender) AddMessage(m *models.Message) {
	sender.mutex.Lock()
	item := &Item{
		value:    m.Content,
		priority: m.Priority,
	}
	heap.Push(sender.messages, item)
	sender.mutex.Unlock()
}

func (sender *Sender) AddChat(chatID int64) {
	sender.mutex.Lock()
	sender.chats = append(sender.chats, chatID)
	sender.mutex.Unlock()
}

func (sender *Sender) RemoveChat(chatID int64) {
	sender.mutex.Lock()
	for i, v := range sender.chats {
		if v == chatID {
			sender.chats = append(sender.chats[:i], sender.chats[i+1:]...) // Spread operator ???
			break
		}
	}
	sender.mutex.Unlock()
}

func (sender *Sender) send(text string) {
	for _, chatID := range sender.chats {
		msg := tgbotapi.NewMessage(chatID, text)
		sender.bot.Send(msg)
	}
}
