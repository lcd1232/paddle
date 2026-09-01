package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseSubscriptionPaymentRefundedWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    SubscriptionPaymentRefunded
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":                {"subscription_payment_refunded"},
					"alert_id":                  {"27120763"},
					"amount":                    {"9.99"},
					"balance_currency":          {"USD"},
					"balance_earnings_decrease": {"9.90"},
					"balance_fee_refund":        {"2.25"},
					"balance_gross_refund":      {"2.30"},
					"balance_tax_refund":        {"1.20"},
					"checkout_id":               {"54832806-chrea1c514a3eb5-25c3040268"},
					"currency":                  {"USD"},
					"earnings_decrease":         {"1.01"},
					"email":                     {"test@example.org"},
					"event_time":                {"2020-04-28 08:42:47"},
					"fee_refund":                {"0.52"},
					"gross_refund":              {"0.10"},
					"initial_payment":           {"1"},
					"instalments":               {"1"},
					"marketing_consent":         {"1"},
					"order_id":                  {"123-456"},
					"passthrough":               {"some data"},
					"quantity":                  {"1"},
					"refund_reason":             {"refund reason"},
					"refund_type":               {"full"},
					"status":                    {"active"},
					"subscription_id":           {"123"},
					"subscription_payment_id":   {"1234"},
					"subscription_plan_id":      {"12345"},
					"tax_refund":                {"1.01"},
					"unit_price":                {"0.25"},
					"user_id":                   {"123456"},
					"p_signature":               {"signature"},
				},
			},
			want: SubscriptionPaymentRefunded{
				AlertName:               AlertSubscriptionPaymentRefunded,
				AlertID:                 "27120763",
				CheckoutID:              "54832806-chrea1c514a3eb5-25c3040268",
				Currency:                "USD",
				Email:                   "test@example.org",
				EventTime:               time.Date(2020, 4, 28, 8, 42, 47, 0, time.UTC),
				MarketingConsent:        true,
				Passthrough:             "some data",
				Status:                  StatusActive,
				SubscriptionID:          "123",
				SubscriptionPlanID:      "12345",
				UserID:                  "123456",
				BalanceCurrency:         "USD",
				BalanceFeeRefund:        "2.25",
				BalanceGrossRefund:      "2.30",
				BalanceTaxRefund:        "1.20",
				EarningsDecrease:        "1.01",
				FeeRefund:               "0.52",
				InitialPayment:          true,
				Instalments:             "1",
				OrderID:                 "123-456",
				Quantity:                1,
				SubscriptionPaymentID:   "1234",
				UnitPrice:               "0.25",
				Amount:                  "9.99",
				BalanceEarningsDecrease: "9.90",
				GrossRefund:             "0.10",
				RefundReason:            "refund reason",
				RefundType:              RefundTypeFull,
				TaxRefund:               "1.01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, vm := NewTestWebhookClient()
			vm.On("Verify",
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(
				nil,
			).Once()
			got, err := c.ParseSubscriptionPaymentRefundedWebhook(tt.args.form)
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
