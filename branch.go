package main

import (
	"errors"
	"net/url"
)

type BranchContainer struct {
	Embedded struct {
		Branches []Branch `json:"branches"`
	} `json:"_embedded"`
}

type Branch struct {
	Name           string        `json:"name"`
	DefaultBranch  bool          `json:"defaultBranch"`
	Revision       string        `json:"revision"`
	Stale          bool          `json:"stale"`
	LastCommitDate string        `json:"lastCommitDate,omitempty"`
	LastCommitter  LastCommitter `json:"lastCommitter"`
}

type LastCommitter struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

// ListRepoBranches List all branches of the repository
func (c *Client) ListRepoBranches(namespace string, name string) (BranchContainer, error) {
	branchContainer := BranchContainer{}
	err := c.getJson("/api/v2/repositories/"+namespace+"/"+name+"/branches/", &branchContainer, nil)
	if err != nil {
		return BranchContainer{}, err
	}
	return branchContainer, nil
}

// GetRepoBranch Get a single branch of the repository by branch name
func (c *Client) GetRepoBranch(namespace string, name string, branchName string) (Branch, error) {
	branch := Branch{}
	err := c.getJson("/api/v2/repositories/"+namespace+"/"+name+"/branches/"+url.PathEscape(branchName), &branch, nil)
	if err != nil {
		return Branch{}, err
	}
	return branch, nil
}

// GetDefaultBranch Get the default branch of the repository
func (c *Client) GetDefaultBranch(namespace string, name string) (Branch, error) {
	branches, err := c.ListRepoBranches(namespace, name)
	if err != nil {
		return Branch{}, err
	}
	for _, b := range branches.Embedded.Branches {
		if b.DefaultBranch {
			return b, nil
		}
	}
	return Branch{}, errors.New("No default branch found")
}
