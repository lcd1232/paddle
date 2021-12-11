package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseTransferPaidWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name        string
		args        args
		want        TransferPaid
		getVerifier func(t *testing.T) (verifier, func())
		wantErr     bool
	}{
		{
			name: "valid form data",
			args: args{
				form: url.Values{
					"alert_name":  {"transfer_paid"},
					"alert_id":    {"27120763"},
					"amount":      {"1234"},
					"currency":    {"USD"},
					"event_time":  {"2020-10-30 00:00:00"},
					"payout_id":   {"1234"},
					"status":      {"paid"},
					"p_signature": {"signature"},
				},
			},
			want: TransferPaid{
				AlertName: AlertTransferPaid,
				AlertID:   "27120763",
				Amount:    "1234",
				Currency:  "USD",
				EventTime: time.Date(2020, 10, 30, 0, 0, 0, 0, time.UTC),
				PayoutID:  "1234",
				Status:    StatusPaid,
			},
		},
		{
			name: "verifier error",
			getVerifier: func(t *testing.T) (verifier, func()) {
				v := new(verifierMock)
				v.On("Verify",
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(errors.New("invalid signature")).Once()
				return v, func() {
					v.AssertExpectations(t)
				}
			},
			wantErr: true,
		},
		{
			name: "decode error",
			args: args{
				form: url.Values{
					"event_time": {"asd"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, vm := NewTestWebhookClient()
			if tt.getVerifier != nil {
				vm, f := tt.getVerifier(t)
				defer f()
				c.verifier = vm
			} else {
				vm.On("Verify",
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(
					nil,
				).Once()
			}
			got, err := c.ParseTransferPaidWebhook(tt.args.form)
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
