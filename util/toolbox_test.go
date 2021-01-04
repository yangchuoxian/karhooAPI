package util

import (
	"testing"
)

func TestRetrieveCredentials(t *testing.T) {
	cred, err := RetrieveCredentials()
	if err != nil {
		t.Error(err)
		return
	}
	if cred == nil {
		t.Error("credentials nil")
		return
	}
	if cred.Username == "" {
		t.Error("username is empty")
	}
	if cred.Password == "" {
		t.Error("password is empty")
	}
}
