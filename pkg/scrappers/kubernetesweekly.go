package scrappers

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

type KubernetesWeekly struct{}

func (s *KubernetesWeekly) GetWeekly() []communityv1alpha1.ArticleSpec {
	var articles []communityv1alpha1.ArticleSpec

	// scrapper logic here
	// your scrapper must populate articles

	response, err := http.Get("https://kubeweekly.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("td#templateBody").ChildrenFiltered(".mcnTextBlock").Each(func(i int, s0 *goquery.Selection) {
		s0.Find("a").Each(func(i int, s1 *goquery.Selection) {
			var article communityv1alpha1.ArticleSpec

			article.Title = s1.Text()
			article.Type = s0.Find("h1").Text()
			article.Url, _ = s1.Attr("href")

			articles = append(articles, article)
		})
	})

	return articles
}

func (s *KubernetesWeekly) GetWeeklyName() string {
	var weeklyName string

	// scrapper logic here
	// your scrapper must populate weeklyName

	response, err := http.Get("https://kubeweekly.io/")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("td#templateHeader").ChildrenFiltered(".mcnTextBlock").Each(func(i int, s0 *goquery.Selection) {
		text := strings.ToLower(s0.Find("strong").Text())
		regex, _ := regexp.Compile(`kubeweekly(.*)#\d+`)
		weeklyName = regex.FindString(text)
	})

	return weeklyName
}
