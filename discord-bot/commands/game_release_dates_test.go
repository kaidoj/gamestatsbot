package commands

import (
	"os"
	"path"
	"testing"

	"github.com/Henry-Sarabia/igdb"

	"github.com/kaidoj/gamestatsbot/discord-bot/config"
)

func TestFetchGameReleaseTimes(t *testing.T) {
	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/kaidoj/gamestatsbot/discord-bot")
	config := config.Init(ap)
	if config == nil {
		t.Errorf("Config not initialized")
	}

	c := &Command{}
	c.Config = config
	_, err := c.fetchGameReleaseTimes()
	if err != nil {
		t.Errorf("Release times could not be fetched")
	}
}

func TestFormatGamesByMonth(t *testing.T) {
	game := &igdb.Game{}
	game.Name = "test"
	game.FirstReleaseDate = 1549655365 //february
	games := []*igdb.Game{game}

	game2 := &igdb.Game{}
	game2.Name = "test"
	game2.FirstReleaseDate = 1551398400 //march
	games = append(games, game2)

	game3 := &igdb.Game{}
	game3.Name = "test"
	game3.FirstReleaseDate = 1551398400 //march
	games = append(games, game3)

	formatedGames := formatGamesByMonth(games)

	i := 0
	for k, games := range formatedGames {

		if i == 0 {
			if k != "February 2019" {
				t.Errorf("Game list doesn't contain February 2019")
			}

			for _, v := range games {
				if v == "" {
					t.Errorf("No games found in February 2019")
				}
			}
		}

		if i == 1 {
			if k != "March 2019" {
				t.Errorf("Game list doesn't contain March 2019")
			}

			for _, v := range games {
				if v == "" {
					t.Errorf("No games found in March 2019")
				}
			}
		}

		i++
	}
}

func TestParseReleaseDate(t *testing.T) {
	res := parseReleaseDate(1549655365)
	if res != "08.02.2019" {
		t.Errorf("Could not parse release date")
	}
}

func TestParseParseMonth(t *testing.T) {
	res := parseMonth(1549655365)
	if res != "February 2019" {
		t.Errorf("Could not parse month")
	}
}
