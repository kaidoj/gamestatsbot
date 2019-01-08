package commands

import "github.com/bwmarrin/discordgo"

//DisplayUserName to client
func DisplayUserName(c *Command) (*discordgo.Message, error) {
	m, err := c.Session.ChannelMessageSend(c.Message.ChannelID, c.User.Username+" <- Mida nime IRV")
	return m, err
}
