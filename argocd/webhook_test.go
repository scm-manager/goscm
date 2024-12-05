package argocd

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
	req.Header.Set("X-SCM-Event", "Push")
	req.Body = io.NopCloser(bytes.NewReader(eventJSON))

	getreq := httptest.NewRequest(http.MethodGet, "/api/webhook", nil)
	getreq.Header.Set("X-SCM-Event", "Push")

	plPush, err := hook.Parse(req, PushEvent)
	require.NoError(t, err, "Error while parsing hook information in Goscm.")

	assert.Equal(t, "https://scm-manager.org/scm/scm-manager/argocd-test", plPush.(PushEventPayload).Repository.SourceUrl)
	assert.Equal(t, "develop", plPush.(PushEventPayload).Branch.Name)
	assert.True(t, plPush.(PushEventPayload).Branch.DefaultBranch)
}

func TestWebhook_VerifyPayload_AssertCorrectMAC(t *testing.T) {
	hook, err := New(Options.Secret("verySecretKey"))
	require.NoError(t, err, "Error while creating new Goscm instance for webhook parse test.")

	signature := "sha1=6149c0ab59e04ac11ac3d1e0e44ae0a96f67d0a4"
	payload := []byte("notSoSecretPayload")

	err = hook.VerifyPayload(signature, payload)

	require.NoError(t, err, "Error while executing VerifyPayload method for webhook test.")
}

func TestWebhook_VerifyPayload_DenyWrongMAC(t *testing.T) {
	hook, err := New(Options.Secret("youShallNotPass"))
	require.NoError(t, err, "Error while creating new Goscm instance for webhook parse test.")

	signature := "sha1=6149c0ab59e04ac11ac3d1e0e44ae0a96f67d0a4"
	payload := []byte("notSoSecretPayload")

	err = hook.VerifyPayload(signature, payload)

	require.Error(t, err, "[FATAL SECURITY ISSUE] Payload verification for webhook did not deny a wrong MAC!!")
}
