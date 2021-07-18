package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	github_repository "github.com/zufardhiyaulhaq/kubernetesweekly/pkg/repository/github"

	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/models"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/repository"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/scrappers/kubernetesweekly"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/settings"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/utils"
)

const URL = "https://www.cncf.io/kubeweekly/"
const CommitMessage = "create new kubernetesweekly"

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	contextLogger := log.WithFields(log.Fields{
		"function": "main",
	})
	contextLogger.Infoln("starting Kubernetesweekly")

	settings, err := settings.NewSettings()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	githubRepository := github_repository.NewGithubRepository(settings)
	kubeweeklyScrapper := kubernetesweekly.NewKubernetesWeekly(URL)

	// get existing weekly from repository
	files, err := githubRepository.GetFiles(repository.RepositoryOptions{
		Organization: settings.GithubOrganization,
		Repository:   settings.GithubRepository,
		Path:         settings.GithubRepositoryPath,
		Branch:       settings.GithubBranch,
	})
	if err != nil {
		contextLogger.Fatalln(err)
	}

	// get newest weekly
	name, err := kubeweeklyScrapper.GetName()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	// build the newest weekly filename
	filename, err := models.NewWeeklyFilenameBuilder().
		SetName(name).
		Build()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	// if newest weekly filename is exist
	for _, file := range files {
		if filename == file {
			contextLogger.Infof("%s already exist", name)
			return
		}
	}

	// get articles
	artiles, err := kubeweeklyScrapper.GetArticles()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	// build weekly spec
	weeklySpec, err := models.NewWeeklySpecBuilder().
		SetName(name).
		SetDate(utils.GetDate()).
		SetCommunity(settings.WeeklyCommunity).
		SetImage(settings.WeeklyImage).
		SetTags(settings.WeeklyTags).
		SetArticles(artiles).
		Build()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	// build weekly object
	weekly, err := models.NewWeeklyBuilder().
		SetName(name).
		SetNamespace(settings.WeeklyNamespace).
		SetSpec(weeklySpec).
		Build()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	weeklyByte, err := weekly.ToYaml()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	// create new file
	contextLogger.Infof("create file %s in the repository", filename)
	err = githubRepository.CreateFile(filename, CommitMessage, weeklyByte, repository.RepositoryOptions{
		Organization: settings.GithubOrganization,
		Repository:   settings.GithubRepository,
		Path:         settings.GithubRepositoryPath,
		Branch:       settings.GithubBranch,
	})
	if err != nil {
		contextLogger.Fatalln(err)
	}
}
