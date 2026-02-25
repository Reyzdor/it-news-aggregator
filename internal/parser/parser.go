package parser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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

	jsonData, err := json.Marshal(articles)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonData))

	return articles, nil
}
