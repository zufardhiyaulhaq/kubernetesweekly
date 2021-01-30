package main

import (
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/zufardhiyaulhaq/kubernetesweekly/handlers"
)

func GetFiles(handler interface{}) []string {
	var files []string

	switch handler.(type) {
	case handlers.Github:
		github := handlers.Github{}
		mapstructure.Decode(handler, &github)

		files = github.GetFiles()
	}

	return files
}

func CreateFile(handler interface{}, fileName string, commitMessage string, fileData []byte) {
	switch handler.(type) {
	case handlers.Github:
		github := handlers.Github{}
		mapstructure.Decode(handler, &github)

		github.CreateFile(fileName, commitMessage, fileData)
	}
}

func GetDate() string {
	dt := time.Now()
	Date := dt.Format("01-02-2006")
	Date = strings.ReplaceAll(Date, "-", "/")

	return Date
}
