package conn

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"jake/method/dev/internal/conn/credentials"
	"net"
	"time"
)

const TCP string = "tcp"

type Conn struct {
	Conn net.Conn
	R    *bufio.Reader
	W    *bufio.Writer
}

// Rate limit to 7500 messages per 30 seconds per Twitch API
var rateLimit = time.NewTicker(time.Second / 250)

// SendMessage sends a message to the specified channel
func (c *Conn) SendMessage(channel, message string) error {
	<-rateLimit.C
	_, err := c.W.WriteString(fmt.Sprintf("PRIVMSG %s :%s\r\n", channel, message))
	if err != nil {
		return err
	}
	err = c.W.Flush()
	return err
}

// Cmd sends a command to the IRC server
func (c *Conn) Cmd(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(c.W, format+"\r\n", args...)
}

// Connect connects to the Twitch IRC server and performs the initial handshake
func Connect(creds credentials.Credentials, addr string) (*Conn, error) {
	conn, err := tls.Dial(TCP, addr, &tls.Config{
		InsecureSkipVerify: true, // Consider setting this to false in production for greater security
	})
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
