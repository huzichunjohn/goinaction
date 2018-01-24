package data

import (
	"testing"
)

func Test_CreateThread(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread.")
	}
	if conv.UserId != users[0].Id {
		t.Error("User not linked with thread")
	}
}
