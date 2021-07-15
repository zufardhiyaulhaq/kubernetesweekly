package scrappers

import (
	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

type Scrapper interface {
	GetName() (string, error)
	GetArticles() ([]communityv1alpha1.ArticleSpec, error)
}
