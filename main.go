package main

import (
	"jake/method/dev/internal/chuck"
	"jake/method/dev/internal/conn"
	"jake/method/dev/internal/conn/credentials"
	"log"
	"strings"
)

const (
	tlsUrl          string = "irc.chat.twitch.tv:6697"
	chuckApi        string = "https://api.chucknorris.io/jokes/random"
	chuckMessage    string = "!chucknorris"
	credentialsPath string = "internal/conn/credentials/credentials.json"
	ping string = "PING"
	pongMessage string = "PONG :tmi.twitch.tv"
)

func main() {
	// Load credentials from local file
	creds, err := credentials.LoadCredentials(credentialsPath)
	if err != nil {
		log.Fatalf("LoadCredentials error: %v", err)
	}

	// Connect to the Twitch IRC server
	conn, err := conn.Connect(*creds, tlsUrl)
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

			if strings.Contains(message, ping) {
				_, err := conn.Cmd(pongMessage)
				if err != nil {
					log.Printf("Failed to send PONG message: %v", err)
					return
				}
				log.Println("Responded to PING with PONG")
			}

			if strings.Contains(message, chuckMessage) {
				go func() {
					log.Printf("Received %s command", chuckMessage)
					joke := chuck.GetJoke(chuckApi)
					err := conn.SendMessage("#"+creds.ChannelName, joke)
					if err != nil {
						log.Printf("Failed to send message: %v", err)
						return
					}
					log.Printf("Sent: %s", joke)

				}()
			}
		}
	}()

	// Prevent main from exiting
	select {}
}
