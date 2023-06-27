package main

import (
	"jake/method/dev/src/bot"
	"jake/method/dev/src/bot/credentials"
	"jake/method/dev/src/chuck"
	"log"
	"strings"
)

const TCP string = "tcp"
const URL string = "irc.chat.twitch.tv:6667"

func main() {
	// Load credentials from file
	creds, err := credentials.LoadCredentials("src/bot/credentials/credentials.json")
	if err != nil {
		log.Fatalf("LoadCredentials error: %v", err)
	}

	// Connect to the Twitch IRC server
	conn, err := bot.Connect(*creds, URL)
	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}
	defer conn.Conn.Close()

	reader := conn.R

	// Start a goroutine with each connection to read incoming messages
	go func() {
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("ReadString error: %v", err)
				continue
			}

			// Keep Server Alive w/ PING/PONG
			if "PING :" == message[:5] {
				conn.Cmd("PONG :" + message[5:])
			}

			// Check if the message contains the word "!chucknorris", ignoring case
			if strings.Contains(strings.ToLower(message), chuck.ChuckMessage) {
				// Retrieve a joke from the Chuck Norris API
				joke := chuck.GetJoke(chuck.ChuckApi)

				// Send the joke to the channel
				err = conn.SendMessage("#"+creds.ChannelName, joke)
				if err != nil {
					log.Printf("SendMessage error: %v", err)
				}
			}
		}
	}()

	// Prevent main from exiting
	select {}
}
