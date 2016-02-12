package onesignal

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	key := "fake key"
	httpClient := &http.Client{}

	c := NewClient(key, httpClient)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, should be %v", got, want)
	}

	if got, want := c.Key, key; got != want {
		t.Errorf("NewClient Key is %v, should be %v", got, want)
	}

	if got, want := c.Client, httpClient; got != want {
		t.Errorf("NewClient Client is %v, should be %v", got, want)
	}
}
