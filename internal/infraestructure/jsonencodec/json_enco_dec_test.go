package jsonencodec

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJSONEncoderFactory(t *testing.T) {
	factory := NewJSONEncoderFactory()
	assert.NotNil(t, factory)

	buffer := &bytes.Buffer{}
	encoder := factory(buffer)
	assert.NotNil(t, encoder)

	obj := map[string]string{"hello": "world"}
	err := encoder.Encode(obj)
	assert.NoError(t, err)

	var decoded map[string]string
	err = json.Unmarshal(buffer.Bytes(), &decoded)
	assert.NoError(t, err)
	assert.Equal(t, obj, decoded)
}
