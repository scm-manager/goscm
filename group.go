package goscm

import (
	"encoding/json"
	"time"
)

type GroupContainer struct {
	Page      int `json:"page"`
	PageTotal int `json:"pageTotal"`
	Embedded  struct {
		Groups []Group `json:"groups"`
	} `json:"_embedded"`
}

type Group struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	LastModified string   `json:"lastModified,omitempty"`
	Type         string   `json:"type"`
	Members      []string `json:"members"`
	External     bool     `json:"external"`
}

// GetGroups returns all existing Groups
func (c *Client) GetGroups() (GroupContainer, error) {
	groupContainer := GroupContainer{}
	err := c.getJson("/api/v2/groups", &groupContainer, nil)
	if err != nil {
		return GroupContainer{}, err
	}
	return groupContainer, nil
}

// GetGroup returns the Group with name groupName
func (c *Client) GetGroup(groupName string) (Group, error) {
	group := Group{}
	err := c.getJson("/api/v2/groups/"+groupName, &group, nil)
	if err != nil {
		return Group{}, err
	}
	return group, nil
}

func (c *Client) CreateGroup(name string, description string) error {
	headers := make(map[string]string)
	headers["Content-Type"] = mimeTypeGroup

	groupData := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{
		Name:        name,
		Description: description,
	}

	bytes, err := json.Marshal(groupData)
	if err != nil {
		return err
	}

	return c.post("/api/v2/groups", bytes, headers)
}

func (c *Client) DeleteGroup(groupName string) error {
	return c.delete("/api/v2/groups/"+groupName, nil)
}

func (c *Client) DeleteUserFromGroup(userName string, groupName string) error {
	group, err := c.GetGroup(groupName)
	if err != nil {
		return err
	}

	for i := 0; i < len(group.Members); i++ {
		if group.Members[i] == userName {
			group.Members = remove(group.Members, i)
			group.LastModified = time.Now().Format("2006-01-02T15:04:05Z")

			bytes, err := json.Marshal(group)
			if err != nil {
				return err
			}

			headers := make(map[string]string)
			headers["Content-Type"] = mimeTypeGroup

			err = c.put("/api/v2/groups/"+groupName, bytes, headers)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return nil
}
func (c *Client) AddUserToGroup(userName string, groupName string) error {
	group, err := c.GetGroup(groupName)
	if err != nil {
		return err
	}
	group.Members = append(group.Members, userName)
	bytes, err := json.Marshal(group)
	if err != nil {
		return err
	}
	group.LastModified = time.Now().Format("2006-01-02T15:04:05Z")
	headers := make(map[string]string)
	headers["Content-Type"] = mimeTypeGroup
	err = c.put("/api/v2/groups/"+groupName, bytes, headers)
	if err != nil {
		return err
	}
	return nil
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (c *Client) DeleteUserFromAllGroups(userName string) error {
	groups, err := c.GetGroups()
	if err != nil {
		return err
	}

	for _, group := range groups.Embedded.Groups {
		err = c.DeleteUserFromGroup(userName, group.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) CopyGroupMembershipsFromOtherUser(userName string, templateUserName string) error {
	groups, err := c.GetGroups()
	if err != nil {
		return err
	}

	for _, group := range groups.Embedded.Groups {
		for _, member := range group.Members {
			if member == templateUserName {
				err = c.AddUserToGroup(userName, group.Name)
				if err != nil {
					return err
				}
				break
			}
		}
	}

	return nil
}
