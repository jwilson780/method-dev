package bot

import (
	"strings"
)

const TwitchUrl string = "irc.chat.twitch.tv:6697"

type Message struct {
	User    string
	Content string
}

// ParseMessage parses a raw message from Twitch IRC
func ParseMessage(raw string) Message {
	parts := strings.Split(raw, " ")

	user := strings.Split(parts[1], "!")[0]
	content := strings.Join(parts[3:], " ")[1:]

	return Message{User: user, Content: content}
}
