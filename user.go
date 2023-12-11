package goscm

import (
	"encoding/base64"
	"encoding/json"
)

type UserData struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Mail        string `json:"mail"`
	External    bool   `json:"external"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
}

// LoginUser attempts to use the user data to log in.
// This can be used to initialize a user profile that
// is authorized to but was never logged in before.
func LoginUser(baseUrl string, username string, password string) error {
	c, err := NewClient(baseUrl, "")
	if err != nil {
		return err
	}

	headers := make(map[string]string)
	token := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	headers["Authorization"] = token

	return c.getJson("/api/v2/me", nil, headers)
}

func (c *Client) CreateUser(userData UserData) error {
	headers := make(map[string]string)
	headers["Content-Type"] = mimeTypeUser

	bytes, err := json.Marshal(userData)
	if err != nil {
		return err
	}

	err = c.post("/api/v2/users", bytes, headers)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteUser(name string) error {
	return c.delete("/api/v2/users/"+name, nil)
}

func (c *Client) DeleteUserAndGroupMembership(name string) error {
	err := c.DeleteUser(name)
	if err != nil {
		return err
	}

	err = c.DeleteUserFromAllGroups(name)
	if err != nil {
		return err
	}

	return nil
}
