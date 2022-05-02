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
package cmd

import (
	"context"

	"github.com/aaqaishtyaq/tools/git-gh/git"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// squashCmd represents the fetch command
	squashCmd = &cobra.Command{
		Use:   "squash owner repo [..labels] [flags]",
		Short: "Squash and Rebase all the commits from Github PRs given the label",
		Long: `Fetches Github pull request's metadata
		and squash and rebase the commits from the PR head's`,
		Args:       cobra.MinimumNArgs(2),
		ArgAliases: []string{"owner", "repos"},
		Run:        contextAdder.withContext(squash),
	}
)

func init() {
	rootCmd.AddCommand(squashCmd)
}

func squash(ctx context.Context, cmd *cobra.Command, args []string) {
	owner := args[0]
	repo := args[1]
	labels := args[2:]

	log = logrus.New()

	gRepo, err := git.NewGitRepository(owner, repo, log)
	if err != nil {
		log.Warn("Unable to find git repository")
		return
	}

	err = gRepo.Squash(ctx, labels, log)
	if err != nil {
		log.Error(err)
	}
}
