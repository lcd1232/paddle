[![Go Report Card](https://goreportcard.com/badge/github.com/lcd1232/paddle?style=flat-square)](https://goreportcard.com/report/github.com/lcd1232/paddle)
[![Workflow Status GitHub](https://img.shields.io/github/workflow/status/lcd1232/paddle/test)](https://github.com/lcd1232/paddle/actions)
[![Coverage](https://img.shields.io/codecov/c/github/lcd1232/paddle)](https://codecov.io/gh/lcd1232/paddle)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/lcd1232/paddle/v2)
[![Releases](https://img.shields.io/github/v/tag/lcd1232/paddle.svg?style=flat-square)](https://github.com/lcd1232/paddle/releases)
[![LICENSE](https://img.shields.io/github/license/lcd1232/paddle)]((https://github.com/lcd1232/paddle/blob/master/LICENSE))

# Paddle

> Go library to work with [Paddle API](https://developer.paddle.com/api-reference/intro)

## Installation

```shell
go get github.com/lcd1232/paddle/v2
```

## Quickstart

### Handle webhooks

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lcd1232/paddle/v2"
)

func main() {
	// omitting errors for simplicity
	wc, _ := paddle.NewWebhookClient("YOURPUBLICKEY") // get it from here https://vendors.paddle.com/public-key
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		_ = req.ParseForm()
		form := req.PostForm
		alertType, _ := paddle.GetAlertName(form)
		switch alertType {
		case paddle.AlertSubscriptionCreated:
			webhook, _ := wc.ParseSubscriptionCreatedWebhook(form)
			fmt.Println(webhook.NextBillDate)
		case paddle.AlertSubscriptionPaymentSucceeded:
			webhook, _ := wc.ParseSubscriptionPaymentSucceededWebhook(form)
			fmt.Println(webhook.BalanceEarnings)
		}
	}

	http.HandleFunc("/paddle", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Using API

If you don't need authentication:
```go
package main

import (
	"context"

	"github.com/lcd1232/paddle/v2"
)

func main() {
	client, err := paddle.NewClient(paddle.Settings{}) // you can use it for methods without auth
	if err != nil {
		panic(err)
	}
	order, err := client.Order(context.Background(), "219233-chre53d41f940e0-58aqh94971")
	if err != nil {
		panic(err)
	}
}
```
With authentication:
```go
package main

import (
	"context"
	"fmt"

	"github.com/lcd1232/paddle/v2"
)

func main() {
	client, err := paddle.NewClient(paddle.Settings{
		VendorID:       "123",
		VendorAuthCode: "123abc",
	})
	if err != nil {
		panic(err)
	}
	urlStr, err := client.GeneratePayLink(context.Background(), paddle.GeneratePayLinkRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(urlStr)
	// use sandbox
	client, err = paddle.NewClient(paddle.Settings{
		URL:            paddle.SandboxBaseURL,
		VendorID:       "12",
		VendorAuthCode: "12bc",
	})
	if err != nil {
		panic(err)
	}
}
```

## Project versioning

paddle uses [semantic versioning](http://semver.org). API should not change between patch and minor releases. New minor
versions may add additional features to the API.

## Licensing

The code in this project is licensed under MIT license.
