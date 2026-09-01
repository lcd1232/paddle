package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseSubscriptionCreatedWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    SubscriptionCreated
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":           {"subscription_created"},
					"alert_id":             {"27120763"},
					"cancel_url":           {"https://checkout.paddle.com/subscription/cancel-url"},
					"checkout_id":          {"54832806-chrea1c514a3eb5-25c3040268"},
					"currency":             {"USD"},
					"email":                {"test@example.org"},
					"event_time":           {"2020-04-28 08:42:47"},
					"marketing_consent":    {"1"},
					"next_bill_date":       {"2020-05-06"},
					"passthrough":          {"some data"},
					"quantity":             {"1"},
					"source":               {"example.org"},
					"status":               {"trialing"},
					"subscription_id":      {"1234"},
					"subscription_plan_id": {"12345"},
					"unit_price":           {"9.99"},
					"update_url":           {"https://checkout.paddle.com/subscription/update-url"},
					"user_id":              {"123456"},
					"p_signature":          {"signature"},
				},
			},
			want: SubscriptionCreated{
				AlertName:          AlertSubscriptionCreated,
				AlertID:            "27120763",
				CancelURL:          "https://checkout.paddle.com/subscription/cancel-url",
				CheckoutID:         "54832806-chrea1c514a3eb5-25c3040268",
				Currency:           "USD",
				Email:              "test@example.org",
				EventTime:          time.Date(2020, 4, 28, 8, 42, 47, 0, time.UTC),
				MarketingConsent:   true,
				NextBillDate:       time.Date(2020, 5, 6, 0, 0, 0, 0, time.UTC),
				Passthrough:        "some data",
				Quantity:           1,
				Source:             "example.org",
				Status:             StatusTrialing,
				SubscriptionID:     "1234",
				SubscriptionPlanID: "12345",
				UnitPrice:          "9.99",
				UserID:             "123456",
				UpdateURL:          "https://checkout.paddle.com/subscription/update-url",
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
			got, err := c.ParseSubscriptionCreatedWebhook(tt.args.form)
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
