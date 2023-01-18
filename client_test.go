package goscm

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = c.GetIndex()
	if err != nil {
		t.Fatal(err.Error())
		return
	}
}
