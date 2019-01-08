package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	Model "github.com/kaidoj/gamestatsbot/discord-bot/models/user"
)

//DisplayTopUsersByMessageCount will display top users by message count
func DisplayTopUsersByMessageCount(c *Command) (*discordgo.Message, error) {
	userModel := &Model.User{}
	users := userModel.GetUsersByMessageCount(10)

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
