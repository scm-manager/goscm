package main

import (
	"testing"
)

const testuser1 userId = "testuser1"

// Muss händisch im SCM hinzugefügt werden
const testGroup = "MathusanTestGruppe"

func TestClient_AddUserToGroup(t *testing.T) {
	c, err := NewClient("https://ecosystem.cloudogu.com/scm", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = c.AddUserToGroup("testuser1", "MathusanTestGruppe")
	if err != nil {
		t.Fatal(err.Error())
	}
	g, err := c.GetGroup("MathusanTestGruppe")
	for i := 0; i < len(g.Members); i++ {
		if g.Members[i] == "testuser1" {
			return
		}
	}
	t.Fatal("User has not been added to group")

}

func TestClient_DeleteUserFromAllGroups(t *testing.T) {
	c, err := NewClient("https://ecosystem.cloudogu.com/scm", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = c.DeleteUserFromAllGroups(testuser1)
	if err != nil {
		t.Fatal(err.Error())
	}

}

func TestClient_DeleteUserFromGroup(t *testing.T) {
	c, err := NewClient("https://ecosystem.cloudogu.com/scm", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = c.DeleteUserFromGroup(testuser1, testGroup)
	if err != nil {
		t.Fatal(err.Error())
	}
	g, err := c.GetGroup(testGroup)
	for i := 0; i < len(g.Members); i++ {
		if g.Members[i] == testuser1 {
			t.Fatal("User has not been deleted")
		}
	}
}
