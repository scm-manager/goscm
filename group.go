package main

import (
	json2 "encoding/json"
	"time"
)

type userId string

type GroupContainer struct {
	Page      int `json:"page"`
	PageTotal int `json:"pageTotal"`
	Embedded  struct {
		Groups []Group `json:"groups"`
	} `json:"_embedded"`
}

type Group struct {
	Description  string   `json:"description"`
	LastModified string   `json:"lastModified,omitempty"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Members      []userId `json:"members"`
	External     bool     `json:"external"`
}

func (c *Client) GetGroups() (GroupContainer, error) {
	groupContainer := GroupContainer{}
	err := c.GetJson("/api/v2/groups", &groupContainer, nil)
	if err != nil {
		return GroupContainer{}, err
	}
	return groupContainer, nil
}

func (c *Client) GetGroup(groupID string) (Group, error) {
	group := Group{}
	err := c.GetJson("/api/v2/groups/"+groupID, &group, nil)
	if err != nil {
		return Group{}, err
	}
	return group, nil
}
func (c *Client) DeleteUserFromGroup(id userId, groupID string) error {
	group, err := c.GetGroup(groupID)
	if err != nil {
		return err
	}
	for i := 0; i < len(group.Members); i++ {
		if group.Members[i] == id {
			group.Members = remove(group.Members, i)
			group.LastModified = time.Now().Format("2006-01-02T15:04:05Z")
			var json []byte
			json, err = json2.Marshal(group)
			if err != nil {
				return err
			}
			headers := make(map[string]string)
			headers["Content-Type"] = "application/vnd.scmm-group+json;v=2"
			err = c.Put("/api/v2/groups/"+groupID, json, headers)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return nil
}
func (c *Client) AddUserToGroup(id userId, groupID string) error {
	group, err := c.GetGroup(groupID)
	if err != nil {
		return err
	}
	group.Members = append(group.Members, id)
	json, err := json2.Marshal(group)
	if err != nil {
		return err
	}
	group.LastModified = time.Now().Format("2006-01-02T15:04:05Z")
	headers := make(map[string]string)
	headers["Content-Type"] = "application/vnd.scmm-group+json;v=2"
	err = c.Put("/api/v2/groups/"+groupID, json, headers)
	if err != nil {
		return err
	}
	return nil
}

func remove(s []userId, i int) []userId {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (c *Client) DeleteUserFromAllGroups(id userId) error {
	groups, err := c.GetGroups()
	if err != nil {
		return err
	}

	for i := 0; i < len(groups.Embedded.Groups); i++ {
		err = c.DeleteUserFromGroup(id, groups.Embedded.Groups[i].Name)
		if err != nil {
			return err
		}
	}
	return nil

}

func (c *Client) CopyGroupMembershipFromOtherUser(id userId, copyId userId) error {
	groups, err := c.GetGroups()
	if err != nil {
		return err
	}
	for i := 0; i < len(groups.Embedded.Groups); i++ {
		group, err := c.GetGroup(groups.Embedded.Groups[i].Name)
		if err != nil {
			return err
		}
		for j := 0; j < len(group.Members); j++ {
			if group.Members[j] == copyId {
				err = c.AddUserToGroup(id, groups.Embedded.Groups[i].Name)
				if err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

/* Anschauungsmaterial:
func (SCM *Group) CopyWithOutMembers(group *Group) {
	group.External = SCM.External
	group.LastModified = time.Now().Format("2006-01-02T15:01:05.000Z")
	// group.LastModified = strings.Split(group.LastModified, "+")[0] // TODO: update?
	group.Type = SCM.Type
	group.Description = SCM.Description
	group.Name = SCM.Name
}

func deleteSCMGroupMembership(wg *sync.WaitGroup, group *SCMGroups, userID string, Logfile *strings.Builder) {
	var containsFlag bool
	var CopySCMGroup SCMGroups
	group.CopyWithOutMembers(&CopySCMGroup)
	for i := 0; i < len(group.Members); i++ {
		if userID == group.Members[i] {
			containsFlag = true
		} else {
			CopySCMGroup.Members = append(CopySCMGroup.Members, group.Members[i])
		}
	}
	if containsFlag {
		Logfile.WriteString(fmt.Sprintf("Removed from %s!\n", CopySCMGroup.Name))
		var jsonData, _ = json.Marshal(CopySCMGroup)
		url := "https://" + Server + "/scm/api/v2/groups/" + CopySCMGroup.Name
		err := MakeRequestEco("PUT", "application/vnd.scmm-group+json;v=2", jsonData, url, Logfile, nil, GetEcosystemPassword())
		if err != nil {
			Logfile.WriteString(err.Error())
			wg.Done()
			return
		}
	}
	if wg != nil {
		wg.Done()
	}
}


*/
