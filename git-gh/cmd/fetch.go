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

	"github.com/aaqaishtyaq/tools/git-gh/github"
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

	Run: contextAdder.withContext(fetch),
}

// var owner, repo string

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().StringVarP(&owner, "owner", "o", "", "Owner of the git repository")
	// fetchCmd.Flags().StringVarP(&repo, "repo", "r", "", "the git repository")
}

// func fetch(args []string) {
// 	owner := args[0]
// 	repos := args[1:]

// 	var repos github.GithubRepositories
// 	for _, r := range repos {
// 		client := github.NewGithubClient(owner, r)
// 	}
// }

func fetch(ctx context.Context, cmd *cobra.Command, args []string) {
	owner := args[0]
	repos := args[1:]

	// cmd.Printf("CC -- ctx: %s\n", ctx)
	// var repos github.GithubRepositories
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*200))
	// defer func() {
	// 	fmt.Println("Cancelling the context...")
	// 	cancel()
	// }()
	for _, r := range repos {
		client := github.NewGithubClient(ctx, owner, r)
		client.Pulls(ctx)
	}
}
