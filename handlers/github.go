package handlers

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	Organization string
	Token        string
	Repository   string
	Path         string
	Branch       string
	Options      *github.RepositoryContentGetOptions
	Session      *github.Client
}

func (g *Github) Start() {
	g.credentials()

	context := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Token},
	)
	tc := oauth2.NewClient(context, ts)

	g.Session = github.NewClient(tc)
}

func (g *Github) credentials() {
	g.Organization = os.Getenv("GITHUB_ORGANIZATION")
	g.Token = os.Getenv("GITHUB_TOKEN")
	g.Repository = os.Getenv("GITHUB_REPOSITORY")
	g.Path = os.Getenv("GITHUB_REPOSITORY_PATH")
	g.Branch = os.Getenv("GITHUB_BRANCH")
	g.Options = &github.RepositoryContentGetOptions{Ref: g.Branch}
}

func (g *Github) GetFiles() []string {
	var files []string
	context := context.Background()

	_, dir, _, err := g.Session.Repositories.GetContents(context, g.Organization, g.Repository, g.Path, g.Options)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range dir {
		files = append(files, *v.Name)
	}

	return files
}

func (g *Github) CreateFile(fileName string, commitMessage string, fileData []byte) {
	context := context.Background()

	res, _, _, err := g.Session.Repositories.GetContents(context, g.Organization, g.Repository, g.Path, g.Options)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = g.Session.Repositories.CreateFile(
		context,
		g.Organization,
		g.Repository,
		g.Path+fileName,
		&github.RepositoryContentFileOptions{
			Message: &commitMessage,
			Content: fileData,
			Branch:  &g.Branch,
			SHA:     github.String(res.GetSHA()),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
