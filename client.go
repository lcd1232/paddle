package paddle

import "net/http"

type Client struct {
	client         *http.Client
	vendorID       string
	vendorAuthCode string
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		client: client,
	}
}

func NewClientWithAuthentication(client *http.Client, vendorID string, vendorAuthCode string) *Client {
	c := NewClient(client)
	c.vendorID = vendorID
	c.vendorAuthCode = vendorAuthCode
	return c
}
