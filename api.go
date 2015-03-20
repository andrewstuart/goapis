package apis

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Url string
	c   http.Client
	q   Query
}

func NewClient(url string) *Client {
	return &Client{url, http.Client{}, Query{}}
}

//DefaultQuery sets a map to be used for each query and extended
func (c *Client) DefaultQuery(q Query) {
	c.q = q
}

func (c *Client) Get(u string, q Query) (*http.Response, error) {
	qs := url.Values{}

	//Get defaults
	for k, v := range c.q {
		qs[k] = []string{v}
	}

	//Add instance query
	for k, v := range q {
		qs[k] = []string{v}
	}

	u = fmt.Sprintf("%s/%s?%s", c.Url, u, qs.Encode())

	return c.c.Get(u)
}
