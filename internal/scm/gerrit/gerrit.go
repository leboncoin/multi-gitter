package gerrit

import (
	"context"
	"github.com/lindell/multi-gitter/internal/scm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Gerrit struct{}

func New(username, token, baseURL string) (*Gerrit, error) {
	return &Gerrit{}, nil
}

func (Gerrit) GetRepositories(ctx context.Context) ([]scm.Repository, error) {
	log.Infof("GetRepositories %v", ctx)
	return nil, errors.New("implement me")
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
