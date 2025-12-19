package unrealircd

import (
	"testing"
)

func TestMessagePrivmsg(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.privmsg",
		response:       true,
	}
	msg := &Message{querier: mock}
	_, err := msg.Privmsg("testnick", "test message")
	if err != nil {
		t.Errorf("Privmsg failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}

func TestMessagePrivmsgMultipleTargets(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.privmsg",
		response:       true,
	}
	msg := &Message{querier: mock}
	targets := []string{"nick1", "#chan"}
	_, err := msg.Privmsg(targets, "broadcast message")
	if err != nil {
		t.Errorf("Privmsg (multiple) failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}

func TestMessageNotice(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.notice",
		response:       true,
	}
	msg := &Message{querier: mock}
	_, err := msg.Notice("testnick", "test notice")
	if err != nil {
		t.Errorf("Notice failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}

func TestMessageNoticeMultipleTargets(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.notice",
		response:       true,
	}
	msg := &Message{querier: mock}
	targets := []string{"nick1", "#chan"}
	_, err := msg.Notice(targets, "test notice to many")
	if err != nil {
		t.Errorf("Notice (multiple) failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}

func TestMessageNumeric(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.numeric",
		response:       true,
	}
	msg := &Message{querier: mock}
	_, err := msg.Numeric("testnick", 123, "test numeric message")
	if err != nil {
		t.Errorf("Numeric failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}

func TestMessageStandardReply(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.standardreply",
		response:       true,
	}
	msg := &Message{querier: mock}
	_, err := msg.StandardReply("testnick", "NOTE", "TEST", "test description", nil)
	if err != nil {
		t.Errorf("StandardReply failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}

func TestMessageWallops(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.send_wallops",
		response:       true,
	}
	msg := &Message{querier: mock}
	_, err := msg.Wallops("test wallops")
	if err != nil {
		t.Errorf("Wallops failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}

func TestMessageGlobops(t *testing.T) {
	mock := &mockQuerier{
		expectedMethod: "message.send_globops",
		response:       true,
	}
	msg := &Message{querier: mock}
	_, err := msg.Globops("test globops")
	if err != nil {
		t.Errorf("Globops failed: %v", err)
	}
	if !mock.called {
		t.Error("Query was not called")
	}
}
