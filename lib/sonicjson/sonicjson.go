// package json is a thin shim over the bytedance sonic library for fast json
// parsing

package json

import (
	"io"

	"github.com/bytedance/sonic"
)

var config = sonic.ConfigFastest

type Encoder sonic.Encoder

func NewEncoder(w io.Writer) Encoder {
	return Encoder(config.NewEncoder(w))
}

type Decoder sonic.Decoder

func NewDecoder(r io.Reader) Decoder {
	return Decoder(config.NewDecoder(r))
}

func Marshal(val any) ([]byte, error) {
	return config.Marshal(val)
}

func Unmarshal(data []byte, val any) error {
	return config.Unmarshal(data, val)
}

func MarshalIndent(val any, prefix, indent string) ([]byte, error) {
	return config.MarshalIndent(val, prefix, indent)
}
