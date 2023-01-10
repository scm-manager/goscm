package main

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	err := CreateUser("https://ecosystem.cloudogu.com/scm", "mkannathasan", "")
	if err != nil {
		t.Fatal(err.Error())
	}
}
