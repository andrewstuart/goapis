package apis

import (
	"fmt"
	"io"
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
	u = fmt.Sprintf("%s/%s?%s", c.Url, u, c.getQuery(&q).Encode())
	return c.c.Get(u)
}

func (c *Client) getQuery(q *Query) *url.Values {
	qs := url.Values{}

	//Get defaults
	for k, v := range c.q {
		qs[k] = []string{v}
	}

	if q != nil {
		//Add instance query
		for k, v := range *q {
			qs[k] = []string{v}
		}
	}

	return &qs
}

func (c *Client) Post(u, t string, q Query, r io.Reader) (*http.Response, error) {
	u = fmt.Sprintf("%s/%s?%s", c.Url, u, c.getQuery(&q).Encode())
	return c.c.Post(u, t, r)
}
