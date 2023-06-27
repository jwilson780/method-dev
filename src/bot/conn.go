package bot

import (
	"bufio"
	"fmt"
	"jake/method/dev/src/bot/credentials"
	"net"
)

const TCP string = "tcp"

type Conn struct {
	Conn net.Conn
	R    *bufio.Reader
	W    *bufio.Writer
}

func (c *Conn) Cmd(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(c.W, format+"\r\n", args...)
}

// SendMessage sends a message to the specified channel
func (c *Conn) SendMessage(channel, message string) error {
	_, err := c.W.WriteString(fmt.Sprintf("PRIVMSG %s :%s\r\n", channel, message))
	if err != nil {
		return err
	}
	err = c.W.Flush()
	return err
}

// Connect connects to the Twitch IRC server and performs the initial handshake
func Connect(creds credentials.Credentials, addr string) (*Conn, error) {
	conn, err := net.Dial(TCP, addr)
	if err != nil {
		return nil, err
	}

	c := &Conn{
		Conn: conn,
		R:    bufio.NewReader(conn),
		W:    bufio.NewWriter(conn),
	}

	if _, err := c.Cmd("PASS %s", creds.OauthToken); err != nil {
		return nil, err
	}

	if _, err := c.Cmd("NICK %s", creds.Username); err != nil {
		return nil, err
	}

	if _, err := c.Cmd("JOIN #%s", creds.ChannelName); err != nil {
		return nil, err
	}

	if err := c.W.Flush(); err != nil {
		return nil, err
	}

	return c, nil
}
