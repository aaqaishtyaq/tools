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

func (g *GithubRepository) Pulls(ctx context.Context, log *logrus.Logger) error {
	pulls, err := g.OpenPullRequests(ctx, log)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Github Information %s,  %s\n", g.Owner, g.Repo)
	fmt.Println("Count of the pulls")
	fmt.Println(len(pulls))
	for _, r := range pulls {
		head := r.Head
		gitRef := head.Ref
		fmt.Println(*r.Number)
		fmt.Println("------------")
		fmt.Println(*gitRef)
	}

	return nil
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
