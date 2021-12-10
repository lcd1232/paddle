package paddle

import (
	"github.com/gorilla/schema"
)

var encoder *schema.Encoder

func init() {

	encoder = schema.NewEncoder()
	{
		cb := customBool(true)
		encoder.RegisterEncoder(cb, customBoolEncoder)
		encoder.RegisterEncoder(&cb, customBoolEncoder)
	}
	{
		ct := customTime{}
		encoder.RegisterEncoder(ct, customTimeEncoder)
		encoder.RegisterEncoder(&ct, customTimeEncoder)
	}
}
