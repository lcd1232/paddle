package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParsePaymentDisputeCreatedWebhook(t *testing.T) {
	tests := []struct {
		name      string
		form      url.Values
		verifyErr error
		want      PaymentDisputeCreated
		wantErr   bool
	}{
		{
			name: "valid form data",
			form: url.Values{
				"alert_id":          {"649721479"},
				"alert_name":        {"payment_dispute_created"},
				"amount":            {"473.05"},
				"balance_amount":    {"452.78"},
				"balance_currency":  {"EUR"},
				"balance_fee":       {"424.33"},
				"checkout_id":       {"9-f4ddb7a7b26753e-34bf2cbd3d"},
				"currency":          {"USD"},
				"email":             {"roberts.rashawn@example.com"},
				"event_time":        {"2023-03-08 09:33:29"},
				"fee_usd":           {"0.46"},
				"marketing_consent": {"1"},
				"order_id":          {"2"},
				"passthrough":       {"Example String"},
				"status":            {"open"},
				"p_signature":       {"signature"},
			},
			want: PaymentDisputeCreated{
				AlertName:        AlertPaymentDisputeCreated,
				AlertID:          "649721479",
				Amount:           "473.05",
				BalanceAmount:    "452.78",
				BalanceCurrency:  "EUR",
				BalanceFee:       "424.33",
				CheckoutID:       "9-f4ddb7a7b26753e-34bf2cbd3d",
				Currency:         "USD",
				Email:            "roberts.rashawn@example.com",
				EventTime:        time.Date(2023, 3, 8, 9, 33, 29, 0, time.UTC),
				FeeUSD:           "0.46",
				MarketingConsent: true,
				OrderID:          "2",
				Passthrough:      "Example String",
				Status:           PaymentDisputeStatusOpen,
			},
		},
		{
			name:      "invalid signature",
			form:      url.Values{"alert_name": {"payment_dispute_created"}, "p_signature": {"bad"}},
			verifyErr: ErrInvalidSignature,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, vm := NewTestWebhookClient()
			vm.On("Verify", mock.Anything, mock.Anything, mock.Anything).Return(tt.verifyErr).Once()
			t.Cleanup(func() { vm.AssertExpectations(t) })
			got, err := c.ParsePaymentDisputeCreatedWebhook(tt.form)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
