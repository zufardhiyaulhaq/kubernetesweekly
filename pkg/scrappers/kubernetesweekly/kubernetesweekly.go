package kubernetesweekly

import (
	"net/http"
	"strings"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	log "github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/models"
	"github.com/zufardhiyaulhaq/kubernetesweekly/pkg/scrappers"
)

type KubernetesWeekly struct {
	URL string
}

func (k KubernetesWeekly) GetName() (string, error) {
	response, err := http.Get(k.URL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	name, exists := document.Find(".kubeweekly-box").First().Find("a").Attr("title")
	if !exists {
		return "", &models.KubeweeklyNameNotFoundError{}
	}

	return name, nil
}

func (k KubernetesWeekly) GetArticles() ([]communityv1alpha1.ArticleSpec, error) {
	contextLogger := log.WithFields(log.Fields{
		"scrappers": "KubernetesWeekly",
		"function":  "GetArtiles",
	})

	response, err := http.Get(k.URL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	url, exists := document.Find(".kubeweekly-box").First().Find("a").Attr("href")
	if !exists {
		return nil, &models.KubeweeklyNameNotFoundError{}
	}

	// start getting articles data
	var articles []communityv1alpha1.ArticleSpec

	response, err = http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	document, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	ICYMI, err := k.getICYMI(document)
	if err != nil {
		contextLogger.Errorf("cannot get ICYMI")
	}

	editorial, err := k.getEditorial(document)
	if err != nil {
		contextLogger.Errorf("cannot get Editorial")
	}

	technical, err := k.getTechnical(document)
	if err != nil {
		contextLogger.Errorf("cannot get Editorial")
	}

	articles = append(articles, ICYMI...)
	articles = append(articles, editorial...)
	articles = append(articles, technical...)

	return articles, nil
}

func (k KubernetesWeekly) getICYMI(document *goquery.Document) ([]communityv1alpha1.ArticleSpec, error) {
	var articles []communityv1alpha1.ArticleSpec

	document.Find("div#section-1").First().Find("table").Each(func(i int, table *goquery.Selection) {

		table.Find("h2").Each(func(i int, h2 *goquery.Selection) {
			if strings.Contains(h2.Text(), "ICYMI") {

				table.Find("a").Each(func(i int, a *goquery.Selection) {
					var article communityv1alpha1.ArticleSpec

					article.Title = a.Text()
					article.Url, _ = a.Attr("href")
					article.Type = "ICYMI"

					articles = append(articles, article)
				})
			}
		})

	})

	return articles, nil
}

func (k KubernetesWeekly) getTechnical(document *goquery.Document) ([]communityv1alpha1.ArticleSpec, error) {
	var articles []communityv1alpha1.ArticleSpec

	document.Find("div#section-1").First().Find("table").Each(func(i int, table *goquery.Selection) {

		table.Find("h2").Each(func(i int, h2 *goquery.Selection) {
			if strings.Contains(h2.Text(), "The technical") {

				table.Find("a").Each(func(i int, a *goquery.Selection) {
					var article communityv1alpha1.ArticleSpec

					article.Title = a.Text()
					article.Url, _ = a.Attr("href")
					article.Type = "Technical"

					articles = append(articles, article)
				})
			}
		})

	})

	return articles, nil
}

func (k KubernetesWeekly) getEditorial(document *goquery.Document) ([]communityv1alpha1.ArticleSpec, error) {
	var articles []communityv1alpha1.ArticleSpec

	document.Find("div#section-1").First().Find("table").Each(func(i int, table *goquery.Selection) {

		table.Find("h2").Each(func(i int, h2 *goquery.Selection) {
			if strings.Contains(h2.Text(), "The editorial") {

				table.Find("a").Each(func(i int, a *goquery.Selection) {
					var article communityv1alpha1.ArticleSpec

					article.Title = a.Text()
					article.Url, _ = a.Attr("href")
					article.Type = "Editorial"

					articles = append(articles, article)
				})
			}
		})

	})

	return articles, nil
}

func NewKubernetesWeekly(URL string) scrappers.Scrapper {
	return KubernetesWeekly{
		URL: URL,
	}
}
