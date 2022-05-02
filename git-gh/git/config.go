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
	"go/build"
	"os"
	"path/filepath"
)

var (
	GithubRemote = "github.com"
	GitRepoPaths = []string{
		"Developer",
		"workspace",
		"Workspace",
		"workspace",
		"project",
	}
)

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
