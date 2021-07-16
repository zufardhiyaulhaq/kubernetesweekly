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
	WeeklyNamespace      string   `envconfig:"WEEKLY_NAMESPACE"`
	WeeklyCommunity      string   `envconfig:"WEEKLY_COMMUNITY"`
	WeeklyTags           []string `envconfig:"WEEKLY_TAGS"`
	WeeklyImage          string   `envconfig:"WEEKLY_IMAGE"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
