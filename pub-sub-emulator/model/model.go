package model

import (
	"bytes"
	"encoding/json"
)

type (
	PubRequest struct {
		Topic string
		Data  interface{}
	}

	SubRequest struct {
		Topic  string
		Subs   string
		Data   []byte
		IsNack bool
	}
)

func (p *PubRequest) DataToByteArray() []byte {

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(p.Data); err != nil {
		return nil
	}

	return b.Bytes()
}
