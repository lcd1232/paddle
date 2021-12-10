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
	panic("implement me")
}
