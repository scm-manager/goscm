package main

import (
	"os"
	"testing"
)

func TestClient_PullRequests(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	prs, err := c.ListPullRequests("scm-manager", "scm-manager", &PullRequestListFilter{Limit: 3, Status: "ALL"})
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(prs.Embedded.PullRequests) == 0 {
		t.Fatal("Could not find pull requests")
	}

	if len(prs.Embedded.PullRequests) > 3 {
		t.Fatal("Filter limit doesn't work")
	}

	pr, err := c.GetPullRequest("scm-manager", "scm-manager", prs.Embedded.PullRequests[0].Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	if pr.Source == "" || pr.Title == "" {
		t.Fatal("Could not find pull request")
	}
}
