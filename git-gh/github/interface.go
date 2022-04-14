package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type GithubRepository struct {
	Owner  string
	Repo   string
	Client *github.Client
}

type GithubRepositories struct {
	Repositories []GitRepository
}

type GitRepository struct {
}

type GithubResponse struct {
}

func NewGithubClient(ctx context.Context, owner, repo string) *GithubRepository {
	client := github.NewClient(nil)
	return &GithubRepository{
		Owner:  owner,
		Repo:   repo,
		Client: client,
	}
}

func (g *GithubRepository) Pulls(ctx context.Context) error {
	PRResp, _, err := g.Client.PullRequests.List(ctx, g.Owner, g.Repo, &github.PullRequestListOptions{State: "open"})

	fmt.Printf("Github Information %s,  %s\n", g.Owner, g.Repo)
	for _, r := range PRResp {
		// fmt.Println(r)
		num := r.Number
		fmt.Println(*num)
	}
	if err != nil {
		return err
	}

	return nil
}
