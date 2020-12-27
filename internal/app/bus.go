package app

import "github.com/marktsoy/gomonolith_sample/internal/app/utils"

type Bus struct {
	Chats    chan *utils.Action
	Messages chan *utils.Action
}

func NewBus() *Bus {
	return &Bus{
		Chats:    make(chan *utils.Action),
		Messages: make(chan *utils.Action),
	}
}
