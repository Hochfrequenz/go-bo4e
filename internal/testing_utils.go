package internal

import (
	"bytes"
	"encoding/json"
)

func CompareAsJson[T any, T2 any](a T, a2 T2) (areEqual bool, err error) {
	bytesFromA, err := json.Marshal(a)
	if err != nil {
		return
	}
	bytesFromA2, err := json.Marshal(a2)
	if err != nil {
		return
	}

	return bytes.Equal(bytesFromA, bytesFromA2), nil
}
