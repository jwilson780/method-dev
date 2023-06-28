package conn

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"jake/method/dev/internal/conn/credentials"
)

// would most likely skip this test in production or in CI
func TestConnect(t *testing.T) {
	// Note: you will need to gen the cert.pem and key.pem files yourself in order to run this test
	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem") // Provide your own self-signed certificate and key
	if err != nil {
		t.Fatal(err) // make sure key doesn't have a passkey
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen(TCP, "127.0.0.1:0", config)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				t.Error(err)
				return
			}
			go handleConnection(conn, t)
		}
	}()

	creds := credentials.Credentials{
		Username:    "test",
		OauthToken:  "test",
		ChannelName: "test",
	}

	c, err := Connect(creds, ln.Addr().String())
	if err != nil {
		t.Fatalf("Connect error: %v", err)
	}
	defer c.Conn.Close()

	reader := bufio.NewReader(c.Conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("ReadString error: %v", err)
	}

	expected := fmt.Sprintf("PASS %s\r\n", creds.OauthToken)
	if !strings.Contains(line, expected) {
		t.Errorf("Expected to receive %q, got %q", expected, line)
	}

	// avoid test to run indefinitely
	time.Sleep(250 * time.Millisecond)
}

func handleConnection(conn net.Conn, t *testing.T) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			t.Errorf("Read error: %v", err)
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			t.Errorf("Write error: %v", err)
			return
		}
	}
}
