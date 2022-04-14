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

type GithubPullRequest struct {
	ID      int64
	Commits []string
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
	pulls, err := g.OpenPullRequests(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fmt.Println(pulls)
	fmt.Printf("Github Information %s,  %s\n", g.Owner, g.Repo)
	fmt.Println("Count of the pulls")
	fmt.Println(len(pulls))
	for _, r := range pulls {
		fmt.Println(*r.Number)
		fmt.Println("------------")
		fmt.Println(*&r.GetCommits)
		// fmt.Println(*r.Commits)
		// fmt.Println(*r.CommitsURL)
	}

	return nil
}

func (g *GithubRepository) OpenPullRequests(ctx context.Context) ([]*github.PullRequest, error) {
	pulls, _, err := g.Client.PullRequests.List(ctx, g.Owner, g.Repo, &github.PullRequestListOptions{State: "open"})

	// fmt.Println(pulls)
	// fmt.Println(res)
	// fmt.Println("\n\n\n\n")
	// fmt.Println("error")
	// fmt.Println(err.Error())
	// fmt.Println("\n\n\n\n")
	if err != nil {
		return []*github.PullRequest{}, err
	}

	return pulls, nil
}
