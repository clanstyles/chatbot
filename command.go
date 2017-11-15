package ignorance

import (
	"chatbot/commands"

	twitch "github.com/gempir/go-twitch-irc"
)

type Command interface {
	Start() error
	Stop() error
	Parse(channel string, user twitch.User, message twitch.Message)
}

func Commands(c *twitch.Client) []Command {
	return []Command{
		commands.Counting(c),
	}
}
