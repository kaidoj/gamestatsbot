package commands

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	UserModel "github.com/kaidoj/gamestatsbot/discord-bot/models/user"
)

var (
	fetchMessagesCount = 100
	wg                 sync.WaitGroup
)

//UpdateUserMessagesCount update users message count from discord api
func UpdateUserMessagesCount(c *Command) error {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
		fmt.Println("Ended")
	}()

	var err error
	messagesCh := make(chan []*discordgo.Message)

	wg.Add(1)
	go getUsersMessages(c, "", messagesCh)

	usersList := countUsersMessages(c, messagesCh)
	wg.Wait()

	//reset users message counts to 0
	UserModel.ResetMessagesCount()

	for _, u := range usersList {
		res := u.CreateOrUpdate()
		if res != nil {
			err = res
		}
	}

	return err
}

func getUsersMessages(c *Command, beforeID string, messagesCh chan []*discordgo.Message) {
	defer wg.Done()
	messages, err := c.Session.ChannelMessages(c.Message.ChannelID, fetchMessagesCount, beforeID, "", "")
	if err != nil {
		log.Printf("Error fetching channel messages: %v", err)
	}

	messagesCh <- messages

	messagesLen := len(messages)
	if messagesLen == fetchMessagesCount {
		wg.Add(1)
		go getUsersMessages(c, messages[messagesLen-1].ID, messagesCh)
	} else {
		close(messagesCh)
	}

}

func countUsersMessages(c *Command, messagesCh chan []*discordgo.Message) map[string]*UserModel.User {
	users := make(map[string]*UserModel.User)

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
				users[author.ID] = &UserModel.User{
					UserID:       author.ID,
					MessageCount: 1,
					Username:     author.Username,
				}
			}
		}
	}

	return users
}
