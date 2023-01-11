package main

// CreateUser TO-DO: guten namen finden
// Funktion wird benutzt, um einen authorisation aber noch nicht im SCM
// erstellten User zu erstellen.
func CreateUser(baseUrl string, username string, password string) error {
	c, err := NewClient(baseUrl, "")
	if err != nil {
		return err
	}
	headers := make(map[string]string)
	headers["Authorization"] = basicAuth(username, password)
	err = c.GetJson("/api/v2/me", nil, headers)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteUser(id string) error {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/vnd.scmm-user+json;v=2"
	err := c.Delete("/api/v2/users/"+id, headers)
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
