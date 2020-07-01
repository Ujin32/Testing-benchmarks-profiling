package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "https://jsonplaceholder.typicode.com"

type Post interface {
	GetList() ([]PostData, error)
}

type postClient struct{}

func NewPostClient() Post {
	return &postClient{}
}

func (c *postClient) GetList() ([]PostData, error) {
	client := http.Client{Timeout: time.Second * 2}

	res, err := client.Get(url + "/posts")
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var posts []PostData
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal []Post: %w", err)
	}

	return posts, nil
}

type PostData struct {
	ID     int
	Title  string
	Body   string
	UserID int
}
