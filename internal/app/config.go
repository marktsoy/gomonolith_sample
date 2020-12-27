package app

type Config struct {
	Store    string
	TgKey    string
	BindAddr string
}

func NewConfig() *Config {
	return &Config{
		Store:    "postgres://localhost/sandbox?sslmode=disable",
		TgKey:    "1392295365:AAHHzzpaO8fCVyBq_oI8v6Yqi5OqjfAGONo",
		BindAddr: "localhost:3333",
	}
}
