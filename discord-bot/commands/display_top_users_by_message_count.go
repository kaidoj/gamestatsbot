package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

//DisplayTopUsersByMessageCount will display top users by message count
func (c *Command) DisplayTopUsersByMessageCount() (*discordgo.Message, error) {
	users := c.DB.GetUsersByMessageCount(10)

	fields := []*discordgo.MessageEmbedField{}
	for _, u := range users {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   u.Username,
			Value:  strconv.Itoa(u.MessageCount) + " sõnumit",
			Inline: false,
		})
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00, // Green
		Fields: fields,
		Title:  "Top 10 kasutajad",
	}

	m, err := c.Session.ChannelMessageSendEmbed(c.Message.ChannelID, embed)

	return m, err
}
