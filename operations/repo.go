package operations

import (
	"github.com/go-git/go-git/v5/plumbing"
	"log"

	"github.com/go-git/go-git/v5"
)

func GetRepo() *git.Repository {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal("Could not open git repo (did you run 'git init'?): ", err)
	}

	return repo
}

func EnsureBranch(repo *git.Repository, branchName string) error {
	headRef, _ := repo.Head()

	// Create the full reference name, e.g., refs/heads/branch-a
	refName := plumbing.ReferenceName("refs/heads/" + branchName)

	// Check if it exists
	_, err := repo.Reference(refName, true)
	if err != nil {
		// If not, create it pointing to the same hash as HEAD
		ref := plumbing.NewHashReference(refName, headRef.Hash())
		repo.Storer.SetReference(ref)
		log.Printf("Created branch: %s", branchName)
		return nil
	}

	return err
}

func pushToOrigin(r *git.Repository, ref *plumbing.Reference) error {
	return nil
}
