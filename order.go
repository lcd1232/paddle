package paddle

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type OrderResponse struct {
	Checkout Checkout  `json:"checkout"`
	Lockers  []Locker  `json:"lockers"`
	Order    OrderInfo `json:"order"`
	State    string    `json:"state"`
}

type OrderInfo struct {
	Completed                  OrderCompleted `json:"completed"`
	CouponCode                 *string        `json:"coupon_code"`
	Currency                   string         `json:"currency"`
	Customer                   OrderCustomer  `json:"customer"`
	CustomerSuccessRedirectUrl string         `json:"customer_success_redirect_url"`
	FormattedTax               string         `json:"formatted_tax"`
	FormattedTotal             string         `json:"formatted_total"`
	HasLocker                  bool           `json:"has_locker"`
	IsSubscription             bool           `json:"is_subscription"`
	OrderId                    int64          `json:"order_id"`
	Quantity                   int            `json:"quantity"`
	ReceiptUrl                 string         `json:"receipt_url"`
	Total                      string         `json:"total"`
	TotalTax                   string         `json:"total_tax"`
}

type OrderCompleted struct {
	Date         Time   `json:"date"`
	Timezone     string `json:"timezone"`
	TimezoneType int    `json:"timezone_type"`
}

type OrderCustomer struct {
	Email            string `json:"email"`
	MarketingConsent bool   `json:"marketing_consent"`
}

type Locker struct {
	Download     string `json:"download"`
	Instructions string `json:"instructions"`
	LicenseCode  string `json:"license_code"`
	LockerId     int    `json:"locker_id"`
	ProductId    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
}

type Checkout struct {
	CheckoutId string `json:"checkout_id"`
	ImageUrl   string `json:"image_url"`
	Title      string `json:"title"`
}

type orderResponse struct {
	baseAPIResponse
	OrderResponse
}

type orderRequest struct {
	CheckoutID string `schema:"checkout_id"`
}

// Order return order details.
// For more - https://developer.paddle.com/api-reference/fea392d1e2f4f-get-order-details
func (c *Client) Order(ctx context.Context, checkoutID string) (*OrderResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, "1.0/order", orderRequest{
		CheckoutID: checkoutID,
	},
		true,
		false,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data orderResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, errors.WithStack(err)
		}
		if data.IsSuccess() {
			return &data.OrderResponse, nil
		}
		return nil, &data
	}
	return nil, errors.Errorf("got status code: %d", resp.StatusCode)
}
