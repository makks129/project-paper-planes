package utils

import (
	"bytes"
	"encoding/json"

	"github.com/makks129/project-paper-planes/src/utils"
)

func FromJson[T any](b *bytes.Buffer) T {
	var str T
	err := json.Unmarshal(b.Bytes(), &str)
	if err != nil {
		utils.Log("ERROR FromJson", err)
	}
	return str
}
