package unrealircd

import (
	"testing"
)

func TestServer_Rehash(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "server.rehash",
		response: map[string]interface{}{
			"success": true,
			"log":     []interface{}{},
		},
	}
	s := &Server{querier: mock}

	result, err := s.Rehash(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected rehash object, got nil")
	}
}

func TestServer_Connect(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "server.connect",
		response:       true,
	}
	s := &Server{querier: mock}

	ok, err := s.Connect("irc2.example.net")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !ok {
		t.Fatalf("Expected true, got false")
	}
}

func TestServer_Disconnect(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "server.disconnect",
		response:       true,
	}
	s := &Server{querier: mock}

	ok, err := s.Disconnect("irc2.example.net")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !ok {
		t.Fatalf("Expected true, got false")
	}
}

func TestServer_ModuleList(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "server.module_list",
		response: map[string]interface{}{
			"list": []interface{}{"mod1", "mod2"},
		},
	}
	s := &Server{querier: mock}

	result, err := s.ModuleList(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected module list, got nil")
	}
}

func TestServer_ConfigTest(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "server.config_test",
		response: map[string]interface{}{
			"success":   true,
			"exit_code": 0,
		},
	}
	s := &Server{querier: mock}

	res, err := s.ConfigTest()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("Expected config test object, got nil")
	}
}
