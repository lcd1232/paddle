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
			name: "1 is true",
			args: args{
				s: "1",
			},
			want: true,
		},
		{
			name: "0 is false",
			args: args{
				s: "0",
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				s: "true",
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				s: "false",
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

func TestSetVerification(t *testing.T) {
	c := Client{}
	c.SetVerification(true)
	_, ok := c.verifier.(*realVerifier)
	assert.True(t, ok)
	c.SetVerification(false)
	_, ok = c.verifier.(*noonVerifier)
	assert.True(t, ok)
}
