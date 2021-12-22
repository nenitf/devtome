package devtome

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nenitf/devtome/pkg/devto"
	"github.com/pkg/errors"
)

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type Client struct {
	url   string
	token string
}

func NewClient(url string, token string) Client {
	return Client{url, token}
}

func (c Client) GetAll() (articles []Article, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", c.url+"/api/articles/me/all", nil)
	if err != nil {
		return articles, errors.Wrap(err, "unable to complete Get request")
	}

	req.Header = http.Header{
		"api-key": []string{c.token},
	}

	res, err := client.Do(req)
	if err != nil {
		return articles, errors.Wrap(err, "unable to complete Get request")
	}

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return articles, errors.Wrap(err, "unable to read response data")
	}

	var articlesDevTo []devto.Article
	err = json.Unmarshal(out, &articlesDevTo)
	if err != nil {
		return articles, errors.Wrap(err, "unable to parse response body")
	}

	for _, a := range articlesDevTo {
		articles = append(articles, mapAPI2Article(a))
	}

	return
}

func mapAPI2Article(a1 devto.Article) (a2 Article) {
	a2 = Article{
		Id:          a1.Id,
		Title:       a1.Title,
		Description: a1.Description,
		Content:     a1.BodyMarkdown,
	}

	return
}
