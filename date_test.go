package paddle

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateUnmarshalJSON(t *testing.T) {
	for _, td := range []struct {
		data string
		want time.Time
	}{
		{
			data: `"2019-08-01 21:24:35.000000"`,
			want: time.Date(2019, 8, 1, 21, 24, 35, 0, time.UTC),
		},
	} {
		t.Run(td.data, func(t *testing.T) {
			got := Time{}
			require.NoError(t, got.UnmarshalJSON([]byte(td.data)))
			assert.True(t, td.want.Equal(time.Time(got)))
		})
	}
}
