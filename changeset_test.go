package goscm

import (
	"os"
	"testing"
)

func TestClient_Changesets(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	changesets, err := c.ListChangesets("scm-manager", "scm-manager", "develop", c.NewChangesetListFilter())
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(changesets.Embedded.Changesets) == 0 {
		t.Fatal("Could not find changesets")
	}

	if len(changesets.Embedded.Changesets) > 10 {
		t.Fatal("Filter limit doesn't work")
	}

	changeset, err := c.GetChangeset("scm-manager", "scm-manager", changesets.Embedded.Changesets[5].Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	if changeset.Description != changesets.Embedded.Changesets[5].Description {
		t.Fatal("This is the wrong changeset")
	}

	headChangeset, err := c.GetHeadChangesetForBranch("scm-manager", "scm-manager", "develop")
	if err != nil {
		t.Fatal(err.Error())
	}

	if headChangeset.Id != changesets.Embedded.Changesets[0].Id {
		t.Fatal("This is not the head?")
	}
}
