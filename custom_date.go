package paddle

import (
	"time"

	"github.com/pkg/errors"
)

const dateLayout = "2006-01-02"

type customDate time.Time

func (cd *customDate) UnmarshalText(text []byte) error {
	s := string(text)
	t, err := parseDate(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*cd = customDate(t)
	return nil
}

func parseDate(s string) (time.Time, error) {
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return t, nil
}
