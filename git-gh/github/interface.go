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
package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

type GithubRepository struct {
	Owner  string
	Repo   string
	Client *github.Client
}

type GithubPullRequest struct {
	ID      int64
	Commits []string
}

type GitRepository struct {
	Branch string
	Head   string
	Label  string
	Base   string
}

func NewGithubClient(ctx context.Context, owner, repo string) *GithubRepository {
	client := github.NewClient(nil)
	return &GithubRepository{
		Owner:  owner,
		Repo:   repo,
		Client: client,
	}
}

func (g *GithubRepository) Pulls(ctx context.Context, log *logrus.Logger) ([]*github.PullRequest, error) {
	pulls, err := g.OpenPullRequests(ctx, log)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return pulls, nil
}

func (g *GithubRepository) PullsWithLabel(ctx context.Context, labels []string, log *logrus.Logger) ([]*github.PullRequest, error) {
	var filteredPulls []*github.PullRequest

	pulls, err := g.OpenPullRequests(ctx, log)
	if err != nil {
		fmt.Println(err)
		return filteredPulls, err
	}

	for _, label := range labels {
		log.WithFields(logrus.Fields{
			"label": label,
		}).Info("Searching")
		for _, pr := range pulls {
			prLabels := pr.Labels
			for _, ghLabel := range prLabels {
				labelName := ghLabel.Name
				if label == *labelName {
					filteredPulls = append(filteredPulls, pr)
				}
			}
		}
	}

	return filteredPulls, err
}

func (g *GithubRepository) RefForLabel(ctx context.Context, labels []string, log *logrus.Logger) ([]string, error) {
	pulls, err := g.PullsWithLabel(ctx, labels, log)
	if err != nil {
		return []string{}, err
	}

	var branches []string
	for _, pull := range pulls {
		head := pull.Head
		branch := head.Ref
		branches = append(branches, *branch)
	}

	return branches, nil
}

func (g *GithubRepository) OpenPullRequests(ctx context.Context, log *logrus.Logger) ([]*github.PullRequest, error) {
	pulls, _, err := g.Client.PullRequests.List(ctx, g.Owner, g.Repo, &github.PullRequestListOptions{State: "open"})

	if err != nil {
		return []*github.PullRequest{}, err
	}

	return pulls, nil
}
