package paddle

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomBoolEncoder(t *testing.T) {
	type A struct {
		B customBool  `schema:"b"`
		C *customBool `schema:"c"`
	}
	pTrue := customBool(true)
	pFalse := customBool(false)
	for _, tc := range []struct {
		name string
		data A
		key  string
		want string
	}{
		{
			name: "true",
			data: A{
				B: customBool(true),
			},
			key:  "b",
			want: "1",
		},
		{
			name: "false",
			data: A{
				B: customBool(false),
			},
			key:  "b",
			want: "0",
		},
		{
			name: "pointer of true",
			data: A{
				C: &pTrue,
			},
			key:  "c",
			want: "1",
		},
		{
			name: "pointer of false",
			data: A{
				C: &pFalse,
			},
			key:  "c",
			want: "0",
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
