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

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   appName,
		Short: "Git Github extension",
		Long: `Github extension for the git(1) cli.
Allow manupulation of commits and PRs. Interacts
with Github's REST APIs.
`,
	}
	appName = "git-gh"
	log     *logrus.Logger
)

// ExecuteContext is the same as cmd.Execute(), but sets the ctx on the command.
func ExecuteContext(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
