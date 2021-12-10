package paddle

import (
	"reflect"
	"time"

	"github.com/pkg/errors"
)

type customTime time.Time

func (ct *customTime) UnmarshalText(text []byte) error {
	s := string(text)
	t, err := parseTime(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*ct = customTime(t)
	return nil
}

func customTimeEncoder(value reflect.Value) string {
	var t time.Time
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return ""
		}
		ct := value.Interface().(*customTime)
		t = time.Time(*ct)
	} else {
		t = time.Time(value.Interface().(customTime))
	}
	return t.Format("2006-01-02 15:04:05")
}
