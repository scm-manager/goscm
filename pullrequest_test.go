package goscm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPullRequest_ListPullRequests_Client(t *testing.T) {
	server := setupSingleTestServer("testdata/pullrequest/scm-listpullrequests-three-results.json", "/api/v2/pull-requests/scm-manager/scm-manager?status=ALL&pageSize=3", t)
	defer server.Close()

	c, err := NewClient(server.URL, "")
	c.SetHttpClient(server.Client())

	prs, err := c.ListPullRequests("scm-manager", "scm-manager", &PullRequestListFilter{Limit: 3, Status: "ALL"})
	require.NoError(t, err, "The ListPullRequests method of the client threw an error.")

	assert.Len(t, prs.Embedded.PullRequests, 3, "Amount of pull request objects in response doesn't match expectation.")
}

func TestPullRequest_GetPullRequest_Client(t *testing.T) {
	server := setupSingleTestServer("testdata/pullrequest/scm-getpullrequest-second-result.json", "/api/v2/pull-requests/scm-manager/scm-manager/2", t)
	defer server.Close()

	c, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	pr, err := c.GetPullRequest("scm-manager", "scm-manager", "2")
	require.NoError(t, err, "The GetPullRequest method of the client threw an error.")

	assert.Equal(t, "2", pr.Id, "The Id of the received pull request doesn't match the expected one.")
	assert.Equal(t, "feature/branch-with-a-slash", pr.Source, "The Source of the received pull request doesn't match the expected one.")
	assert.Equal(t, "main", pr.Target, "The Target of the received pull request doesn't match the expected one.")
	assert.Equal(t, "heartOfGold", pr.Labels[0], "The Target of the received pull request doesn't match the expected one.")
}
