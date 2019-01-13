package commands

import "github.com/bwmarrin/discordgo"

//DisplayUserName to client
func (c *Command) DisplayUserName() (*discordgo.Message, error) {
	m, err := c.Session.ChannelMessageSend(c.Message.ChannelID, c.User.Username+" <- Mida nime IRV")
	return m, err
}
