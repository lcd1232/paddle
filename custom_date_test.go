package paddle

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomDateEncoder(t *testing.T) {
	type A struct {
		B customDate  `schema:"b"`
		C *customDate `schema:"c"`
	}
	cd := customDate(time.Date(2021, 05, 11, 15, 21, 55, 0, time.UTC))
	for _, tc := range []struct {
		name string
		data A
		key  string
		want string
	}{
		{
			name: "value",
			data: A{
				B: cd,
			},
			key:  "b",
			want: "2021-05-11",
		},
		{
			name: "pointer",
			data: A{
				C: &cd,
			},
			key:  "c",
			want: "2021-05-11",
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
