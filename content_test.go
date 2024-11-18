package goscm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContent_GetContent_Client(t *testing.T) {
	server := setupSingleTestServer(
		"testdata/content/fourthchange",
		"/api/v2/repositories/scm-manager/scm-manager/content/b2efa637d9f5595174f1fbe6d547c7a3d6811c26/fourthchange", t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	body, _, err := client.GetContent("scm-manager", "scm-manager", "b2efa637d9f5595174f1fbe6d547c7a3d6811c26", "fourthchange")
	require.NoError(t, err, "The GetContent method of the client threw an error.")

	assert.Equal(t, "Use the forth", body)
}

func TestContent_GetContentForDefaultBranch_Client(t *testing.T) {
	server := setupTestServer(map[string]string{
		"/api/v2/repositories/scm-manager/scm-manager/content/22442f21ce44c78517f5100c252b1681089dd70c/eighthchange": "testdata/content/eighthchange",
		"/api/v2/repositories/scm-manager/scm-manager/branches/":                                                     "testdata/branch/scm-listrepobranches-example.json",
	}, false, t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	body, _, err := client.GetContentForDefaultBranch("scm-manager", "scm-manager", "eighthchange")
	require.NoError(t, err, "The GetContentForDefaultBranch method of the client threw an error.")

	assert.Equal(t, "Eighth change", body)
}
