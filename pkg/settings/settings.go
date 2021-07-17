package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	GithubOrganization   string   `required:"true" envconfig:"GITHUB_ORGANIZATION"`
	GithubToken          string   `required:"true" envconfig:"GITHUB_TOKEN"`
	GithubRepository     string   `required:"true" envconfig:"GITHUB_REPOSITORY"`
	GithubRepositoryPath string   `required:"true" envconfig:"GITHUB_REPOSITORY_PATH"`
	GithubBranch         string   `required:"true" envconfig:"GITHUB_BRANCH"`
	WeeklyNamespace      string   `required:"true" envconfig:"WEEKLY_NAMESPACE"`
	WeeklyCommunity      string   `required:"true" envconfig:"WEEKLY_COMMUNITY"`
	WeeklyTags           []string `required:"true" envconfig:"WEEKLY_TAGS"`
	WeeklyImage          string   `required:"true" envconfig:"WEEKLY_IMAGE"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
