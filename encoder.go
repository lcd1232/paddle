package paddle

import (
	"github.com/gorilla/schema"
)

var encoder *schema.Encoder

func init() {
	cb := customBool(true)

	encoder = schema.NewEncoder()
	encoder.RegisterEncoder(cb, customBoolEncoder)
	encoder.RegisterEncoder(&cb, customBoolEncoder)
}
