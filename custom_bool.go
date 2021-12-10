package paddle

import (
	"reflect"

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

func customBoolEncoder(value reflect.Value) string {
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return ""
	}
	b := value.Bool()
	if b {
		return "1"
	}
	return "0"
}
