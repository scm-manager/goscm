package goscm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser_CreateUser_Client(t *testing.T) {
	server := setupTestServer(map[string]string{
		"/api/v2/users": "POST",
	}, true, t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	err = client.CreateUser(UserData{})
	require.NoError(t, err, "The CreateGroup method of the client threw an error.")
}

func TestUser_DeleteUser_Client(t *testing.T) {
	server := setupTestServer(map[string]string{
		"/api/v2/users/exampleUser": "DELETE",
	}, true, t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	err = client.DeleteUser("exampleUser")
	require.NoError(t, err, "The CreateGroup method of the client threw an error.")
}
