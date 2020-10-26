package paddle

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewTestClient() (*WebhookClient, *verifierMock) {
	vm := new(verifierMock)
	return &WebhookClient{
		verifier: vm,
	}, vm
}

func Test_parseTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "valid time",
			args: args{
				s: "2020-04-29 03:49:35",
			},
			want: time.Date(2020, 04, 29, 3, 49, 35, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTime(tt.args.s)
			if tt.wantErr {
				require.NoError(t, err)
				return
			}
			require.NoError(t, err)
			assert.True(t, tt.want.Equal(got))
		})
	}
}

func Test_parseDate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "valid date",
			args: args{
				s: "2020-05-02",
			},
			want: time.Date(2020, 5, 2, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDate(tt.args.s)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.True(t, tt.want.Equal(got))
		})
	}
}

func Test_parseBool(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "true",
			args: args{
				s: "1",
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				s: "0",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBool(tt.args.s)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
