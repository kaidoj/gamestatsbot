package main

import (
	"fmt"
	"log"

	"github.com/kaidoj/gamestatsbot/discord-bot/commands"
	"github.com/kaidoj/gamestatsbot/discord-bot/models"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	botID string
)

type Bot struct {
	config  *viper.Viper
	session *discordgo.Session
	DB      models.ConnectionInterface
}

func (b *Bot) Run() error {
	discord, err := discordgo.New("Bot " + b.config.GetString("api_key"))
	b.session = discord
	b.errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	b.errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(b.commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, b.config.GetString("status_message"))
		if err != nil {
			b.errCheck("Error attempting to set my status", err)
		}
		servers := discord.State.Guilds
		fmt.Printf(b.config.GetString("b_name")+" has started on %d servers", len(servers))
	})

	err = discord.Open()
	b.errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	<-make(chan struct{})

	return err
}

func (b *Bot) errCheck(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %+v", msg, err)
	}
}

func (b *Bot) commandHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the b is talking
		return
	}

	//create or update user message count
	userModel := &models.User{
		UserID:   user.ID,
		Username: user.Username,
	}
	b.DB.UpdateMessageCount(userModel)

	content := message.Content
	command := &commands.Command{
		Content: content,
		User:    user,
		IsAdmin: b.isAdmin(user.ID),
		Session: session,
		Message: message,
		BotID:   botID,
		DB:      b.DB,
	}
	err := command.Execute()
	if err != nil {
		log.Printf("There was problems with command %v %v", command, err)
	}
}

func (b *Bot) isAdmin(userID string) bool {
	if userID == b.config.GetString("admin_id") {
		return true
	}

	return false
}
