package goscm

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_Webhook_Parse(t *testing.T) {
	hook, err := New(Options.Secret(""))
	eventJSON, errReadFile := os.ReadFile("testdata/scm-webhook-data.json")

	if errReadFile != nil {
		t.Fatal(errReadFile.Error())
	}

	if err != nil {
		t.Fatal(err.Error())
	}

	req := httptest.NewRequest(http.MethodPost, "/api/webhook", nil)
	req.Header.Set("X-SCM-PushEvent", "Push")
	req.Body = io.NopCloser(bytes.NewReader(eventJSON))

	getreq := httptest.NewRequest(http.MethodGet, "/api/webhook", nil)
	getreq.Header.Set("X-SCM-PushEvent", "Push")

	plPush, err := hook.Parse(req, PushEvent)
	if err != nil {
		t.Fatal(err)
	}

	UNUSED(plPush)
}

func UNUSED(x ...interface{}) {}
