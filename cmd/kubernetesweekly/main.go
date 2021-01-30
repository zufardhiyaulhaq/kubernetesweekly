package main

import (
	"log"
	"strings"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	"github.com/zufardhiyaulhaq/kubernetesweekly/handlers"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/scrappers"
	"gopkg.in/yaml.v2"
)

func main() {
	handler := handlers.Github{}
	handler.Start()

	// Init scrapper
	scrapper := scrappers.KubernetesWeekly{}

	// get latest weekly from handlers
	recentWeeklyNames := GetFiles(handler)

	// compare weekly logic
	// compare newest weekly from scrapper
	// with latest list of weekly from datastore
	newestWeeklyName := scrapper.GetWeeklyName()

	// remove unsupported character from name
	newestWeeklyName = strings.Replace(newestWeeklyName, "#", "", -1)

	for _, v := range recentWeeklyNames {
		if strings.ToLower(strings.ReplaceAll(newestWeeklyName, " ", "-"))+".yaml" == v {
			log.Println("Weekly already in datastore")
			return
		}
	}

	// Scapper logic
	// must return list fo ArticleSpec defined in community-operator
	// communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	var weekly []communityv1alpha1.ArticleSpec
	weekly = scrapper.GetWeekly()

	// delete empty data.
	for index, article := range weekly {
		if article.Type == "" || article.Title == "" || article.Url == "" {
			weekly[index] = weekly[len(weekly)-1]
			weekly[len(weekly)-1] = communityv1alpha1.ArticleSpec{}
			weekly = weekly[:len(weekly)-1]
		}
	}

	// Init builder
	builder := Builder{}

	// Build
	builder.build(newestWeeklyName, weekly)

	// Add to Github
	crd, err := yaml.Marshal(builder)
	if err != nil {
		log.Fatal(err)
	}

	commitMessage := "Adding " + newestWeeklyName
	CreateFile(handler, strings.ToLower(strings.ReplaceAll(newestWeeklyName, " ", "-"))+".yaml", commitMessage, crd)
}
