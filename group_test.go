package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"testing"
	"time"
)

func Test_Group(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	groupName := "SOS-CI-Test-Group"

	// Create Group
	err = c.CreateGroup(groupName, "Test Group created by go-scm Tests.")
	if err != nil {
		t.Fatal(err.Error())
	}

	prime, _ := rand.Prime(rand.Reader, 64)
	password := base64.StdEncoding.EncodeToString([]byte(time.Now().String() + prime.String()))

	userData := UserData{
		Name:        "SOS-CI-Test-Groupuser",
		DisplayName: "SOS CI Test-Groupuser",
		Mail:        "",
		External:    false,
		Password:    password,
		Active:      true,
	}

	// Create User
	err = c.CreateUser(userData)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Add User to Group
	err = c.AddUserToGroup(userData.Name, groupName)
	if err != nil {
		t.Fatal(err.Error())
	}
	group, err := c.GetGroup(groupName)
	contains := false
	for i := 0; i < len(group.Members); i++ {
		if group.Members[i] == userData.Name {
			contains = true
			break
		}
	}
	if !contains {
		t.Fatalf("user %q not present in group %q", userData.Name, groupName)
	}

	// Delete User From All Groups
	err = c.DeleteUserFromAllGroups(userData.Name)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Delete User
	err = c.DeleteUser(userData.Name)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Delete Group
	err = c.DeleteGroup(groupName)
	if err != nil {
		t.Fatal(err.Error())
	}
}
