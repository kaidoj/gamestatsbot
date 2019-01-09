package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"

	"github.com/kaidoj/gamestatsbot/discord-bot/commands"
	Model "github.com/kaidoj/gamestatsbot/discord-bot/models/user"
)

var (
	botID string
)

type Bot struct {
	config  *viper.Viper
	session *discordgo.Session
}

func (b *Bot) Run() {
	b.connect()
}

func (b *Bot) connect() {
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
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf(b.config.GetString("b_name")+" has started on %d servers", len(servers))
	})

	err = discord.Open()
	b.errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	<-make(chan struct{})
}

func (b *Bot) errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
	}
}

func (b *Bot) commandHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the b is talking
		return
	}

	//create or update user message count
	userModel := &Model.User{
		UserID:   user.ID,
		Username: user.Username,
	}
	userModel.UpdateMessageCount()

	content := message.Content
	command := &commands.Command{
		Content: content,
		User:    user,
		IsAdmin: b.isAdmin(user.ID),
		Session: session,
		Message: message,
		BotID:   botID,
	}
	err := commands.Execute(command)
	if err != nil {
		b.errCheck("There was problems with some commands", err)
	}
}

func (b *Bot) isAdmin(userID string) bool {
	if userID == b.config.GetString("admin_id") {
		return true
	}

	return false
}
