package main

import (
	"github.com/kaidoj/gamestatsbot/discord-bot/config"
)

func main() {
	initBot := &Bot{config.Init(), nil}
	initBot.Run()
}
