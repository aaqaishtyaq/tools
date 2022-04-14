/*
Copyright © 2022 Aaqa Ishtyaq aaqaishtyaq@gmail.com

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
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "git-gh",
		Short: "Git Github extension",
		Long: `Github extension for the git(1) cli.
Allow manupulation of commits and PRs. Interacts
with Github's REST APIs.
`,
		// Run: runWithoutContext,
	}
	contextAdder ctxAdder
)

func init() {
	// This is a common pattern for structuring programs using the
	// flag or cobra packages.
	cmd1 := &cobra.Command{
		Use: "cmd1",
		Run: addContext(context.Background(), runWithContext),
	}

	cmd2 := &cobra.Command{
		Use: "cmd2",
		Run: contextAdder.withContext(runWithContext),
	}

	rootCmd.AddCommand(cmd1)

	rootCmd.AddCommand(cmd2)
}

func runWithoutContext(cmd *cobra.Command, args []string) {
	cmd.Printf("called as: %s\n", cmd.CalledAs())
	cmd.Printf("name: %s\n", cmd.Name())
	cmd.Printf("args: %v\n", args)
}

func runWithContext(ctx context.Context, cmd *cobra.Command, args []string) {
	cmd.Printf("called as: %s\n", cmd.CalledAs())
	cmd.Printf("name: %s\n", cmd.Name())
	cmd.Printf("ctx: %s\n", ctx)
	cmd.Printf("args: %v\n", args)

	tick := time.NewTicker(500 * time.Millisecond)
	defer tick.Stop()

	for n := 0; n < 10; n++ {
		cmd.Println("Working...")

		select {
		case <-tick.C:

		case <-ctx.Done():
			cmd.Println("Context done")
			return
		}
	}
}

type commandWithContext func(context.Context, *cobra.Command, []string)

func addContext(ctx context.Context, fn commandWithContext) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		fn(ctx, cmd, args)
	}
}

type ctxAdder struct {
	ctx context.Context
}

func (c *ctxAdder) setContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *ctxAdder) withContext(fn commandWithContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fn(c.ctx, cmd, args)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) error {
	contextAdder.setContext(ctx)
	return rootCmd.Execute()
}
