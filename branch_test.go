package goscm

import (
	"os"
	"testing"
)

func TestClient_Branches(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	branches, err := c.ListRepoBranches("scm-manager", "scm-manager")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(branches.Embedded.Branches) < 1 {
		t.Fatal("Could not find branches")
	}

	branch, err := c.GetRepoBranch("scm-manager", "scm-manager", branches.Embedded.Branches[0].Name)
	if err != nil {
		t.Fatal(err.Error())
	}

	if branch != branches.Embedded.Branches[0] {
		t.Fatal("Got the wrong branch?")
	}

	defaultBranch, err := c.GetDefaultBranch("scm-manager", "scm-manager")
	if err != nil {
		t.Fatal(err.Error())
	}

	if !defaultBranch.DefaultBranch {
		t.Fatal("This is not the default branch!")
	}
}
