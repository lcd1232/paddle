package paddle

import (
	"context"
	"time"
)

type OrderResponse struct {
	Checkout Checkout  `json:"checkout"`
	Lockers  []Locker  `json:"lockers"`
	Order    OrderInfo `json:"order"`
	State    string    `json:"state"`
}

type OrderInfo struct {
	Completed                  time.Time
	CouponCode                 *string       `json:"coupon_code"`
	Currency                   string        `json:"currency"`
	Customer                   OrderCustomer `json:"customer"`
	CustomerSuccessRedirectUrl string        `json:"customer_success_redirect_url"`
	FormattedTax               string        `json:"formatted_tax"`
	FormattedTotal             string        `json:"formatted_total"`
	HasLocker                  bool          `json:"has_locker"`
	IsSubscription             bool          `json:"is_subscription"`
	OrderId                    int64         `json:"order_id"`
	Quantity                   int           `json:"quantity"`
	ReceiptUrl                 string        `json:"receipt_url"`
	Total                      string        `json:"total"`
	TotalTax                   string        `json:"total_tax"`
}

type OrderCompleted struct {
	Date         time.Time `json:"date"`
	Timezone     string    `json:"timezone"`
	TimezoneType int       `json:"timezone_type"`
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
}

func (c *Client) Order(ctx context.Context, orderID string) (*OrderResponse, error) {
	panic("implement me")
}
