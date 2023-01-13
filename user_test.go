package main

import (
	"testing"
)

// Muss händisch im SCM hinzugefügt werden
func TestClient_DeleteUser(t *testing.T) {
	//Über diesen cURL Befehl kann man sich einen API Key mit * erzeugen, der dann alles darf, was auch der User darf:curl -vu Username https://next-scm.cloudogu.com/scm/api/v2/cli/login -X POST -H "Content-Type: application/json" --data '{"apiKey":"something"}'
	//
	//Das something ist der Name füë den Key und kann frei gewählt werden.
	c, err := NewClient("https://ecosystem.cloudogu.com/scm", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = c.DeleteUser("MathusanTest")
	if err != nil {
		t.Fatal(err.Error())
	}

}
