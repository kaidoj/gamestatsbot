package commands

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
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
}

//Execute if found from contents
func Execute(c *Command) error {

	_, err := regexp.MatchString(commandPrefix, c.Content)
	if err != nil {
		return err
	}

	command := strings.Replace(c.Content, commandPrefix, "", -1)
	command = strings.Trim(command, " ")

	switch command {
	case "top":
		_, err := DisplayTopUsersByMessageCount(c)
		return err
	case "name":
		_, err := DisplayUserName(c)
		return err
	}

	if c.IsAdmin {
		switch command {
		case "admin messages":
			err := UpdateUserMessagesCount(c)
			return err
		}
	}

	return err
}
