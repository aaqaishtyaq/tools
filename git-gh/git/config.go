package git

import (
	"go/build"
	"os"
	"path/filepath"
)

var GitRepoPaths = []string{
	"Developer",
	"workspace",
	"Workspace",
}

var GithubRemote = "github.com"

func GoGitRepoPath(owner, repo string) string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	gosrc := filepath.Join(gopath, "src")
	repoPath := filepath.Join(gosrc, GithubRemote, owner, repo)
	return repoPath
}

func RepoPath(owner, repo string) (string, bool) {
	repoPath := GoGitRepoPath(owner, repo)
	if ok := ValidPath(repoPath); ok {
		return repoPath, true
	}

	if path, ok := PathMatrix(repo); ok {
		return path, true
	}

	return "", false
}

func PathMatrix(repo string) (string, bool) {
	for _, suite := range GitRepoPaths {
		path := filepath.Join(suite, repo)
		if ok := ValidPath(path); ok {
			return path, ok
		}
	}

	return "", false
}

func ValidPath(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}

	return false
}
