package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParseTransferPaidWebhook(t *testing.T) {
	type args struct {
		form url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    TransferPaid
		wantErr bool
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
