package main

import (
	"encoding/json"
	"strconv"
)

type RepoContainer struct {
	Page      int `json:"page"`
	PageTotal int `json:"pageTotal"`
	Embedded  struct {
		Repositories []Repository `json:"repositories"`
	} `json:"_embedded"`
}

type Repository struct {
	Namespace    string `json:"namespace"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Contact      string `json:"contact"`
	Archived     bool   `json:"archived"`
	LastModified string `json:"lastModified,omitempty"`
	Links        Links  `json:"_links,omitempty"`
}

type Links struct {
	ProtocolUrl []ProtocolUrl `json:"protocol,omitempty"`
}

type ProtocolUrl struct {
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

type RepoListFilter struct {
	Limit        int  `json:"limit"`
	ShowArchived bool `json:"showArchived"`
}

func (c *Client) NewRepoListFilter() *RepoListFilter {
	something := &RepoListFilter{}
	something.Limit = 10
	something.ShowArchived = true
	return something
}

// ListRepos List all repositories which the user may see
func (c *Client) ListRepos(filter *RepoListFilter) (RepoContainer, error) {
	repoContainer := RepoContainer{}
	err := c.getJson("/api/v2/repositories?pageSize="+strconv.FormatInt(int64(filter.Limit), 10)+"&archived="+strconv.FormatBool(filter.ShowArchived), &repoContainer, nil)
	if err != nil {
		return RepoContainer{}, err
	}
	return repoContainer, nil
}

// ListReposByNamespace List all repositories which the user may see by namespace
func (c *Client) ListReposByNamespace(namespace string, filter *RepoListFilter) (RepoContainer, error) {
	repoContainer := RepoContainer{}
	err := c.getJson("/api/v2/repositories/"+namespace+"?pageSize="+strconv.FormatInt(int64(filter.Limit), 10)+"&archived="+strconv.FormatBool(filter.ShowArchived), &repoContainer, nil)
	if err != nil {
		return RepoContainer{}, err
	}
	return repoContainer, nil
}

// GetRepo Get single repository by namespace and name
func (c *Client) GetRepo(namespace string, name string) (Repository, error) {
	repo := Repository{}
	err := c.getJson("/api/v2/repositories/"+namespace+"/"+name, &repo, nil)
	if err != nil {
		return Repository{}, err
	}
	return repo, nil
}

// CreateRepo Create new repository
func (c *Client) CreateRepo(repo Repository) (Repository, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = mimeTypeRepo

	bytes, err := json.Marshal(repo)
	if err != nil {
		return repo, err
	}
	err = c.post("/api/v2/repositories/", bytes, headers)
	if err != nil {
		return repo, err
	}
	return repo, nil
}

// DeleteRepo Delete repository
func (c *Client) DeleteRepo(namespace string, name string) error {
	err := c.delete("/api/v2/repositories/"+namespace+"/"+name, nil)
	if err != nil {
		return err
	}
	return nil
}
