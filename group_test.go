package goscm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroup_GetGroups_Client(t *testing.T) {
	server := setupSingleTestServer("testdata/group/scm-getgroups-list.json", "/api/v2/groups", t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	groups, err := client.GetGroups()
	require.NoError(t, err, "The GetGroups method of the client threw an error.")

	assert.Len(t, groups.Embedded.Groups, 3)
}

func TestGroup_GetGroup_Client(t *testing.T) {
	server := setupSingleTestServer("testdata/group/scm-getgroup-internal-init.json", "/api/v2/groups/anInternalGroup", t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	group, err := client.GetGroup("anInternalGroup")
	require.NoError(t, err, "The GetGroup method of the client threw an error.")

	assert.Equal(t, "anInternalGroup", group.Name)
}

func TestGroup_CreateGroup_Client(t *testing.T) {
	server := setupTestServer(map[string]string{
		"/api/v2/groups": "POST",
	}, true, t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	err = client.CreateGroup("newGroup", "")
	require.NoError(t, err, "The CreateGroup method of the client threw an error.")
}

func TestGroup_DeleteGroup_Client(t *testing.T) {
	server := setupTestServer(map[string]string{
		"/api/v2/groups/newGroup": "DELETE",
	}, true, t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	err = client.DeleteGroup("newGroup")
	require.NoError(t, err, "The DeleteGroup method of the client threw an error.")
}
