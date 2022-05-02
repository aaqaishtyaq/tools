package git

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aaqaishtyaq/tools/git-gh/github"
	"github.com/aaqaishtyaq/tools/git-gh/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/sirupsen/logrus"
)

type GitRepository struct {
	Repo   string
	Owner  string
	Remote string
	Path   string
}

func NewGitRepository(owner, repo string, log *logrus.Logger) (*GitRepository, error) {
	le := log.WithFields(
		logrus.Fields{
			"owner": owner,
			"repo":  repo,
		},
	)
	repoPath, ok := RepoPath(owner, repo)
	if !ok {
		le.WithField("path", repoPath).Error("Repository does not exist on disk")

		return &GitRepository{}, errors.New("git repo not found on disk")
	}

	le.WithField("path", repoPath).Info("Found git repository")

	return &GitRepository{
		Repo:   repo,
		Owner:  owner,
		Remote: GithubRemote,
		Path:   repoPath,
	}, nil
}

func (g *GitRepository) Squash(ctx context.Context, labels []string, log *logrus.Logger) error {
	branches, err := g.BranchesFromLabel(ctx, labels, log)
	if err != nil {
		return err
	}

	branch := branches[0]
	err = g.RebaseandSquash(ctx, branch, log)
	if err != nil {
		return err
	}

	return nil
}

func (g *GitRepository) BranchesFromLabel(ctx context.Context, labels []string, log *logrus.Logger) ([]string, error) {
	client := github.NewGithubClient(ctx, g.Owner, g.Repo)
	branches, err := client.RefForLabel(ctx, labels, log)
	if err != nil {
		log.Error("Unable to get Refs for the branches")
		return nil, err
	}

	branches = utils.UniqueStrings(branches)
	for _, b := range branches {
		log.WithField("branch", b).Info("Found")
	}

	if len(branches) == 0 {
		return nil, errors.New("no git branches found for the given PR label")
	}

	return branches, nil
}

func (g *GitRepository) RebaseandSquash(ctx context.Context, branch string, log *logrus.Logger) error {
	le := log.WithFields(logrus.Fields{
		"path":  g.Path,
		"owner": g.Owner,
		"repo":  g.Repo,
	})

	rand := utils.GenerateRand(5)
	tempBranchName := fmt.Sprintf("%s/%s/%s", g.Owner, g.Repo, rand)

	// Instantiate a new repository targeting the given path
	r, err := git.PlainOpen(g.Path)
	if err != nil {
		return err
	}

	le.Info("git show-ref --head HEAD")
	ref, err := r.Head()
	if err != nil {
		return err
	}
	le.Info(ref.Hash())

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Checkout to a new branch
	le.Info("checkout to new branch")
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(tempBranchName),
		Create: true,
	})
	if err != nil {
		return err
	}

	sshKeyPath, err := utils.DefaultSSHKeyPath()
	if err != nil {
		return err
	}

	// user, err := user.Current()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// publicKeys, err := ssh.NewPublicKeysFromFile("aaqa@hackerrank.com", sshKey, "")
	// if err != nil {
	// 	return err
	// }

	sshKey, err := ioutil.ReadFile(sshKeyPath)
	if err != nil {
		return err
	}

	publicKey, err := ssh.NewPublicKeys("aaqaishtyaq", []byte(sshKey), "")
	if err != nil {
		return err
	}

	le.Info("Pulling the latest commits")
	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.PullContext(ctx, &git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(branch),
		Auth:          publicKey,
		Progress:      os.Stdout,
	})

	if err != nil {
		return err
	}

	le.Info("git show-ref --head HEAD post rebase")
	ref, err = r.Head()
	if err != nil {
		return err
	}
	le.Info(ref.Hash())
	return nil
}
