package utils

import (
	"bytes"
	"encoding/json"
)

func FromJson[T any](b *bytes.Buffer) T {
	var str T
	json.Unmarshal(b.Bytes(), &str)
	return str
}
