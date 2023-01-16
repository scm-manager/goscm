package main

import (
	json2 "encoding/json"
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}
	g := GroupContainer{}
	err = c.getJson("/api/v2/groups", &g, nil)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	print(g.Embedded.Groups[0].Name)
}

func TestNewClient2(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}
	g := Group{}
	err = c.getJson("/api/v2/groups/MathusanTestGruppe", &g, nil)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	print(g.Name)
}

func TestClient_Put(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}
	g := Group{}
	headers := make(map[string]string)
	headers["Content-Type"] = mimeTypeGroup
	err = c.getJson("/api/v2/groups/MathusanTestGruppe", &g, headers)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	g.LastModified = time.Now().Format("2006-01-02T15:04:05Z")
	var NewUser string
	NewUser = "testUser"
	g.Members = append(g.Members, NewUser)
	json, err := json2.Marshal(g)
	err = c.put("/api/v2/groups/MathusanTestGruppe", json, headers)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

}
