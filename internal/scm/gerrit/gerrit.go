package gerrit

import (
	"context"
	gogerrit "github.com/andygrunwald/go-gerrit"
	"github.com/lindell/multi-gitter/internal/scm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type Gerrit struct {
	client     *gogerrit.Client
	baseUrl    string
	username   string
	token      string
	repoSearch string
}

func New(username, token, baseURL, repoSearch string) (*Gerrit, error) {
	ctx := context.Background()
	client, err := gogerrit.NewClient(ctx, baseURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gerrit client")
	}

	client.Authentication.SetBasicAuth(username, token)

	return &Gerrit{
		client:     client,
		baseUrl:    baseURL,
		username:   username,
		token:      token,
		repoSearch: repoSearch,
	}, nil
}

func (g Gerrit) GetRepositories(ctx context.Context) ([]scm.Repository, error) {
	opt := &gogerrit.ProjectOptions{
		Description: true,
		Substring:   g.repoSearch,
		Type:        "CODE",
		ProjectBaseOptions: gogerrit.ProjectBaseOptions{
			Limit: 2500, // Maybe we should make this configurable
		},
	}
	projects, _, err := g.client.Projects.ListProjects(ctx, opt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list projects")
	}

	repos := make([]scm.Repository, 0)
	for name, project := range *projects {
		if project.State != "ACTIVE" {
			log.Debug("Skipping repository since state is not ACTIVE")
			continue
		}

		repo, err := g.convertRepo(name)
		if err != nil {
			return nil, err
		}

		repos = append(repos, repo)
	}

	return repos, nil
}

func (g Gerrit) convertRepo(name string) (repository, error) {
	// Note: maybe we should support cloning via ssh
	u, err := url.Parse(g.baseUrl)
	if err != nil {
		return repository{}, err
	}
	u.User = url.UserPassword(g.username, g.token)
	u.Path = "/a/" + name
	repoUrl := u.String()

	return repository{
		url:           repoUrl,
		name:          name,
		defaultBranch: "master", // Some projects might have a different default branch
	}, nil
}

func (Gerrit) CreatePullRequest(ctx context.Context, repo scm.Repository, prRepo scm.Repository, newPR scm.NewPullRequest) (scm.PullRequest, error) {
	log.Infof("CreatePullRequest %v, %v, %v", ctx, repo, prRepo, newPR)
	return nil, errors.New("implement me")
}

func (Gerrit) UpdatePullRequest(ctx context.Context, repo scm.Repository, pullReq scm.PullRequest, updatedPR scm.NewPullRequest) (scm.PullRequest, error) {
	log.Infof("UpdatePullRequest %v, %v, %v", ctx, repo, pullReq, updatedPR)
	return nil, errors.New("implement me")
}

func (Gerrit) GetPullRequests(ctx context.Context, branchName string) ([]scm.PullRequest, error) {
	log.Infof("GetPullRequests %v, %s", ctx, branchName)
	return nil, errors.New("implement me")
}

func (Gerrit) GetOpenPullRequest(ctx context.Context, repo scm.Repository, branchName string) (scm.PullRequest, error) {
	log.Infof("GetOpenPullRequest %v, %v, %s", ctx, repo, branchName)
	return nil, errors.New("implement me")
}

func (Gerrit) MergePullRequest(ctx context.Context, pr scm.PullRequest) error {
	log.Infof("MergePullRequest %v, %v", ctx, pr)
	return errors.New("implement me")
}

func (Gerrit) ClosePullRequest(ctx context.Context, pr scm.PullRequest) error {
	log.Infof("ClosePullRequest %v, %v", ctx, pr)
	return errors.New("implement me")
}

func (Gerrit) ForkRepository(ctx context.Context, repo scm.Repository, newOwner string) (scm.Repository, error) {
	log.Infof("ForkRepository %v, %v, %s", ctx, repo, newOwner)
	return nil, errors.New("implement me")
}
