package robification

import "testing"

func TestNewFdChat(t *testing.T) {
	newChat := NewFdChat("123456789", "test")
	if newChat.Flow_Token != "123456789" || newChat.Content != "test" {
		t.Error("Cannot create new FD Chat")
	}
}
