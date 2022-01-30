package client

import "net/http"

type Client struct{}

func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleReq(w, r)
}

func NewClient() http.Handler {
	return new(Client)
}
