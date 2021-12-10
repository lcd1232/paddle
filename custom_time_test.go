package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomTimeEncoder(t *testing.T) {
	type A struct {
		B customTime  `schema:"b"`
		C *customTime `schema:"c"`
	}
	ct := customTime(time.Date(2021, 05, 11, 15, 21, 55, 0, time.UTC))
	for _, tc := range []struct {
		name string
		data A
		key  string
		want string
	}{
		{
			name: "2021-05-11 15:21:55",
			data: A{
				B: ct,
			},
			key:  "b",
			want: "2021-05-11 15:21:55",
		},
		{
			name: "nil",
			data: A{
				C: nil,
			},
			key:  "c",
			want: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			form := url.Values{}
			require.NoError(t, encoder.Encode(tc.data, form))
			assert.Equal(t, tc.want, form.Get(tc.key))
		})
	}
}
