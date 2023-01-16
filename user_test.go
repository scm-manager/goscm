package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"testing"
	"time"
)

func Test_User(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	prime, _ := rand.Prime(rand.Reader, 64)
	password := base64.StdEncoding.EncodeToString([]byte(time.Now().String() + prime.String()))

	userData := UserData{
		Name:        "SOS-CI-Test-User",
		DisplayName: "SOS CI Test-User",
		Mail:        "",
		External:    false,
		Password:    password,
		Active:      true,
	}

	t.Log("Create User ...")
	err = c.CreateUser(userData)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log("Login User ...")
	err = LoginUser("https://stagex.cloudogu.com/scm", userData.Name, userData.Password)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log("Delete User ...")
	err = c.DeleteUser("Test-User-SOS")
	if err != nil {
		t.Fatal(err.Error())
	}
}
