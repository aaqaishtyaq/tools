package git

import (
	"fmt"
	"go/build"
	"os"
	"testing"
)

func TestGoGitRepoPath(t *testing.T) {
	t.Run("returns unique string slice from a string slice", func(t *testing.T) {
		owner := "aaqaishtyaq"
		repo := "git-gh"

		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}

		got := GoGitRepoPath(owner, repo)
		want := fmt.Sprintf("%s/src/%s/%s/%s", gopath, GithubRemote, owner, repo)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
