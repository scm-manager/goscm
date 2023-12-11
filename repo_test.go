package goscm

import (
	"os"
	"testing"
)

func TestClient_Repository(t *testing.T) {
	c, err := NewClient("https://stagex.cloudogu.com/scm", os.Getenv("SCM_BEARER_TOKEN"))
	if err != nil {
		t.Fatal(err.Error())
	}

	// Cleanup before tests => Delete must be idempotent
	err = c.DeleteRepo("hitchhiker", "HeartOfGold")
	if err != nil {
		t.Fatal(err.Error())
	}

	if err != nil {
		t.Fatal(err.Error())
	}

	// Create new repository
	repoData := Repository{
		Namespace:   "hitchhiker",
		Name:        "HeartOfGold",
		Description: "Hitchhiker's Guide through the galaxy",
		Contact:     "trillian@scm-manager.org",
		Type:        "git",
		Archived:    false,
	}

	newRepo, err := c.CreateRepo(repoData)
	if err != nil {
		t.Fatal(err.Error())
	}

	if newRepo.Type != repoData.Type {
		t.Fatal("Wrong repository type")
	}

	// Get single repo
	repo, err := c.GetRepo("hitchhiker", "HeartOfGold")
	if err != nil {
		t.Fatal(err.Error())
	}

	if repo.Type == "" || repo.Namespace == "" || repo.Name == "" {
		t.Fatalf("Could not get repo")
	}

	if len(repo.Links.ProtocolUrl) == 0 {
		t.Fatalf("The repository is missing protocol urls")
	}

	// List repos
	repos, err := c.ListRepos(c.NewRepoListFilter())
	if err != nil {
		t.Fatal(err.Error())
	}
	// There should be at least one repository which the user just created
	if len(repos.Embedded.Repositories) < 1 {
		t.Fatal("Could not list repositories")
	}

	err = c.DeleteRepo("hitchhiker", "HeartOfGold")
	if err != nil {
		t.Fatal(err.Error())
	}
}
