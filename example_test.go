package paddle

import (
	"context"
	"fmt"
)

func ExampleNewClient() {
	client, err := NewClient(Settings{
		VendorID:       "123",
		VendorAuthCode: "123ab",
	})
	if err != nil {
		panic(err)
	}
	url, err := client.GeneratePayLink(context.Background(), GeneratePayLinkRequest{
		ProductID: 25,
		Prices: map[string]string{
			"USD": "25.45",
			"EUR": "21.39",
			"RUB": "999",
		},
		Affiliates: map[string]string{
			"1234": "0.50",
			"1235": "0.25",
		},
		CustomerEmail:    "user@example.com",
		MarketingConsent: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
