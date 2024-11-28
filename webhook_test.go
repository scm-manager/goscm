package goscm

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWebhook_Parse(t *testing.T) {
	hook, err := New(Options.Secret(""))
	eventJSON, errReadFile := os.ReadFile("testdata/webhook/scm-webhook-data.json")

	require.NoError(t, err, "Error while creating new Goscm instance for webhook parse test.")
	require.NoError(t, errReadFile, "Error thrown while loading JSON for webhook parse test.")

	req := httptest.NewRequest(http.MethodPost, "/api/webhook", nil)
	req.Header.Set("X-SCM-PushEvent", "Push")
	req.Body = io.NopCloser(bytes.NewReader(eventJSON))

	getreq := httptest.NewRequest(http.MethodGet, "/api/webhook", nil)
	getreq.Header.Set("X-SCM-PushEvent", "Push")

	plPush, err := hook.Parse(req, PushEvent)
	require.NoError(t, err, "Error while parsing hook information in Goscm.")

	assert.Equal(t, "https://scm-manager.org/scm/scm-manager/argocd-test", plPush.(PushEventPayload).SourceUrl)
	assert.Equal(t, "develop", plPush.(PushEventPayload).Branch.Name)
	assert.True(t, plPush.(PushEventPayload).Branch.DefaultBranch)
}
