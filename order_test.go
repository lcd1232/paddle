package paddle

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrder(t *testing.T) {
	for _, tc := range []struct {
		name         string
		checkoutID   string
		responseCode int
		responseBody func(t *testing.T) []byte
		context      func() (context.Context, context.CancelFunc)
		want         *OrderResponse
		wantErr      bool
		wantQuery    url.Values
	}{
		{
			name:         "success order",
			checkoutID:   "219233-chre53d41f940e0-58aqh94971",
			responseCode: http.StatusOK,
			responseBody: func(t *testing.T) []byte {
				b, err := ioutil.ReadFile("testdata/order-1.json")
				require.NoError(t, err)
				return b
			},
			want: &OrderResponse{
				Checkout: Checkout{
					CheckoutId: "219233-chre53d41f940e0-58aqh94971",
					ImageUrl:   "https://paddle.s3.amazonaws.com/user/91/XWsPdfmISG6W5fgX5t5C_icon.png",
					Title:      "My Product",
				},
				Lockers: []Locker{
					{
						Download:     "https://mysite.com/download/my-app",
						Instructions: "Simply enter your license code and click 'Activate'.",
						LicenseCode:  "ABC-123",
						LockerId:     1127139,
						ProductId:    514032,
						ProductName:  "My Product Name",
					},
				},
				Order: OrderInfo{
					Completed:  time.Date(2019, 8, 1, 18, 24, 35, 0, time.UTC),
					CouponCode: nil,
					Currency:   "GBP",
					Customer: OrderCustomer{
						Email:            "example@paddle.com",
						MarketingConsent: true,
					},
					CustomerSuccessRedirectUrl: "https://example.com/success",
					FormattedTax:               "£1.73",
					FormattedTotal:             "£9.99",
					HasLocker:                  true,
					IsSubscription:             false,
					OrderId:                    123456,
					Quantity:                   1,
					ReceiptUrl:                 "https://my.paddle.com/receipt/826289/3219233-chre53d41f940e0-58aqh94971",
					Total:                      "9.99",
					TotalTax:                   "1.73",
				},
				State: "processed",
			},
			wantQuery: url.Values{
				"checkout_id": {"219233-chre53d41f940e0-58aqh94971"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			WithTestServer(t, tc.responseCode, tc.responseBody(t), func(url string, rCh <-chan *http.Request) {
				ctx := context.Background()
				if tc.context != nil {
					var cancel context.CancelFunc
					ctx, cancel = tc.context()
					defer cancel()
				}
				c, err := NewClient(Settings{
					URL: url,
				})
				require.NoError(t, err)
				order, err := c.Order(ctx, tc.checkoutID)
				if tc.wantErr {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
				assert.Equal(t, tc.want, order)
				r := <-rCh
				require.NoError(t, r.ParseForm())
				assert.Equal(t, tc.wantQuery, r.URL.Query())
				assert.Equal(t, "/1.0/order", r.URL.Path)
			})
		})
	}
}
