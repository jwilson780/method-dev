package main

import (
	"bufio"
	"fmt"
	"jake/method/dev/src/bot"
	"jake/method/dev/src/chuck"
	"os"
	"strings"
	"time"
)


func main() {

	// connect to channel
	conn, err := bot.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %s\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(conn.R)

	for {
		// read in line from terminal
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read: %s\n", err)
			os.Exit(1)
		}

		line = strings.TrimSuffix(line, "\r\n")

		// if line starts with PING, send PONG
		if strings.HasPrefix(line, "PING") {
			_, err = conn.Cmd("PONG %s", strings.Split(line, " ")[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write PONG: %s\n", err)
				os.Exit(1)
			}
			continue
		}

		if !strings.Contains(line, "PRIVMSG") {
			continue
		}

		msg := bot.ParseMessage(line)
		// print message to terminal using time format RFC3339
		fmt.Printf("[%s] %s: %s\n", time.Now().Format(time.RFC3339), msg.User, msg.Content)

		// if message is "!chucknorris", send joke to channel
		if msg.Content == chuck.ChuckMessage {
			joke := chuck.GetJoke()
			_, err = conn.Cmd("PRIVMSG #%s :%s", "YOUR_CHANNEL_NAME", joke)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write message: %s\n", err)
				os.Exit(1)
			}
		}
	}
}
