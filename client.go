package paddle

import "net/http"

const APIBaseURL = "https://vendors.paddle.com/api/2.0/"

type Client struct {
	client         *http.Client
	vendorID       string
	vendorAuthCode string
	apiURL         string
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		client: client,
		apiURL: APIBaseURL,
	}
}

func NewClientWithAuthentication(client *http.Client, vendorID string, vendorAuthCode string) *Client {
	c := NewClient(client)
	c.vendorID = vendorID
	c.vendorAuthCode = vendorAuthCode
	return c
}
