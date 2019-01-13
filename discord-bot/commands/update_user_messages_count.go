package commands

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kaidoj/gamestatsbot/discord-bot/models"
)

var (
	fetchMessagesCount = 100
	wg                 sync.WaitGroup
)

//UpdateUserMessagesCount update users message count from discord api
func (c *Command) UpdateUserMessagesCount() error {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
		fmt.Println("Ended")
	}()

	var err error
	messagesCh := make(chan []*discordgo.Message)

	wg.Add(1)
	go c.getUsersMessages(messagesCh, "")

	usersList := c.countUsersMessages(messagesCh)
	wg.Wait()

	//reset users message counts to 0
	c.DB.ResetMessagesCount()

	for _, u := range usersList {
		res := c.DB.CreateOrUpdate(u)
		if res != nil {
			err = res
		}
	}

	return err
}

func (c *Command) getUsersMessages(messagesCh chan []*discordgo.Message, beforeID string) {
	defer wg.Done()
	messages, err := c.Session.ChannelMessages(c.Message.ChannelID, fetchMessagesCount, beforeID, "", "")
	if err != nil {
		log.Printf("Error fetching channel messages: %v", err)
	}

	messagesCh <- messages

	messagesLen := len(messages)
	if messagesLen == fetchMessagesCount {
		wg.Add(1)
		go c.getUsersMessages(messagesCh, messages[messagesLen-1].ID)
	} else {
		close(messagesCh)
	}

}

func (c *Command) countUsersMessages(messagesCh chan []*discordgo.Message) map[string]*models.User {
	users := make(map[string]*models.User)

	for messages := range messagesCh {
		for _, m := range messages {
			author := m.Author
			if author.ID == c.BotID || author.Bot {
				continue
			}

			if userExists, ok := users[author.ID]; ok {
				userExists.MessageCount++
				users[author.ID] = userExists
			} else {
				users[author.ID] = &models.User{
					UserID:       author.ID,
					MessageCount: 1,
					Username:     author.Username,
				}
			}
		}
	}

	return users
}
