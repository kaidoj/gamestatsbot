package commands

import (
	"github.com/bwmarrin/discordgo"
)

func (c *Command) DisplayCommands() (*discordgo.Message, error) {
	fields := []*discordgo.MessageEmbedField{}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "!mk top",
		Value:  "Top kasutajad sõnumite arvu kaupa",
		Inline: false,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "!mk name",
		Value:  "Lihtne nime kuvamine",
		Inline: false,
	})

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "!mk releases",
		Value:  "Uute mängude väljalaske kuupäevad",
		Inline: false,
	})

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00, // Green
		Fields: fields,
		Title:  "Käsud",
	}

	m, err := c.Session.ChannelMessageSendEmbed(c.Message.ChannelID, embed)

	return m, err
}
