package main

import (
	"log"
	"os"

	"github.com/kaidoj/gamestatsbot/discord-bot/config"
	"github.com/kaidoj/gamestatsbot/discord-bot/models"
)

func main() {
	log.SetOutput(os.Stdout)
	cnf := config.Init(".")
	db := &models.DB{nil, cnf}
	bot := &Bot{cnf, nil, db.Init()}
	bot.Run()
}
