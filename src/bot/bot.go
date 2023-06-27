package bot

import (
	"fmt"
	"jake/method/dev/src/bot/credentials"
	"net/textproto"
	"strings"
)

const TwitchUrl string = "irc.chat.twitch.tv:6697"
const TCP string = "tcp"

type Message struct {
	User    string
	Content string
}

func Connect() (*textproto.Conn, error) {
	creds, err := credentials.LoadCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to load credentials: %w", err)
	}

	conn, err := textproto.Dial(TCP, TwitchUrl)
	if err != nil {
		return nil, err
	}

	_, err = conn.Cmd("PASS oauth:%s", creds.OauthToken)
	if err != nil {
		return nil, err
	}

	_, err = conn.Cmd("NICK %s", creds.Username)
	if err != nil {
		return nil, err
	}

	_, err = conn.Cmd("JOIN #%s", creds.ChannelName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ParseMessage(raw string) Message {
	parts := strings.Split(raw, " ")

	user := strings.Split(parts[1], "!")[0]
	content := strings.Join(parts[3:], " ")[1:]

	return Message{User: user, Content: content}
}
