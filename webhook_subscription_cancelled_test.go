package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseSubscriptionCancelledWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    SubscriptionCancelled
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":                  {"subscription_cancelled"},
					"alert_id":                    {"27120763"},
					"cancellation_effective_date": {"2020-05-14"},
					"checkout_id":                 {"54832806-chrea1c514a3eb5-25c3040268"},
					"currency":                    {"USD"},
					"email":                       {"test@example.org"},
					"event_time":                  {"2020-04-28 08:42:47"},
					"marketing_consent":           {"1"},
					"passthrough":                 {"some data"},
					"quantity":                    {"1"},
					"status":                      {"deleted"},
					"subscription_id":             {"1234"},
					"subscription_plan_id":        {"12345"},
					"unit_price":                  {"9.99"},
					"user_id":                     {"123456"},
					"p_signature":                 {"signature"},
				},
			},
			want: SubscriptionCancelled{
				AlertName:                 AlertSubscriptionCancelled,
				AlertID:                   "27120763",
				CancellationEffectiveDate: time.Date(2020, 5, 14, 0, 0, 0, 0, time.UTC),
				CheckoutID:                "54832806-chrea1c514a3eb5-25c3040268",
				Currency:                  "USD",
				Email:                     "test@example.org",
				EventTime:                 time.Date(2020, 4, 28, 8, 42, 47, 0, time.UTC),
				MarketingConsent:          true,
				Passthrough:               "some data",
				Quantity:                  1,
				Status:                    StatusDeleted,
				SubscriptionID:            "1234",
				SubscriptionPlanID:        "12345",
				UnitPrice:                 "9.99",
				UserID:                    "123456",
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
			got, err := c.ParseSubscriptionCancelledWebhook(tt.args.form)
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
