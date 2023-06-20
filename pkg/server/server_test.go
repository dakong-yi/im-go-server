package server

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	// Testing valid IP, port, and message buffer size
	s1 := NewServer("127.0.0.1", 8080, 10)
	if s1.Ip != "127.0.0.1" || s1.Port != 8080 || len(s1.Message) != 0 {
		t.Errorf("NewServer() failed with valid IP, port, and message buffer size.")
	}

	// Testing invalid IP
	s2 := NewServer("invalid", 8080, 10)
	if s2 != nil {
		t.Errorf("NewServer() failed with invalid IP.")
	}

	// Testing invalid port
	s3 := NewServer("127.0.0.1", -1, 10)
	if s3 != nil {
		t.Errorf("NewServer() failed with invalid port.")
	}

	// Testing invalid message buffer size
	s4 := NewServer("127.0.0.1", 8080, -1)
	if s4 != nil {
		t.Errorf("NewServer() failed with invalid message buffer size.")
	}
}
