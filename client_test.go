package main

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	err := CreateUser("https://ecosystem.cloudogu.com/scm", "mkannathasan", "")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestNewClient(t *testing.T) {

}
