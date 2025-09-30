package jsonencodec

import (
	"encoding/json"
	"io"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/handler"
)

func NewJSONEncoderFactory() handler.EncoderFactory {
	return func(w io.Writer) handler.Encoder {
		return json.NewEncoder(w)
	}
}
