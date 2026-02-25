package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title       string
	Description string
}

func Pars() ([]Article, error) {
	res, err := http.Get("https://habr.com/ru/news/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []Article

	doc.Find("article.tm-articles-list__item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a.tm-title__link").Text()
		description := s.Find("div.article-formatted-body").Text()

		articles = append(articles, Article{
			Title:       title,
			Description: description,
		})
		fmt.Println(title, description)
	})

	return articles, nil
}
