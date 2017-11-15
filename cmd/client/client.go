package main

import (
	"flag"
	"log"

	"chatbot"

	twitch "github.com/gempir/go-twitch-irc"
)

var (
	counter int

	username string
	oauth    string
)

func init() {
	flag.StringVar(&username, "username", "", "")
	flag.StringVar(&oauth, "oauth", "", "")
	flag.Parse()
}

func main() {
	client := twitch.NewClient(username, oauth)

	client.OnNewMessage(func(channel string, user twitch.User, message twitch.Message) {
		log.Println(user.Username, "]", message.Text)

		for _, command := range ignorance.Commands(client) {
			command.Parse(channel, user, message)
		}
	})

	client.Join("hardlydifficult")

	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
}
