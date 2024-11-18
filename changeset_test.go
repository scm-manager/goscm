package goscm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChangeset_ListChangesets_Client(t *testing.T) {
	server := setupSingleTestServer(
		"testdata/changeset/scm-listchangesets-tenelements.json",
		"/api/v2/repositories/scm-manager/scm-manager/branches/feature%2Fbranch-with-a-slash/changesets?&pageSize=10", t)
	defer server.Close()

	c, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	changesets, err := c.ListChangesets("scm-manager", "scm-manager", "feature/branch-with-a-slash", c.NewChangesetListFilter())
	require.NoError(t, err, "The ListChangesets method of the client threw an error.")

	assert.Len(t, changesets.Embedded.Changesets, 10)
	assert.Equal(t, "Twelfth change", changesets.Embedded.Changesets[0].Description)
}

func TestChangeset_GetChangeset_Client(t *testing.T) {
	server := setupSingleTestServer("testdata/changeset/scm-getchangeset-example-commit.json", "/api/v2/repositories/scm-manager/scm-manager/changesets/903afbf1e", t)
	defer server.Close()

	c, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	changeset, err := c.GetChangeset("scm-manager", "scm-manager", "903afbf1e")
	require.NoError(t, err, "The GetChangeset method of the client threw an error.")

	assert.Equal(t, "Fourth change", changeset.Description)
}

func TestChangeset_GetHeadChangesetForBranch_Client(t *testing.T) {
	server := setupSingleTestServer(
		"testdata/changeset/scm-listchangesets-oneelement.json",
		"/api/v2/repositories/scm-manager/scm-manager/branches/feature%2Fbranch-with-a-slash/changesets?&pageSize=1", t)
	defer server.Close()

	c, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	headChangeset, err := c.GetHeadChangesetForBranch("scm-manager", "scm-manager", "feature/branch-with-a-slash")
	require.NoError(t, err, "The GetHeadChangeset method of the client threw an error.")

	assert.Equal(t, headChangeset.Id, "889023f4d4c8f2672ff752cf8d3a1d8d2f0fb4af", "Head changeset id doesn't match.")
}
