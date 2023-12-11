package goscm

import (
	"os"
	"testing"
)

func TestClient_Content(t *testing.T) {

	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	content, r, err := c.GetContent("scm-manager", "scm-manager", "develop", "Jenkinsfile")
	if err != nil {
		t.Fatal(err.Error())
	}

	if content == "" {
		t.Fatal("Could not find content")
	}

	content, r, err = c.GetContent("scm-manager", "scm-manager", "develop", "docker")
	if err != nil {
		t.Fatal(err.Error())
	}

	if content != "" && r.StatusCode == 404 {
		t.Fatal("Expecting a 404 error since it is a directory")
	}

	content, r, err = c.GetContentForDefaultBranch("scm-manager", "scm-manager", "build.gradle")
	if err != nil {
		t.Fatal(err.Error())
	}

	if content == "" {
		t.Fatal("Could not find content for default branch")
	}
}
