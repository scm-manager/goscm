package main

import "encoding/base64"

// LoginUser attempts to use the user data to log in.
// This can be used to initialize a user profile that
// is authorized to but never was logged in before.
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

func (c *Client) DeleteUser(id string) error {
	headers := make(map[string]string)
	headers["Content-Type"] = mimeTypeUser
	err := c.delete("/api/v2/users/"+id, headers)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteUserAndGroupMembership(id userId) error {
	err := c.DeleteUser(string(id))
	if err != nil {
		return err
	}
	err = c.DeleteUserFromAllGroups(id)
	if err != nil {
		return err
	}
	return nil
}
