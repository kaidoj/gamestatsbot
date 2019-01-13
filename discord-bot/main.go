package main

import (
	"github.com/kaidoj/gamestatsbot/discord-bot/config"
	"github.com/kaidoj/gamestatsbot/discord-bot/models"
)

func main() {
	cnf := config.Init(".")
	db := &models.DB{nil, cnf}
	bot := &Bot{cnf, nil, db.Init()}
	bot.Run()
}
