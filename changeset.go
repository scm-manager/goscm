package goscm

import "strconv"

type ChangesetContainer struct {
	Page      int `json:"page"`
	PageTotal int `json:"pageTotal"`
	Embedded  struct {
		Changesets []Changeset `json:"changesets"`
	} `json:"_embedded"`
}

type Changeset struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Author      Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type ChangesetListFilter struct {
	Limit int `json:"limit"`
}

func (c *Client) NewChangesetListFilter() *ChangesetListFilter {
	filter := ChangesetListFilter{}
	filter.Limit = 10
	return &filter
}

// ListChangesets List all changesets for repository by branch
func (c *Client) ListChangesets(namespace string, name string, branch string, filter *ChangesetListFilter) (ChangesetContainer, error) {
	changesetContainer := ChangesetContainer{}
	err := c.getJson("/api/v2/repositories/"+namespace+"/"+name+"/branches/"+branch+"/changesets?&pageSize="+strconv.FormatInt(int64(filter.Limit), 10), &changesetContainer, nil)
	if err != nil {
		return ChangesetContainer{}, err
	}
	return changesetContainer, nil
}

// GetChangeset Get a single repository changeset by id
func (c *Client) GetChangeset(namespace string, name string, id string) (Changeset, error) {
	changeset := Changeset{}
	err := c.getJson("/api/v2/repositories/"+namespace+"/"+name+"/changesets/"+id, &changeset, nil)
	if err != nil {
		return Changeset{}, err
	}
	return changeset, err
}

// GetHeadChangesetForBranch Get head changeset for repository by branch
func (c *Client) GetHeadChangesetForBranch(namespace string, name string, branch string) (Changeset, error) {
	filter := c.NewChangesetListFilter()
	filter.Limit = 1
	changesets, err := c.ListChangesets(namespace, name, branch, filter)
	if err != nil {
		return Changeset{}, err
	}
	return changesets.Embedded.Changesets[0], nil
}
