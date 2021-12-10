package paddle

import (
	"reflect"
	"time"

	"github.com/pkg/errors"
)

const timeLayout = "2006-01-02 15:04:05"

type customTime int64

func (ct *customTime) UnmarshalText(text []byte) error {
	s := string(text)
	t, err := parseTime(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*ct = customTime(t.Unix())
	return nil
}

func (ct *customTime) Time() time.Time {
	if ct == nil {
		return time.Time{}
	}
	return time.Unix(int64(*ct), 0)
}

func parseTime(s string) (time.Time, error) {
	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return t, nil
}

func customTimeEncoder(value reflect.Value) string {
	var t time.Time
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return ""
		}
		ct := value.Interface().(*customTime)
		t = ct.Time()
	} else {
		ct := value.Interface().(customTime)
		t = ct.Time()
	}
	return t.UTC().Format(timeLayout)
}
