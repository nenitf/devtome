package devtome

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nenitf/devtome/pkg/devto"
	"github.com/pkg/errors"
)

// THANKS:
// https://golang.cafe/blog/golang-httptest-example.html
// https://dev.to/eminetto/testing-apis-in-golang-using-apitest-1860
// https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html

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

func (c Client) GetAll() ([]Article, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", c.url+"/api/articles/me/all", nil)
	if err != nil {
		return []Article{}, errors.Wrap(err, "unable to complete Get request")
	}

	req.Header = http.Header{
		"api-key": []string{c.token},
	}

	res, err := client.Do(req)
	if err != nil {
		return []Article{}, errors.Wrap(err, "unable to complete Get request")
	}

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Article{}, errors.Wrap(err, "unable to read response data")
	}

	var articles []Article
	err = json.Unmarshal(out, &articles)
	if err != nil {
		return []Article{}, errors.Wrap(err, "unable to read response data")
	}

	return articles, nil
}
