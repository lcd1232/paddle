package paddle

import "net/http"

type Client struct {
	client *http.Client
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		client: client,
	}
}
