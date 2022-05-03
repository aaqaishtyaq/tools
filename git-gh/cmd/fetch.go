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
	"github.com/aaqaishtyaq/tools/git-gh/github"
	"github.com/aaqaishtyaq/tools/git-gh/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch owner repos [..repos] [flags]",
	Short: "Fetch Github PR information",
	Long: `Fetches Github pull request's metadata
		and returns the list of Pull requests and there information.`,
	Args:       cobra.MinimumNArgs(2),
	ArgAliases: []string{"owner", "repos"},
	RunE:       fetch,
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func fetch(cmd *cobra.Command, args []string) error {
	owner := args[0]
	repo := args[1]
	labels := args[2:]

	log = logrus.New()
	ctx := cmd.Root().Context()
	client := github.NewGithubClient(ctx, owner, repo)
	branches, err := client.RefForLabel(ctx, labels, log)
	if err != nil {
		log.WithError(err)
		return err
	}

	branches = utils.UniqueStrings(branches)
	for _, b := range branches {
		log.WithFields(logrus.Fields{
			"branch": b,
		}).Info("Found")
	}

	return nil
}
