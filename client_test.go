package main

import (
	json2 "encoding/json"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	err := CreateUser("https://ecosystem.cloudogu.com/scm", "mkannathasan", "")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestNewClient(t *testing.T) {
	c, err := NewClient("https://ecosystem.cloudogu.com/scm", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	g := GroupContainer{}
	err = c.GetJson("/api/v2/groups", &g, nil)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	print(g.Embedded.Groups[0].Name)
}

func TestNewClient2(t *testing.T) {
	c, err := NewClient("https://ecosystem.cloudogu.com/scm", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	g := Group{}
	err = c.GetJson("/api/v2/groups/MathusanTestGruppe", &g, nil)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	print(g.Name)
}

func TestClient_Put(t *testing.T) {
	c, err := NewClient("https://ecosystem.cloudogu.com/scm", "")
	if err != nil {
		t.Fatal(err.Error())
	}
	g := Group{}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/vnd.scmm-group+json;v=2"
	err = c.GetJson("/api/v2/groups/MathusanTestGruppe", &g, headers)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	g.LastModified = time.Now().Format("2006-01-02T15:04:05Z")
	var NewUser userId
	NewUser = "testUser"
	g.Members = append(g.Members, NewUser)
	json, err := json2.Marshal(g)
	err = c.Put("/api/v2/groups/MathusanTestGruppe", json, headers)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

}
