package goscm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepo_CreateRepo(t *testing.T) {
	server := setupTestServer(map[string]string{
		"/api/v2/repositories/": "POST",
	}, true, t)
	defer server.Close()

	client, err := NewClient(server.URL, "")
	require.NoError(t, err, "Error during client initialization")

	_, err = client.CreateRepo(Repository{Name: "exampleRepo", Namespace: "exampleNamespace"})
	require.NoError(t, err, "The CreateRepo method of the client threw an error.")
}
