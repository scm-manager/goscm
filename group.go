package main

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
