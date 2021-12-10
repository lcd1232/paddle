package paddle

import (
	"reflect"
	"time"

	"github.com/pkg/errors"
)

const dateLayout = "2006-01-02"

type customDate int64

func (cd *customDate) UnmarshalText(text []byte) error {
	s := string(text)
	t, err := parseDate(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*cd = customDate(t.Unix())
	return nil
}

func (cd *customDate) Time() time.Time {
	if cd == nil {
		return time.Time{}
	}
	return time.Unix(int64(*cd), 0)
}

func parseDate(s string) (time.Time, error) {
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return t, nil
}

func customDateEncoder(value reflect.Value) string {
	var t time.Time
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return ""
		}
		cd := value.Interface().(*customDate)
		t = cd.Time()
	} else {
		cd := value.Interface().(customDate)
		t = cd.Time()
	}
	return t.UTC().Format(dateLayout)
}
