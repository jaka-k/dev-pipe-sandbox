package operations

import (
	"github.com/go-git/go-git/v5"
	"io"
	"net/http"
	"strings"

	"github.com/go-git/go-git/v5/plumbing"
)

func HandleBranch(w http.ResponseWriter, r *http.Request, repo *git.Repository) {
	branchName := strings.TrimPrefix(r.URL.Path, "/branch/")
	if branchName == "" {
		http.Error(w, "Branch name is required", http.StatusBadRequest)
		return
	}

	// Construct the reference name for the branch
	refName := plumbing.ReferenceName("refs/heads/" + branchName)

	// Resolve the revision
	ref, err := repo.Reference(refName, true)
	if err != nil {
		http.Error(w, "Branch not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Get the commit object
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		http.Error(w, "Failed to get commit: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the tree from the commit
	tree, err := commit.Tree()
	if err != nil {
		http.Error(w, "Failed to get tree: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Find the file in the tree
	file, err := tree.File("serve/source.html")
	if err != nil {
		http.Error(w, "File not found in branch: "+err.Error(), http.StatusNotFound)
		return
	}

	// Open the file reader
	reader, err := file.Reader()
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	// Serve the content
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, reader)
}

func SourceFromBranch(w http.ResponseWriter, repo *git.Repository, branchName string) {
	repo.Fetch(&git.FetchOptions{RemoteName: "origin"})

	refName := plumbing.ReferenceName("refs/remotes/origin/" + branchName)
	ref, err := repo.Reference(refName, true)
	if err != nil {
		http.Error(w, "Branch not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// retrieves the commit object using the ref hash
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		http.Error(w, "Failed to get commit: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tree, err := commit.Tree()
	if err != nil {
		http.Error(w, "Failed to get tree: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := tree.File("serve/source.html")
	if err != nil {
		http.Error(w, "File not found in branch: "+err.Error(), http.StatusNotFound)
		return
	}

	reader, err := file.Reader()
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, reader)
}
