package goscm

import (
	"strconv"
)

type PullRequestContainer struct {
	Page      int `json:"page"`
	PageTotal int `json:"pageTotal"`
	Embedded  struct {
		PullRequests []PullRequest `json:"pullRequests"`
	} `json:"_embedded"`
}

type PullRequest struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Source       string   `json:"source"`
	Target       string   `json:"target"`
	Status       string   `json:"status"`
	CreationDate string   `json:"creationDate,omitempty"`
	CloseDate    string   `json:"closeDate,omitempty"`
	Labels       []string `json:"labels"`
}

type PullRequestListFilter struct {
	Status string `json:"status"`
	Limit  int    `json:"limit"`
}

func (c *Client) NewPullRequestListFilter() *PullRequestListFilter {
	filter := PullRequestListFilter{}
	filter.Status = "OPEN"
	filter.Limit = 10
	return &filter
}

// ListPullRequests List all pull requests for repository
func (c *Client) ListPullRequests(namespace string, name string, filter *PullRequestListFilter) (PullRequestContainer, error) {
	pullRequestContainer := PullRequestContainer{}
	err := c.getJson("/api/v2/pull-requests/"+namespace+"/"+name+"?status="+filter.Status+"&pageSize="+strconv.FormatInt(int64(filter.Limit), 10), &pullRequestContainer, nil)
	if err != nil {
		return PullRequestContainer{}, err
	}
	return pullRequestContainer, nil
}

// GetPullRequest Get single pull request for repository
func (c *Client) GetPullRequest(namespace string, name string, id string) (PullRequest, error) {
	pullRequest := PullRequest{}
	err := c.getJson("/api/v2/pull-requests/"+namespace+"/"+name+"/"+id, &pullRequest, nil)
	if err != nil {
		return PullRequest{}, err
	}
	return pullRequest, err
}
