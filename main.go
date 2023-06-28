package main

import (
	"jake/method/dev/src/bot"
	"jake/method/dev/src/bot/credentials"
	"jake/method/dev/src/chuck"
	"log"
	"strings"
)

const tlsUrl string = "irc.chat.twitch.tv:6697"
const chuckApi string = "https://api.chucknorris.io/jokes/random"
const chuckMessage string = "!chucknorris"
const credentialsPath string = "src/bot/credentials/credentials.json"

func main() {
	// Load credentials from file
	creds, err := credentials.LoadCredentials(credentialsPath)
	if err != nil {
		log.Fatalf("LoadCredentials error: %v", err)
	}

	// Connect to the Twitch IRC server
	conn, err := bot.Connect(*creds, tlsUrl)
	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}
	defer conn.Conn.Close()
	// Start a goroutine with each connection to read incoming messages
	go func() {
		for {
			message, err := conn.R.ReadString('\n')
			if err != nil {
				log.Printf("Failed to read message: %v", err)
				return
			}

			message = strings.TrimSuffix(message, "\n")
			log.Printf("Received: %s", message)

			if strings.Contains(message, "PING") {
				conn.Cmd("PONG :tmi.twitch.tv")
				log.Println("Responded to PING with PONG")
			}

			if strings.Contains(message, chuckMessage) {
				go func() {
					log.Printf("Received %s command", chuckMessage)
					joke := chuck.GetJoke(chuckApi)
					err := conn.SendMessage("#"+creds.ChannelName, joke)
					if err != nil {
						log.Printf("Failed to send message: %v", err)
					} else {
						log.Printf("Sent: %s", joke)
					}
				}()
			}
		}
	}()

	// Prevent main from exiting
	select {}
}
