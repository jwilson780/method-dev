package bot

import (
	"jake/method/dev/src/bot/credentials"
	"net"
	"testing"
)

func TestBot(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0") // choose a random open port
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

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Errorf("Read error: %v", err)
			return
		}

		msg := string(buf[:n])
		expected := "PRIVMSG #testchannel :test message\r\n"
		if msg != expected {
			t.Errorf("Expected to receive %q, got %q", expected, msg)
		}
	}()

	c2 := credentials.Credentials{
		Username:    "t",
		OauthToken:  "t",
		ChannelName: "t",
	}

	c, err := Connect(c2, l.Addr().String())
	if err != nil {
		t.Fatalf("Connect error: %v", err)
	}

	msg := "test message"
	err = c.SendMessage("#testchannel", msg)
	if err != nil {
		t.Fatalf("send message error: %v", err)
	}
}
