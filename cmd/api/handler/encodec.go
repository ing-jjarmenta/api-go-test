package handler

import "io"

type Encoder interface {
	Encode(v any) error
}

type EncoderFactory func(w io.Writer) Encoder
