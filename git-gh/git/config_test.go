/*
Copyright © 2022 Aaqa Ishtyaq <aaqaishtyaq@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
