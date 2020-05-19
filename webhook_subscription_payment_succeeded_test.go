package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseSubscriptionPaymentSucceededWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    SubscriptionPaymentSucceeded
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":              {"subscription_payment_succeeded"},
					"alert_id":                {"27120763"},
					"balance_currency":        {"USD"},
					"balance_earnings":        {"10.00"},
					"balance_fee":             {"2.25"},
					"balance_gross":           {"2.30"},
					"balance_tax":             {"1.20"},
					"checkout_id":             {"54832806-chrea1c514a3eb5-25c3040268"},
					"country":                 {"US"},
					"coupon":                  {"coupon"},
					"currency":                {"USD"},
					"customer_name":           {"somename"},
					"earnings":                {"1.01"},
					"email":                   {"test@example.org"},
					"fee":                     {"0.52"},
					"event_time":              {"2020-04-28 08:42:47"},
					"initial_payment":         {"1"},
					"instalments":             {"1"},
					"marketing_consent":       {"1"},
					"next_bill_date":          {"2020-05-13"},
					"next_payment_amount":     {"9.99"},
					"order_id":                {"123-456"},
					"passthrough":             {"some data"},
					"payment_method":          {"card"},
					"payment_tax":             {"0.25"},
					"plan_name":               {"plan1"},
					"quantity":                {"1"},
					"receipt_url":             {"http://my.paddle.com/receipt/receipt-url"},
					"sale_gross":              {"0.10"},
					"status":                  {"active"},
					"subscription_id":         {"123"},
					"subscription_payment_id": {"1234"},
					"subscription_plan_id":    {"12345"},
					"unit_price":              {"0.25"},
					"user_id":                 {"123456"},
					"p_signature":             {"signature"},
				},
			},
			want: SubscriptionPaymentSucceeded{
				AlertName:             AlertSubscriptionPaymentSucceeded,
				AlertID:               "27120763",
				CheckoutID:            "54832806-chrea1c514a3eb5-25c3040268",
				Currency:              "USD",
				Email:                 "test@example.org",
				EventTime:             time.Date(2020, 4, 28, 8, 42, 47, 0, time.UTC),
				MarketingConsent:      true,
				NextBillDate:          time.Date(2020, 5, 13, 0, 0, 0, 0, time.UTC),
				Passthrough:           "some data",
				Status:                StatusActive,
				SubscriptionID:        "123",
				SubscriptionPlanID:    "12345",
				UserID:                "123456",
				BalanceCurrency:       "USD",
				BalanceEarnings:       "10.00",
				BalanceFee:            "2.25",
				BalanceGross:          "2.30",
				BalanceTax:            "1.20",
				Country:               "US",
				Coupon:                "coupon",
				CustomerName:          "somename",
				Earnings:              "1.01",
				Fee:                   "0.52",
				InitialPayment:        true,
				Instalments:           "1",
				NextPaymentAmount:     "9.99",
				OrderID:               "123-456",
				PaymentMethod:         PaymentMethodCard,
				PaymentTax:            "0.25",
				PlanName:              "plan1",
				Quantity:              1,
				ReceiptURL:            "http://my.paddle.com/receipt/receipt-url",
				SaleGross:             "0.10",
				SubscriptionPaymentID: "1234",
				UnitPrice:             "0.25",
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
