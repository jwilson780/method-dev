package bot

import (
	"bufio"
	"fmt"
	"jake/method/dev/src/bot/credentials"
	"net"
	"strings"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	// Start a local server to simulate a Twitch IRC server
	l, err := net.Listen(TCP, "localhost:0") // choose a random open port
	if err != nil {
		t.Fatalf("net.Listen error: %v", err)
	}

	go func() {
		conn, err := l.Accept()
		if err != nil {
			t.Errorf("Accept error: %v", err)
			return
		}
		defer conn.Close()

		// Read incoming messages and echo them back to the client
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				t.Errorf("Read error: %v", err)
				return
			}

			msg := string(buf[:n])
			conn.Write([]byte(msg + "\r\n")) // echo back the message
		}
	}()

	// Test Connect
	c2 := credentials.Credentials{
		Username:    "test",
		OauthToken:  "test",
		ChannelName: "test",
	}

	// Connect to the local server
	c, err := Connect(c2, l.Addr().String()) // will send handshake message
	if err != nil {
		t.Fatalf("Connect error: %v", err)
	}

	// Listen for message from the bot
	reader := bufio.NewReader(c.Conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("ReadString error: %v", err)
	}

	// Check that the bot sent the correct message, which is just the first connection message
	expected := fmt.Sprintf("PASS test\r\n")
	if !strings.Contains(line, expected) {
		t.Errorf("Expected to receive %q, got %q", expected, line)
	}

	// avoid test to run indefinitely
	time.Sleep(250 * time.Millisecond)
}
