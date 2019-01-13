package commands

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/kaidoj/gamestatsbot/discord-bot/models"
)

var (
	commandPrefix = "!mk"
)

//Command structure
type Command struct {
	Content string
	User    *discordgo.User
	IsAdmin bool
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	BotID   string
	DB      models.ConnectionInterface
}

//Execute if found from contents
func (c *Command) Execute() error {

	_, err := regexp.MatchString(commandPrefix, c.Content)
	if err != nil {
		return err
	}

	command := strings.Replace(c.Content, commandPrefix, "", -1)
	command = strings.Trim(command, " ")

	switch command {
	case "top":
		_, err := c.DisplayTopUsersByMessageCount()
		return err
	case "name":
		_, err := c.DisplayUserName()
		return err
	}

	if c.IsAdmin {
		switch command {
		case "admin messages":
			err := c.UpdateUserMessagesCount()
			return err
		}
	}

	return err
}
