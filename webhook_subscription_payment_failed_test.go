package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseSubscriptionPaymentFailedWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    SubscriptionPaymentFailed
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":              {"subscription_payment_failed"},
					"alert_id":                {"27120763"},
					"amount":                  {"2.99"},
					"cancel_url":              {"https://checkout.paddle.com/subscription/cancel-url"},
					"checkout_id":             {"54832806-chrea1c514a3eb5-25c3040269"},
					"currency":                {"USD"},
					"email":                   {"test@example.org"},
					"event_time":              {"2020-04-28 08:42:47"},
					"marketing_consent":       {"1"},
					"next_retry_date":         {"2020-05-13"},
					"passthrough":             {"some data"},
					"quantity":                {"1"},
					"status":                  {"past_due"},
					"subscription_id":         {"123"},
					"subscription_plan_id":    {"12345"},
					"unit_price":              {"0.25"},
					"update_url":              {"https://checkout.paddle.com/subscription/update-url"},
					"subscription_payment_id": {"1234"},
					"installments":            {"1"},
					"order_id":                {"123-4567"},
					"user_id":                 {"123456"},
					"attempt_number":          {"1"},
					"p_signature":             {"signature"},
				},
			},
			want: SubscriptionPaymentFailed{
				AlertName:             AlertSubscriptionPaymentFailed,
				AlertID:               "27120763",
				Amount:                "2.99",
				CancelURL:             "https://checkout.paddle.com/subscription/cancel-url",
				CheckoutID:            "54832806-chrea1c514a3eb5-25c3040269",
				Currency:              "USD",
				Email:                 "test@example.org",
				EventTime:             time.Date(2020, 4, 28, 8, 42, 47, 0, time.UTC),
				MarketingConsent:      true,
				NextRetryDate:         time.Date(2020, 5, 13, 0, 0, 0, 0, time.UTC),
				Passthrough:           "some data",
				Quantity:              1,
				Status:                StatusPastDue,
				SubscriptionID:        "123",
				SubscriptionPlanID:    "12345",
				UnitPrice:             "0.25",
				UpdateURL:             "https://checkout.paddle.com/subscription/update-url",
				SubscriptionPaymentID: "1234",
				Installments:          1,
				OrderID:               "123-4567",
				UserID:                "123456",
				AttemptNumber:         1,
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
			got, err := c.ParseSubscriptionPaymentFailedWebhook(tt.args.form)
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
