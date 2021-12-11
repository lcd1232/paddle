package paddle

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type customBool bool

func (cb *customBool) UnmarshalText(text []byte) error {
	s := string(text)
	b, err := parseBool(s)
	if err != nil {
		return errors.WithStack(err)
	}
	*cb = customBool(b)
	return nil
}

func toCustomBoolPointer(b *bool) *customBool {
	if b == nil {
		return nil
	}
	cb := customBool(*b)
	return &cb
}

func parseBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return b, nil
}

func customBoolEncoder(value reflect.Value) string {
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return ""
		}
		value = value.Elem()
	}
	b := value.Bool()
	if b {
		return "1"
	}
	return "0"
}
