package github_repository

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/repository"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/settings"
	"golang.org/x/oauth2"
)

type GithubRepository struct {
	Context context.Context
	Client  *github.Client
}

func (g GithubRepository) GetFiles(options repository.RepositoryOptions) ([]string, error) {
	var files []string

	_, contents, _, err := g.Client.Repositories.GetContents(
		g.Context,
		options.Organization,
		options.Repository,
		options.Path,
		&github.RepositoryContentGetOptions{
			Ref: options.Branch,
		})
	if err != nil {
		return []string{}, err
	}

	for _, content := range contents {
		files = append(files, *content.Name)
	}

	return files, nil
}

func (g GithubRepository) CreateFile(name string, commit string, data []byte, options repository.RepositoryOptions) error {
	res, _, _, err := g.Client.Repositories.GetContents(
		g.Context,
		options.Organization,
		options.Repository,
		options.Path,
		&github.RepositoryContentGetOptions{
			Ref: options.Branch,
		})
	if err != nil {
		return err
	}

	_, _, err = g.Client.Repositories.CreateFile(
		g.Context,
		options.Organization,
		options.Repository,
		options.Path+name,
		&github.RepositoryContentFileOptions{
			Message: &commit,
			Content: data,
			Branch:  &options.Branch,
			SHA:     github.String(res.GetSHA()),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func NewGithubRepository(settings settings.Settings) repository.Repository {
	context := context.Background()

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: settings.GithubToken},
	)
	httpClient := oauth2.NewClient(context, tokenSource)

	return GithubRepository{
		Context: context,
		Client:  github.NewClient(httpClient),
	}
}
