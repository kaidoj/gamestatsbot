package commands

import (
	"log"
	"strconv"
	"time"

	"github.com/Henry-Sarabia/igdb"
	"github.com/bwmarrin/discordgo"
)

func (c *Command) GetGameReleaseDates() (*discordgo.Message, error) {

	fields := []*discordgo.MessageEmbedField{}

	PC, err := c.fetchGameReleaseTimes()
	monthsWithGames := formatGamesByMonth(PC)

	for key, games := range monthsWithGames {

		var message string
		i := 1
		for _, v := range games {
			message += v + "\n "

			if i == 25 {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:   key,
					Value:  message,
					Inline: false,
				})
				i = 0
				message = ""
			}

			i++
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   key,
			Value:  message,
			Inline: false,
		})
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00, // Green
		Fields: fields,
		Title:  "Uued mängud tulemas",
	}

	m, err := c.Session.ChannelMessageSendEmbed(c.Message.ChannelID, embed)
	return m, err
}

//fetch release times from external api
func (c *Command) fetchGameReleaseTimes() ([]*igdb.Game, error) {
	client := igdb.NewClient(c.Config.GetString("igdb_api_key"), nil)

	today := time.Now()
	months := today.AddDate(0, 4, 0)

	byFirstReleaseDate := igdb.ComposeOptions(
		igdb.SetLimit(50),
		igdb.SetFields("name", "first_release_date"),
		igdb.SetOrder("first_release_date", igdb.OrderAscending),
		igdb.SetFilter("first_release_date", igdb.OpGreaterThanEqual, strconv.FormatInt(today.Unix(), 10)),
		igdb.SetFilter("first_release_date", igdb.OpLessThanEqual, strconv.FormatInt(months.Unix(), 10)),
	)

	PC, err := client.Games.List(
		nil,
		byFirstReleaseDate,
		igdb.SetFilter("platforms", igdb.OpIn, "6"),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return PC, nil
}

//format so they are in their month group
func formatGamesByMonth(games []*igdb.Game) map[string][]string {
	newGames := make(map[string][]string)
	for _, v := range games {
		month := parseMonth(v.FirstReleaseDate)
		gameName := v.Name + " (" + parseReleaseDate(v.FirstReleaseDate) + ")"
		newGames[month] = append(newGames[month], gameName)
	}
	return newGames
}

func parseReleaseDate(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("02.01.2006")
}

func parseMonth(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("January 2006")
}
