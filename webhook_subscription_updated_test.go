package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseSubscriptionUpdatedWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    SubscriptionUpdated
		wantErr bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":               {"subscription_updated"},
					"alert_id":                 {"27120763"},
					"cancel_url":               {"https://checkout.paddle.com/subscription/cancel-url"},
					"checkout_id":              {"54832806-chrea1c514a3eb5-25c3040268"},
					"currency":                 {"USD"},
					"email":                    {"test@example.org"},
					"event_time":               {"2020-04-28 08:42:47"},
					"marketing_consent":        {"1"},
					"new_price":                {"35.880"},
					"new_quantity":             {"1"},
					"new_unit_price":           {"35.88"},
					"next_bill_date":           {"2020-05-13"},
					"old_next_bill_date":       {"2020-05-10"},
					"old_price":                {"35.880"},
					"old_quantity":             {"1"},
					"old_status":               {"trialing"},
					"old_subscription_plan_id": {"12345"},
					"old_unit_price":           {"35.880"},
					"passthrough":              {"some data"},
					"status":                   {"past_due"},
					"subscription_id":          {"1234"},
					"subscription_plan_id":     {"1234567"},
					"update_url":               {"https://checkout.paddle.com/subscription/update-url"},
					"user_id":                  {"123456"},
					"p_signature":              {"signature"},
				},
			},
			want: SubscriptionUpdated{
				AlertName:             AlertSubscriptionUpdated,
				AlertID:               "27120763",
				CancelURL:             "https://checkout.paddle.com/subscription/cancel-url",
				CheckoutID:            "54832806-chrea1c514a3eb5-25c3040268",
				Currency:              "USD",
				Email:                 "test@example.org",
				EventTime:             time.Date(2020, 4, 28, 8, 42, 47, 0, time.UTC),
				MarketingConsent:      true,
				NewPrice:              "35.880",
				NewQuantity:           1,
				NewUnitPrice:          "35.88",
				NextBillDate:          time.Date(2020, 5, 13, 0, 0, 0, 0, time.UTC),
				OldNextBillDate:       time.Date(2020, 5, 10, 0, 0, 0, 0, time.UTC),
				OldPrice:              "35.880",
				OldQuantity:           1,
				OldStatus:             StatusTrialing,
				OldSubscriptionPlanID: "12345",
				OldUnitPrice:          "35.880",
				Passthrough:           "some data",
				Status:                StatusPastDue,
				SubscriptionID:        "1234",
				SubscriptionPlanID:    "1234567",
				UpdateURL:             "https://checkout.paddle.com/subscription/update-url",
				UserID:                "123456",
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
			got, err := c.ParseSubscriptionUpdatedWebhook(tt.args.form)
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
