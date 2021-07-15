package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	GithubOrganization   string   `envconfig:"GITHUB_ORGANIZATION"`
	GithubToken          string   `envconfig:"GITHUB_TOKEN"`
	GithubRepository     string   `envconfig:"GITHUB_REPOSITORY"`
	GithubRepositoryPath string   `envconfig:"GITHUB_REPOSITORY_PATH"`
	GithubBranch         string   `envconfig:"GITHUB_BRANCH"`
	MeetupNamespace      string   `envconfig:"NAMESPACE"`
	MeetupCommunity      string   `envconfig:"COMMUNITY"`
	MeetupTags           []string `envconfig:"TAGS"`
	MeetupImage          string   `envconfig:"IMAGE"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
