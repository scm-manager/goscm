package goscm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBranch_ListRepoBranches_Client(t *testing.T) {
	server := setupSingleTestServer(
		"testdata/branch/scm-listrepobranches-example.json",
		"/api/v2/repositories/scm-manager/scm-manager/branches/", t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	branches, err := client.ListRepoBranches("scm-manager", "scm-manager")
	require.NoError(t, err, "The ListRepoBranches method of the client threw an error.")

	assert.Len(t, branches.Embedded.Branches, 4)
	assert.Equal(t, "branchInCamelCase", branches.Embedded.Branches[0].Name)
	assert.Equal(t, "2024-11-14T10:50:45Z", branches.Embedded.Branches[1].LastCommitDate)
	assert.Equal(t, "889023f4d4c8f2672ff752cf8d3a1d8d2f0fb4af", branches.Embedded.Branches[2].Revision)
	assert.Equal(t, "scm@example.com", branches.Embedded.Branches[3].LastCommitter.Mail)
}

func TestBranch_GetRepoBranch_Client(t *testing.T) {
	server := setupSingleTestServer(
		"testdata/branch/scm-getrepobranch-example.json",
		"/api/v2/repositories/scm-manager/scm-manager/branches/feature%2Fbranch-with-a-slash", t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	branch, err := client.GetRepoBranch("scm-manager", "scm-manager", "feature/branch-with-a-slash")
	require.NoError(t, err, "The GetRepoBranch method of the client threw an error.")

	assert.Equal(t, "feature/branch-with-a-slash", branch.Name)
	assert.Equal(t, "889023f4d4c8f2672ff752cf8d3a1d8d2f0fb4af", branch.Revision)
	assert.Equal(t, false, branch.DefaultBranch)
	assert.Equal(t, "2024-11-14T14:51:42Z", branch.LastCommitDate)
	assert.Equal(t, "SCM Administrator", branch.LastCommitter.Name)
}

func TestBranch_GetDefaultBranch_Client(t *testing.T) {
	server := setupSingleTestServer(
		"testdata/branch/scm-listrepobranches-example.json",
		"/api/v2/repositories/scm-manager/scm-manager/branches/", t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	branch, err := client.GetDefaultBranch("scm-manager", "scm-manager")
	require.NoError(t, err, "The GetDefaultBranch method of the client threw an error.")

	assert.Equal(t, "main", branch.Name)
	assert.Equal(t, "22442f21ce44c78517f5100c252b1681089dd70c", branch.Revision)
	assert.Equal(t, true, branch.DefaultBranch)
	assert.Equal(t, "2024-11-14T14:24:40Z", branch.LastCommitDate)
	assert.Equal(t, "SCM Administrator", branch.LastCommitter.Name)
}
