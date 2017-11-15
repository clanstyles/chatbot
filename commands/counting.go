package commands

import (
	"fmt"
	"log"
	"strconv"
	"time"

	twitch "github.com/gempir/go-twitch-irc"
)

type CountingHighScore struct {
	Score   int
	User    string
	Created time.Time
}

type countingCommand struct {
	*twitch.Client
}

type countingState struct {
	count     int
	lastUser  string
	HighScore CountingHighScore
}

var (
	counting countingState
)

func Counting(c *twitch.Client) *countingCommand {
	return &countingCommand{
		Client: c,
	}
}

func (c *countingCommand) Start() error {
	return nil
}
func (c *countingCommand) Stop() error {
	return nil
}

func (c *countingCommand) Parse(channel string, user twitch.User, message twitch.Message) {
	if message.Text == "!hc" {
		if counting.HighScore.Score == 0 {
			c.Say(channel, "Nobody has the current high score.")
			return
		}

		since := time.Since(counting.HighScore.Created).String()
		c.Say(channel, fmt.Sprintf("%s has the current high score of %d. He's held the top position for %s", counting.HighScore.User, counting.HighScore.Score, since))
		return
	}

	val, err := strconv.Atoi(message.Text)
	if err != nil {
		return
	}

	log.Println("we have a number", val, "and counter", counting.count)
	if counting.count == 0 {
		if val == 1 {
			counting.count++
			c.Say(channel, fmt.Sprintf("%s started the counting game!", user.Username))
			return
		}
	}

	if counting.count > 0 {
		if counting.count+1 == val {
			counting.count++
			counting.lastUser = user.Username
		} else {
			log.Println("game is over")
			log.Println(counting.count, "is >", counting.HighScore.Score)

			if counting.count > counting.HighScore.Score {
				c.Say(channel, fmt.Sprintf("%s broke the count, but, %s has the new high score of %d!", user.Username, counting.lastUser, counting.count))
				log.Println(fmt.Sprintf("%s broke the count, but, %s has the new high score of %d!", user.Username, counting.lastUser, counting.count))
				counting.HighScore = CountingHighScore{
					Score:   counting.count,
					User:    counting.lastUser,
					Created: time.Now().UTC(),
				}
			} else {
				log.Println(fmt.Sprintf("%s broke the count of %d", user.Username, counting.count))
				c.Say(channel, fmt.Sprintf("%s broke the count of %d", user.Username, counting.count))
			}

			counting.count = 0
		}
	}
}
