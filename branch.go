package goscm

import (
	"errors"
	"fmt"
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
		return BranchContainer{}, fmt.Errorf("failed to load branches of %s/%s: %w", namespace, name, err)
	}
	return branchContainer, nil
}

// GetRepoBranch Get a single branch of the repository by branch name
func (c *Client) GetRepoBranch(namespace string, name string, branchName string) (Branch, error) {
	branch := Branch{}
	err := c.getJson("/api/v2/repositories/"+namespace+"/"+name+"/branches/"+url.PathEscape(branchName), &branch, nil)
	if err != nil {
		return Branch{}, fmt.Errorf("failed to load branch %s of %s/%s: %w", branchName, namespace, name, err)
	}
	return branch, nil
}

var ErrEmptyRepository = errors.New("repository is empty")
var ErrNoDefaultBranchFound = errors.New("no default branch found")

// GetDefaultBranch Get the default branch of the repository
func (c *Client) GetDefaultBranch(namespace string, name string) (Branch, error) {
	branches, err := c.ListRepoBranches(namespace, name)
	if err != nil {
		return Branch{}, err
	}
	if len(branches.Embedded.Branches) == 0 {
		return Branch{}, ErrEmptyRepository
	}
	for _, b := range branches.Embedded.Branches {
		if b.DefaultBranch {
			return b, nil
		}
	}
	return Branch{}, ErrNoDefaultBranchFound
}
