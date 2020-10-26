package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParsePaymentSucceededWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    PaymentSucceeded
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":          {"payment_succeeded"},
					"alert_id":            {"27120763"},
					"balance_currency":    {"USD"},
					"balance_earnings":    {"10.00"},
					"balance_fee":         {"2.25"},
					"balance_gross":       {"2.30"},
					"balance_tax":         {"1.20"},
					"checkout_id":         {"54832806-chrea1c514a3eb5-25c3040268"},
					"country":             {"US"},
					"coupon":              {"coupon"},
					"currency":            {"USD"},
					"customer_name":       {"somename"},
					"earnings":            {"1.01"},
					"email":               {"test@example.org"},
					"event_time":          {"2020-04-28 08:42:47"},
					"fee":                 {"0.52"},
					"ip":                  {"127.0.0.1"},
					"marketing_consent":   {"1"},
					"order_id":            {"123-456"},
					"passthrough":         {"some data"},
					"payment_method":      {"card"},
					"payment_tax":         {"0.25"},
					"product_id":          {"123"},
					"product_name":        {"product1"},
					"quantity":            {"1"},
					"receipt_url":         {"http://my.paddle.com/receipt/receipt-url"},
					"sale_gross":          {"0.10"},
					"used_price_override": {"true"},
					"p_signature":         {"signature"},
				},
			},
			want: PaymentSucceeded{
				AlertName:         AlertPaymentSucceeded,
				AlertID:           "27120763",
				BalanceCurrency:   "USD",
				BalanceEarnings:   "10.00",
				BalanceFee:        "2.25",
				BalanceGross:      "2.30",
				BalanceTax:        "1.20",
				CheckoutID:        "54832806-chrea1c514a3eb5-25c3040268",
				Country:           "US",
				Coupon:            "coupon",
				Currency:          "USD",
				CustomerName:      "somename",
				Earnings:          "1.01",
				Email:             "test@example.org",
				EventTime:         time.Date(2020, 4, 28, 8, 42, 47, 0, time.UTC),
				Fee:               "0.52",
				Instalments:       "1",
				MarketingConsent:  true,
				OrderID:           "123-456",
				IP:                "127.0.0.1",
				Passthrough:       "some data",
				PaymentMethod:     PaymentMethodCard,
				PaymentTax:        "0.25",
				ProductID:         "123",
				ProductName:       "product1",
				Quantity:          1,
				ReceiptURL:        "http://my.paddle.com/receipt/receipt-url",
				SaleGross:         "0.10",
				UsedPriceOverride: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, vm := NewTestClient()
			vm.On("Verify",
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(
				nil,
			).Once()
			got, err := c.ParseSubscriptionPaymentSucceededWebhook(tt.args.form)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
			vm.AssertExpectations(t)
		})
	}
}
